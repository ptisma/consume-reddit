package processor

import (
	"fmt"
	"reddit-api-processor/internal/config"
	"reddit-api-processor/internal/model"

	"github.com/buger/jsonparser"
	amqp "github.com/rabbitmq/amqp091-go"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(config *config.PostgresConfig) (*gorm.DB, error) {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai", config.DBHostURI, config.DBUsername, config.DBUserPassword, config.DBName, config.DBPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, err
}

func CreateDatabase(config *config.PostgresConfig) error {
	var err error
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable TimeZone=Asia/Shanghai", config.DBHostURI, config.DBPort, config.DBUsername, config.DBUserPassword)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	createDatabaseCommand := fmt.Sprintf("CREATE DATABASE %s", config.DBName)

	err = db.Exec(createDatabaseCommand).Error
	if err != nil {
		return err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	sqlDB.Close()

	return err
}

func ConnectBroker(config *config.RabbitMQConfig) (*amqp.Channel, string, error) {
	var queueName string
	conn, err := amqp.Dial(config.BrokerURI)
	if err != nil {
		return nil, queueName, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, queueName, err
	}

	if config.AutoCreateRabbitMQ{
		err = ch.ExchangeDeclare(
			config.ExchangeName, // name
			"topic",          // type
			true,             // durable
			false,            // auto-deleted
			false,            // internal
			false,            // no-wait
			nil,              // arguments
		)
	} else {
		err = ch.ExchangeDeclarePassive(
			config.ExchangeName, // name
			"topic",          // type
			true,             // durable
			false,            // auto-deleted
			false,            // internal
			false,            // no-wait
			nil,              // arguments
		)
	}
	
	if err != nil {
		return nil, queueName, err
	}
	var q amqp.Queue
	if config.AutoCreateRabbitMQ{
		q, err = ch.QueueDeclare(
			config.QueueName,    // name
			true, // durable
			false, // delete when unused
			false,  // exclusive
			false, // no-wait
			nil,   // arguments
		)
	} else {
		q, err = ch.QueueDeclarePassive(
			config.QueueName,    // name
			true, // durable
			false, // delete when unused
			false,  // exclusive
			false, // no-wait
			nil,   // arguments
		)
	}
	
	if err != nil {
		return nil, queueName, err
	}

	err = ch.QueueBind(
		q.Name,            // queue name
		config.RoutingKey, // routing key
		config.ExchangeName,  // exchange
		false,
		nil)
	if err != nil {
		return nil, queueName, err
	}

	return ch, q.Name, err

}

type Processor struct {
	DB         *gorm.DB
	Exchange   *amqp.Channel
	QueueName  string
	ExchangeName  string
	RoutingKey string
}

func (p *Processor) ReadFromBroker() (<-chan amqp.Delivery, error) {
	msgs, err := p.Exchange.Consume(
		p.QueueName, // queue
		"",          // consumer
		true,        // auto ack
		false,       // exclusive
		false,       // no local
		false,       // no wait
		nil,         // args
	)

	return msgs, err

}

func (p *Processor) Process(body []byte, category string) (model.Post, error) {

	title, _, _, err := jsonparser.Get(body, "data", "title")
	content, _, _, err := jsonparser.Get(body, "data", "selftext")
	category = category

	post := model.Post{Title: string(title), Content: string(content), Category: category}

	return post, err

}

func (p *Processor) WriteToDB(post *model.Post) error {
	return p.DB.Create(post).Error

}

func (p *Processor) AutoMigrate() error {
	return p.DB.AutoMigrate(&model.Post{})

}

func GetProcessor(dbConfig *config.PostgresConfig, brokerConfig *config.RabbitMQConfig) (*Processor, error) {
	var err error
	// err = CreateDatabase(dbConfig)
	// if err != nil {
	// 	return nil, err
	// }
	db, err := ConnectDB(dbConfig)
	if err != nil {
		return nil, err
	}

	ch, queueName, err := ConnectBroker(brokerConfig)

	return &Processor{DB: db, Exchange: ch, QueueName: queueName, ExchangeName: brokerConfig.ExchangeName, RoutingKey: brokerConfig.RoutingKey}, err

}

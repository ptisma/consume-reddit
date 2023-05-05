package processor

import (
	"fmt"
	"reddit-api-processor/internal/config"
	"reddit-api-processor/internal/model"

	amqp "github.com/rabbitmq/amqp091-go"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(config *config.PostgresConfig) (*gorm.DB, error) {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", config.DBHost, config.DBUserName, config.DBUserPassword, config.DBName, config.DBPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, err
}

func ConnectBroker(config *config.RabbitMQConfig) (*amqp.Channel, error) {
	conn, err := amqp.Dial(config.BrokerURI)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	err = ch.ExchangeDeclare(
		config.TopicName, // name
		"topic",          // type
		true,             // durable
		false,            // auto-deleted
		false,            // internal
		false,            // no-wait
		nil,              // arguments
	)
	if err != nil {
		return nil, err
	}

	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		return nil, err
	}

	err = ch.QueueBind(
		q.Name,            // queue name
		config.RoutingKey, // routing key
		config.TopicName,  // exchange
		false,
		nil)
	if err != nil {
		return nil, err
	}

	return ch, err

}

type Processor struct {
	DB         *gorm.DB
	Exchange   *amqp.Channel
	TopicName  string
	RoutingKey string
}

func (p *Processor) Test() {

}

func (p *Processor) WriteToDB() {
	p.DB.AutoMigrate(&model.Message{})
	p.DB.Create(&model.Message{Body: "kek"})

}

func GetProcessor(dbConfig *config.PostgresConfig, brokerConfig *config.RabbitMQConfig) (*Processor, error) {
	var err error
	db, err := ConnectDB(dbConfig)
	if err != nil {
		return nil, err
	}

	ch, err := ConnectBroker(brokerConfig)

	return &Processor{DB: db, Exchange: ch}, err

}

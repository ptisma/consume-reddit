package main

import (
	"log"
	"net/http"

	"api-server/internal/controllers"
	"api-server/internal/initializers"
	"api-server/internal/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	server *gin.Engine

	PostController      controllers.PostController
	PostRouteController routes.PostRouteController
)

func init() {
	config, err := initializers.LoadConfig("../../")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)
	initializers.Automigrate()

	initializers.ConnectCache(&config)

	PostController = controllers.NewPostController(initializers.DB, initializers.Cache)
	PostRouteController = routes.NewRoutePostController(PostController)

	server = gin.Default()
}

func main() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"*"}

	server.Use(cors.New(corsConfig))

	router := server.Group("/api")
	router.GET("/healthchecker", func(ctx *gin.Context) {
		message := "Welcome to Golang with Gorm and Postgres"
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
	})

	PostRouteController.PostRoute(router)
	log.Fatal(server.Run(":" + config.ServerPort))
}

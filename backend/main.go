package main

import (
	"log"
	"context"
	"ezam/config"
	"ezam/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	router := gin.Default()

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Problem loading the .env file: %s", err)
	}

	config.InitializeDB()
	defer func() {
		if err := config.Client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	router.Use(cors.Default())
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	controllers.StudentRoute(router.Group("/student"))
	controllers.ExamRoute(router.Group("/exam"))
	router.Run()
}

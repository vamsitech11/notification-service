// main.go
package main

import (
	"notification-service/handlers"
	"notification-service/kafka"
	"notification-service/storage"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize MySQL
	storage.Connect()

	// Initialize Kafka
	kafka.InitProducer()

	r := gin.Default()

	r.POST("/subscribe", handlers.SubscribeHandler)
	r.POST("/notifications/send", handlers.SendNotificationHandler)
	r.POST("/unsubscribe", handlers.UnsubscribeHandler)
	r.GET("/subscriptions/:user_id", handlers.FetchSubscriptionsHandler)

	r.Run(":8081") // Default port 8080
}

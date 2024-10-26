// handlers/subscribe.go
package handlers

import (
	"net/http"
	"notification-service/storage"

	"github.com/gin-gonic/gin"
)

func SubscribeHandler(c *gin.Context) {
	var payload struct {
		UserID               string   `json:"user_id"`
		Topics               []string `json:"topics"`
		NotificationChannels struct {
			Email             string `json:"email"`
			Sms               string `json:"sms"`
			PushNotifications bool   `json:"push_notifications"`
		} `json:"notification_channels"`
	}

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Store each topic for the user in MySQL
	for _, topic := range payload.Topics {
		_, err := storage.DB.Exec("INSERT INTO subscriptions (user_id, topic, email, sms, push_notifications) VALUES (?, ?, ?, ?, ?)",
			payload.UserID, topic, payload.NotificationChannels.Email, payload.NotificationChannels.Sms, payload.NotificationChannels.PushNotifications)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to store subscription"})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "Subscription successful"})
}

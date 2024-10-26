// handlers/send_notification.go
package handlers

import (
	"net/http"
	"notification-service/kafka" // Ensure to import your kafka package

	"github.com/gin-gonic/gin"
)

func SendNotificationHandler(c *gin.Context) {
	var payload struct {
		Topic string `json:"topic"`
		Event struct {
			EventID   string                 `json:"event_id"`
			Timestamp string                 `json:"timestamp"`
			Details   map[string]interface{} `json:"details"`
		} `json:"event"`
		Message struct {
			Title string `json:"title"`
			Body  string `json:"body"`
		} `json:"message"`
	}

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Publish the message to Kafka
	err := kafka.Publish(payload.Topic, payload.Message.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send notification"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Notification sent"})
}

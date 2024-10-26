// handlers/unsubscribe.go
package handlers

import (
	"net/http"
	"notification-service/storage" // Ensure to import your storage package

	"github.com/gin-gonic/gin"
)

func UnsubscribeHandler(c *gin.Context) {
	var payload struct {
		UserID string   `json:"user_id"`
		Topics []string `json:"topics"`
	}

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Remove subscriptions from MySQL
	for _, topic := range payload.Topics {
		_, err := storage.DB.Exec("DELETE FROM subscriptions WHERE user_id = ? AND topic = ?", payload.UserID, topic)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unsubscribe"})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "Unsubscribed successfully"})
}

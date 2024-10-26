// handlers/fetch_subscriptions.go
package handlers

import (
	"net/http"
	"notification-service/storage" // Ensure to import your storage package

	"github.com/gin-gonic/gin"
)

func FetchSubscriptionsHandler(c *gin.Context) {
	userID := c.Param("user_id")

	// Query to fetch subscriptions for the given user_id
	rows, err := storage.DB.Query("SELECT topic, email, sms, push_notifications FROM subscriptions WHERE user_id = ?", userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch subscriptions"})
		return
	}
	defer rows.Close()

	// Define the subscriptions slice with the correct structure
	var subscriptions []struct {
		Topic                string
		NotificationChannels struct {
			Email             string
			Sms               string
			PushNotifications bool
		}
	}

	// Loop through the rows and scan into the correct struct
	for rows.Next() {
		var topic string
		var email string
		var sms string
		var pushNotifications bool

		// Scan the current row into variables
		if err := rows.Scan(&topic, &email, &sms, &pushNotifications); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan row"})
			return
		}

		// Create a new entry with the correct nested structure
		subscription := struct {
			Topic                string
			NotificationChannels struct {
				Email             string
				Sms               string
				PushNotifications bool
			}
		}{
			Topic: topic,
			NotificationChannels: struct {
				Email             string
				Sms               string
				PushNotifications bool
			}{
				Email:             email,
				Sms:               sms,
				PushNotifications: pushNotifications,
			},
		}

		// Append the new subscription to the slice
		subscriptions = append(subscriptions, subscription)
	}

	// Return the response
	c.JSON(http.StatusOK, gin.H{"user_id": userID, "subscriptions": subscriptions})
}

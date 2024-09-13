package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web/dao"
)

type ConsumerRequest struct {
	ConsumerId string `json:"consumer_id"` // The consumer ID from the post body
}

type KafkaHandler struct{}

func (KafkaHandler) SendTopic(c *gin.Context) {
	senderStatus := dao.SendMessage()
	if senderStatus != nil {
		c.JSON(200, gin.H{"error": 1, "msg": "can't send message to kafka"})
	}
}

func (KafkaHandler) ReceiveTopic(c *gin.Context) {
	var request ConsumerRequest

	// Bind the incoming JSON to the ConsumerRequest struct
	if err := c.ShouldBindJSON(&request); err != nil {
		// Return a 400 error if the JSON is invalid
		c.JSON(http.StatusBadRequest, gin.H{"error": 1, "msg": "Invalid request body"})
		return
	}

	// Call the dao.ReceiveTopic function with the ConsumerID
	err := dao.ReceiveTopic(request.ConsumerId)
	if err != nil {
		// If there was an error receiving the topic
		c.JSON(http.StatusInternalServerError, gin.H{"error": 1, "msg": "Error receiving messages from Kafka"})
		return
	}

	// Return success if messages are being received
	c.JSON(http.StatusOK, gin.H{"error": 0, "msg": "Successfully started receiving messages from Kafka"})
}

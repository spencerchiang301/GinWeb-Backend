package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web/db"
	"web/global"
	"web/handlers"
	"web/messaing"
)

func main() {
	r := gin.Default()
	global.Mysql = db.InitDB()
	global.MyKafkaWriter = messaing.KafkaWriter()
	global.MyKafkaReader = messaing.KafkaReader()

	r.GET("/", func(c *gin.Context) {
		data := map[string]interface{}{
			"message": "hello world",
			"tag":     "welcome",
		}
		c.AsciiJSON(http.StatusOK, data)
	})

	r.GET("/fishPrice", handlers.FishHandler{}.GetFishPrice)
	r.GET("/fishImage", handlers.FishHandler{}.GetFishImage)
	r.POST("/kafka/sendTopic", handlers.KafkaHandler{}.SendTopic)
	r.POST("/kafka/receiveTopic", handlers.KafkaHandler{}.ReceiveTopic)

	err := r.Run("0.0.0.0:8080")
	if err != nil {
		return
	}
}

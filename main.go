package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web/db"
	"web/global"
	"web/handlers"
)

func main() {
	r := gin.Default()
	global.Mysql = db.InitDB()

	r.GET("/", func(c *gin.Context) {
		data := map[string]interface{}{
			"message": "hello world",
			"tag":     "welcome",
		}
		c.AsciiJSON(http.StatusOK, data)
	})

	r.GET("/fishPrice", handlers.FishHandler{}.GetFishPrice)
	r.GET("/fishImage", handlers.FishHandler{}.GetFishImage)

	err := r.Run("0.0.0.0:8080")
	if err != nil {
		return
	}
}

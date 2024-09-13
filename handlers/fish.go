package handlers

import (
	"github.com/gin-gonic/gin"
	"web/dao"
)

type FishHandler struct{}

func (FishHandler) GetFishPrice(c *gin.Context) {
	fishPrices := dao.GetFishPrice()
	if fishPrices == nil {
		c.JSON(200, gin.H{"error": 1, "msg": "get fish price fail"})
	} else {
		c.JSON(200, gin.H{"error": 0, "msg": "get fish price success", "fish_prices": fishPrices})
	}
}

func (FishHandler) GetFishImage(c *gin.Context) {
	fishImages := dao.GetFishImage()
	if fishImages == nil {
		c.JSON(200, gin.H{"error": 1, "msg": "get fish image fail"})
	} else {
		c.JSON(200, gin.H{"error": 0, "msg": "get fish images success", "fish_images": fishImages})
	}
}

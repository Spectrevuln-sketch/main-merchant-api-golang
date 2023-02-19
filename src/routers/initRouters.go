package routers

import (
	"main-merchant/src/controllers/merchants"

	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine) {
	merchant := router.Group("/v1/merchant")
	merchants.MerchantRoutes(merchant)
}

package merchants

import (
	"main-merchant/src/handler/merchants"

	"github.com/gin-gonic/gin"
)

func MerchantRoutes(merchant *gin.RouterGroup) {

	merchant.GET("all-merchant", merchants.AllMerchant)
	merchant.GET("current-merchant/:id", merchants.CurrnetMerchant)
	merchant.POST("create-merchant", merchants.CreateMerchant)
	merchant.PUT("updated-merchant/:id", merchants.UpdateMerchant)
	merchant.DELETE("delete-merchant/:id", merchants.DeleteMerchant)
}

package merchants

import (
	"main-merchant/src/config"
	"main-merchant/src/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateMerchant(c *gin.Context) {
	var merchant models.Merchant
	if err := c.ShouldBindJSON(&merchant); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"responseCode": "99",
			"responseDesc": "INVALID JSON",
			"message":      err.Error(),
		})
		return
	}

	config.DB.Create(&merchant)
	c.JSON(http.StatusOK, gin.H{
		"responseCode": "00",
		"responseDesc": "SUCCESS",
		"data":         merchant,
	})
}

func AllMerchant(c *gin.Context) {
	var merchant []models.Merchant
	config.DB.Find(&merchant)
	c.JSON(http.StatusOK, gin.H{
		"responseCode": "00",
		"responseDesc": "SUCCESS",
		"data":         merchant,
	})
}

func CurrnetMerchant(c *gin.Context) {
	var merchant models.Merchant
	id := c.Param("id")

	if err := config.DB.First(&merchant, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"responseCode": "99",
				"responseDesc": "INVALID REQUEST",
				"message":      "Data Tidak Di Temukan",
			})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"responseCode": "00",
		"responseDesc": "SUCCESS",
		"data":         merchant,
	})
}

func UpdateMerchant(c *gin.Context) {
	var merchant models.Merchant
	id := c.Param("id")
	if err := c.ShouldBindJSON(&merchant); err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"responseCode": "99",
			"responseDesc": "INVALID REQUEST",
			"message":      "Data Tidak Di Temukan",
		})
		return
	}

	if err := config.DB.First(&merchant, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"responseCode": "99",
				"responseDesc": "INVALID REQUEST",
				"message":      "Data Tidak Di Temukan",
			})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	if config.DB.Model(&merchant).Where("id = ?", id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"responseCode": "99",
			"responseDesc": "CANNOT UPDATE MERCHANT",
			"message":      "tidak dapat update merchant",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"responseCode": "00",
		"responseDesc": "SUCCESS",
		"data":         merchant,
	})
}

func DeleteMerchant(c *gin.Context) {
	var merchant models.Merchant
	id := c.Param("id")

	if err := config.DB.First(&merchant, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"responseCode": "99",
				"responseDesc": "INVALID REQUEST",
				"message":      "Data Tidak Di Temukan",
			})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}
	if config.DB.Delete(&merchant, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"responseCode": "99",
			"responseDesc": "INVALID REQUEST",
			"message":      "Data Tidak Di Temukan",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"responseCode": "00",
		"responseDesc": "DATA SUCCSESSFULLY DELETED",
		"data":         &merchant,
	})
}

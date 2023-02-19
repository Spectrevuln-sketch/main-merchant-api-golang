package main

import (
	"main-merchant/src/config"
	"main-merchant/src/routers"

	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnv()
	config.ConnectDB()
}

func main() {
	router := gin.Default()
	routers.InitRoutes(router)
	router.Run()
}

package main

import (
	"github.com/gin-gonic/gin"
	"go_decryptEnc/src/controller"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/encrypt-decrypt", controller.HandleEncryptDecrypt)
	return r
}

func main() {
	router := setupRouter()
	router.Run() // Starts the Gin server on localhost
}

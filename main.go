package main

import (
	"github.com/gin-gonic/gin"
	"go_decryptEnc/src/controller"
)

func main() {
	r := gin.Default()
	r.POST("/crypto", controller.HandleEncryptDecrypt)
	r.Run(":8080")
}

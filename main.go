package main

import (
	"github.com/gin-gonic/gin"
	"go_decryptEnc/src/controller"
)

func main() {
	r := gin.Default()
	r.POST("/encrypt-decrypt", controller.HandleEncryptDecrypt)
	err := r.Run()
	if err != nil {
		return
	}
}

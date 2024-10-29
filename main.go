package main

import (
	"github.com/gin-gonic/gin"
	"go_decryptEnc/src/controller"
	"net/http"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/encrypt-decrypt", controller.HandleEncryptDecrypt)
	return r
}

func VercelHandler(w http.ResponseWriter, r *http.Request) {
	router := setupRouter()
	router.ServeHTTP(w, r)
}

func main() {
	router := setupRouter()
	router.Run()
}

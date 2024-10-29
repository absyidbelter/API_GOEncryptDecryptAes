package handler

import (
	"github.com/gin-gonic/gin"
	"go_decryptEnc/src/controller"
	"net/http"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/encrypt-decrypt", controller.HandleEncryptDecrypt)
	return r
}

func VercelHandler(w http.ResponseWriter, r *http.Request) {
	router := SetupRouter()
	router.ServeHTTP(w, r)
}

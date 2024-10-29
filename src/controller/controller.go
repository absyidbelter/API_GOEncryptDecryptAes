package controller

import (
	"github.com/gin-gonic/gin"
	"go_decryptEnc/src/model"
	"go_decryptEnc/src/service"
	"net/http"
)

func HandleEncryptDecrypt(c *gin.Context) {
	var payload model.RequestPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	salt := service.GetSalt(payload.EncSalt)
	var result string
	switch payload.Action {
	case "encrypt":
		result = service.EncryptAES(payload.Value, salt)
	case "decrypt":
		result = service.DecryptAES(payload.Value, salt)
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid action. Use 'encrypt' or 'decrypt'"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": result})
}

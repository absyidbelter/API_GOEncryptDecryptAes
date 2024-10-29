package service

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
)

func EncryptAES(plaintext, salt string) string {
	c, err := aes.NewCipher([]byte(salt))
	if err != nil {
		log.Println("Error in NewCipher:", err)
		return ""
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		log.Println("Error in NewGCM:", err)
		return ""
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		log.Println("Error in ReadFull:", err)
		return ""
	}

	value := gcm.Seal(nonce, nonce, []byte(plaintext), nil)
	return hex.EncodeToString(value)
}

// Fungsi untuk mendapatkan salt dari `encSalt`
func GetSalt(encSalt string) string {
	type Turningcheck struct {
		Check string `json:"check"`
	}

	var check Turningcheck
	decoded, _ := base64.StdEncoding.DecodeString(encSalt)
	decodedCheck := reverse(string(decoded))
	resultString, _ := base64.StdEncoding.DecodeString(decodedCheck)

	err := json.Unmarshal([]byte(resultString), &check)
	if err != nil {
		fmt.Println("Error decoding salt:", err)
		return ""
	}

	return check.Check
}

func DecryptAES(value, salt string) string {
	cipherText, err := hex.DecodeString(value)
	if err != nil {
		log.Println("Error decoding value:", err)
		return ""
	}

	c, err := aes.NewCipher([]byte(salt))
	if err != nil {
		log.Println("Error in NewCipher:", err)
		return ""
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		log.Println("Error in NewGCM:", err)
		return ""
	}

	nonceSize := gcm.NonceSize()
	if len(cipherText) < nonceSize {
		log.Println("Invalid ciphertext size")
		return ""
	}

	nonce, ciphertext := cipherText[:nonceSize], cipherText[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		log.Println("Error decrypting:", err)
		return ""
	}

	return string(plaintext)
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

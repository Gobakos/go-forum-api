package services

import (
	// "crypto/hmac"
	// "crypto/sha256"
	// "encoding/hex"
	// "encoding/json"
	// "errors"
	"test/models"
	// "os"
	// "github.com/golang-jwt/jwt/v5"
)

func GetUsers()([]*models.Users,error){
	return models.GetUsers()
}

func GetUser(id string)(models.Users,error){

	return models.GetUser(id)
}
func DeleteUser(id string) (bool, error) {
	deleted, err := models.DelUser(id)
	if err != nil {
		return false, err
	}
	return deleted, nil
}
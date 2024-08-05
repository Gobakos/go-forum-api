package services

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"os"
	"github.com/golang-jwt/jwt/v5"
	"test/models"
)

func IsValidLogin(requestBody []byte,) (*models.Login, error) {
	var login models.Login
	if err := json.Unmarshal(requestBody, &login); err != nil {
		return nil, errors.New("Invalid schema")
	}
	if login.Username == "" || login.Password == "" {
		return nil, errors.New("Missing required fields")
	}
	return &login, nil
}

func IsValidRegister(requestBody []byte,) (*models.Users, error){
	var register models.Users
	if err := json.Unmarshal(requestBody, &register); err != nil {
		return nil, errors.New("Invalid schema")
	}
	if register.Username == "" || register.Password == "" || register.Description == "" || register.Lastname == "" || register.Name == "" {
		return nil, errors.New("Missing required fields")
	}
	return &register, nil
}

func AddUser(user *models.Users) (error) {
	hashedPassword, err := HashedPasswordAuth(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword
	success, err := models.AddUser(user)
	if err != nil {
		return err
	}
	if !success {
		return errors.New("Failed to add user")
	}
	return nil
}

func HashedPasswordAuth(password string) (string, error) {
	/* 			HASH GENERATION				*/
	//Need the hash algorithm and also the key
	//Initializes a HMAC hasher

	hmacHasher := hmac.New(sha256.New, []byte(os.Getenv("PASSWORD_KEY")))
	_, err := hmacHasher.Write([]byte(password))
	if err != nil {
		return "", err
	}
	hashed := hmacHasher.Sum(nil)
	return hex.EncodeToString(hashed), nil
}

func TryLogin(username string, password string) (bool,string, string, error) {
	return models.TryLogin(username, password)
}

func TokenGen(claims jwt.Claims) (string, error) {

	/* 			TOKEN GENERATION			*/
	secretKey := []byte(os.Getenv("SECRET_KEY"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
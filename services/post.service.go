package services

import (
	"encoding/json"
	"errors"
	"time"
	"os"
	"github.com/golang-jwt/jwt/v5"
	"api/models"
	// "api/controllers"
)

type Claims struct {
	Username string `json:"username"`
	Id string `json:"Id"`
	Is_Admin   string   `json:"is_admin"`
	jwt.RegisteredClaims //Embedded standard
}

func IsValidPost(requestBody []byte, tokenString string) error {
	var post models.Posts
	userId,username,err := GetJwtId(tokenString)
	if err != nil {
		return err
	}
	post.User = &models.Users{Id: userId}
	post.Id = generatePostID()
	if err := json.Unmarshal(requestBody, &post); err != nil {
		return errors.New("Invalid Schema")
	}
	post.Username = username
	_,err=GetUser(userId)
	if err != nil{
		return errors.New("This user does not exist")
	}
	return AddPost(&post)
}

func GetJwtId(tokenString string) (string,string, error) {
	secretKey := []byte(os.Getenv("SECRET_KEY"))
	if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
        tokenString = tokenString[7:]
    }
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return "","", err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return "","", errors.New("Invalid token")
	}

	return claims.Id,claims.Username, nil
}

func AddPost(post *models.Posts) error {
	success, err := models.AddPost(post)
	if err != nil {
		return err
	}
	if !success {
		return errors.New("Failed to add post")
	}
	return nil
}

func generatePostID() string {
	return "post_" + time.Now().Format("20060102150405")
}

func GetPostUserId(post_id string) (string,error){
	return models.GetPostUserId(post_id)
}

func DeletePost(postID string) error {
    success, err := models.DeletePost(postID)
    if err != nil {
        return err
    }
    if !success {
        return errors.New("Failed to delete post")
    }
    return nil
}
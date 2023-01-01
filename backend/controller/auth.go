package controller

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func GenerateJWT(uid string) string {
	claims := jwt.MapClaims{
		"userid": uid,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString([]byte("asdfasdf"))
	log.Println(tokenString)
	return tokenString
}

func CheckJWT(c *gin.Context) {
	tokenString := c.Request.Header.Get("token")
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			fmt.Println("err!!!")
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte("Miteh"), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		fmt.Println(claims["user_id"])
		c.JSON(200, claims["user_id"])
	}
}

package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func GenerateJwt(user_id string, role string) (string, error) {
	key := []byte(os.Getenv("JWT_TOKEN"))
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user_id,
		"aud": role,
		"exp": time.Now().Add(time.Hour * 48).Unix(),
		"iat": time.Now().Unix(),
	})
	s, err := t.SignedString(key)

	if err != nil {
		return "", err
	}
	return s, nil
}

func VerifyJwt(token string) (*jwt.Token, error) {
	tokenString, err := jwt.Parse(token, func(tokenString *jwt.Token) (any, error) { return os.Getenv("JWT_TOKEN"), nil })
	if err != nil {
		return nil, err
	}
	if !tokenString.Valid {
		return nil, fmt.Errorf("Invalid Token")	
	}
	return tokenString, nil
}

func AuthMiddleware(c *gin.Context) {
	tokenString, err := c.Cookie("token")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H { "error" : "No cookie in http header"})
		return
	}
	_, err = VerifyJwt(tokenString)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H { "error" : "Token Invalid"})
		return
	}

	c.Next()
}

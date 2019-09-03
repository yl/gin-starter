package jwt

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/yangliulnn/gin-starter/configs"
	"github.com/yangliulnn/gin-starter/httpd/models"
	"strconv"
	"time"
)

func Generate(u *models.User) (string, error) {
	key := []byte("password")
	claims := jwt.StandardClaims{
		ExpiresAt: time.Now().Add(configs.JWT.TTL * time.Second).Unix(),
		Id:        uuid.New().String(),
		IssuedAt:  time.Now().Unix(),
		NotBefore: time.Now().Unix(),
		Subject:   strconv.Itoa(u.ID) + "q",
	}

	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(key)
}

func Check(t string) (*models.User, error) {
	token, err := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v ", token.Header["alg"])
		}

		return []byte("password"), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("Claims is not jwt.MapClaims ")
	}

	err = claims.Valid()
	if err != nil {
		return nil, err
	}

	sub := claims["sub"]
	id, err := strconv.Atoi(sub.(string))
	if err != nil {
		return nil, errors.New("token is invalid")
	}
	user := models.NewUser()
	err = user.FirstBy("id", id)
	if err != nil {
		return nil, err
	}
	fmt.Println(user)
	return user, nil
}

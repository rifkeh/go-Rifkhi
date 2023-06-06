package middleware

import (
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)



func CreateJWT(id uint) (string, error){
	errenv := godotenv.Load()

	if errenv != nil{
		log.Fatal("Error loading .env file")
	}
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["admin_id"] = id
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}
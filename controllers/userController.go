package controllers

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func SignUp() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
func Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}

func GetUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}
func GetUsers() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}
func HashPassword(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		panic(err)
	}
	return string(hashedPassword)
}
func VerifyPassword(userPassword string, providedPassword string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(providedPassword), []byte(userPassword))
	msg := ""
	check := true
	if err != nil {
		check = false
		msg = "password is Incorrect"
	}
	return check, msg
}

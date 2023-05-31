package controllers

import "github.com/gin-gonic/gin"


func GetUsers() gin.HandlerFunc{
	return func(ctx *gin.Context) {

	}
}

func GetUser() gin.HandlerFunc{
	return func(ctx *gin.Context) {

	}
}

func SignUp() gin.HandlerFunc{
	return func(ctx *gin.Context) {

	}
}

func LogIn() gin.HandlerFunc{
	return func(ctx *gin.Context) {

	}
}

func HashPassword(password string) string{

}

func VerifyPassword(password, providedPassword string)(bool, string){
	
}

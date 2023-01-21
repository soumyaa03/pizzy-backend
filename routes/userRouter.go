package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/soumyaa03/pizzy-backend/controllers"
	"github.com/soumyaa03/pizzy-backend/middleware"
)

func UserRoutes(incomingRoutes *gin.Engine) {

	incomingRoutes.GET("/users/signup", controller.SignUp())
	incomingRoutes.GET("/users/login", controller.Login())
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/users/:user_id", controller.GetUser())

}

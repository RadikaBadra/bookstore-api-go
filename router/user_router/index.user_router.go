package user_router

import (
	controller "bookstore-api/controller/user_controller"

	"github.com/gin-gonic/gin"
)

func UserRouter(router *gin.Engine) {
	router.POST("/user/register", controller.Register)
	router.POST("/user/login", controller.Login)
}

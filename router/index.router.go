package router

import (
	"bookstore-api/router/book_router"
	"bookstore-api/router/user_router"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRoute(router *gin.Engine){

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message" : "server connected",
		})
	})

	book_router.BookRouter(router)
	user_router.UserRouter(router)
}
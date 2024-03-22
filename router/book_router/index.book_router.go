package book_router

import (
	controller "bookstore-api/controller/book_controller"

	"github.com/gin-gonic/gin"
)

func BookRouter(router *gin.Engine){
	router.GET("/books", controller.GetAllBooks)
	router.GET("/books/:id", controller.GetBook)
	router.POST("/books", controller.CreateBook)
	router.PATCH("/books/:id", controller.UpdateBook)
	router.DELETE("/books/:id", controller.DeleteBook)
}
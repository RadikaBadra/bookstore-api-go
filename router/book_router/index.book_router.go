package book_router

import (
	controller "bookstore-api/controller/book_controller"
	"bookstore-api/middleware"

	"github.com/gin-gonic/gin"
)

func BookRouter(router *gin.Engine) {
	router.GET("/books", middleware.RequireAuth, controller.GetAllBooks)
	router.GET("/books/:id", middleware.RequireAuth, controller.GetBook)
	router.POST("/books", middleware.RequireAuth, controller.CreateBook)
	router.PATCH("/books/:id", middleware.RequireAuth, controller.UpdateBook)
	router.DELETE("/books/:id", middleware.RequireAuth, controller.DeleteBook)
}

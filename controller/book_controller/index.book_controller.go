package controller

import (
	"bookstore-api/database"
	models "bookstore-api/models/book_model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllBooks(ctx *gin.Context) {
	var books []models.Book

	database.DB.Find(&books)
	ctx.JSON(http.StatusOK, gin.H{
		"data" : books,
	})
}

func GetBook(ctx *gin.Context) {
	var book models.Book

	if err := database.DB.Where("id = ?", ctx.Param("id")).First(&book).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Data not found!"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data" : book,
	})
}

type CreateBookInput struct {
  Title  string `json:"title" binding:"required"`
  Author string `json:"author" binding:"required"`
	Price  int 		`json:"price" binding:"required"`
}

func CreateBook(ctx *gin.Context) {
	var input CreateBookInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
    ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

	book := models.Book{Title: input.Title, Author: input.Author, Price: input.Price}
  database.DB.Create(&book)

  ctx.JSON(http.StatusOK, gin.H{"data": book})
}

type UpdateBookInput struct {
  Title  string `json:"title"`
  Author string `json:"author"`
	Price  int 		`json:"price"`
}

func UpdateBook(ctx *gin.Context) {
	var input UpdateBookInput
	var book models.Book

	if err := database.DB.Where("id = ?", ctx.Param("id")).First(&book).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Data not found!"})
		return
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
    ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

	database.DB.Model(&book).Updates(input)
	ctx.JSON(http.StatusOK, gin.H{
		"data" : book,
	})
}

func DeleteBook(ctx *gin.Context) {
  // Get model if exist
  var book models.Book
  if err := database.DB.Where("id = ?", ctx.Param("id")).First(&book).Error; err != nil {
    ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
    return
  }

  database.DB.Delete(&book)

  ctx.JSON(http.StatusOK, gin.H{"data": true})
}
package bootstrap

import (
	"bookstore-api/database"
	"bookstore-api/router"

	"github.com/gin-gonic/gin"
)

func BootStrap() {
	app := gin.Default()
	database.ConnectDB()
	router.InitRoute(app)
	app.Run(":8080")
}
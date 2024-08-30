package router

import (
	"80HW/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/items", handlers.GetItems)
	r.GET("/items/:id", handlers.GetItem)
	r.POST("/items", handlers.CreateItem)
	r.PUT("/items/:id", handlers.UpdateItem)
	r.DELETE("/items/:id", handlers.DeleteItem)

	return r
}

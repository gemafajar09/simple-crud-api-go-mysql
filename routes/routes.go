package routes

import (
	"api-go-mysql/controllers"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})
	r.GET("/user", controllers.GetUser)
	r.POST("/user", controllers.TambahUser)
	r.GET("/user/:id", controllers.CariUser)
	r.PATCH("/user/:id", controllers.UpdateUser)
	r.DELETE("user/:id", controllers.DeleteUser)
	return r
}

package routes

import (
	"pasta-diary2-backend/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
    api := router.Group("/api")
    {
        api.GET("/posts", controllers.GetPosts)
        api.POST("/posts", controllers.CreatePost)
        api.PUT("/posts/:id", controllers.UpdatePost)
        api.DELETE("/posts/:id", controllers.DeletePost)
    }
}

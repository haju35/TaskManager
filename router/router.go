package router

import (
    "github.com/gin-gonic/gin"
    "Task_manager/controllers"
)

// SetupRouter initializes all routes and returns a Gin engine.
func SetupRouter() *gin.Engine {
    r := gin.Default()

    api := r.Group("/api")
    controllers.RegisterTaskRoutes(api)

    return r
}

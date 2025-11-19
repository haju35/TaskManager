package router

import (
    "github.com/gin-gonic/gin"
    "Task_manager -API/controllers"
    "Task_manager/data"

)

func SetupRouter() *gin.Engine {
    r := gin.Default()

    api := r.Group("/tasks")
    {
        api.POST("", controllers.CreateTaskHandler)
        api.GET("", controllers.GetAllTasksHandler)
        api.GET("/:id", controllers.GetTaskHandler)
        api.PUT("/:id", controllers.UpdateTaskHandler)
        api.DELETE("/:id", controllers.DeleteTaskHandler)
    }

    return r
}

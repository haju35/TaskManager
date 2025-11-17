package controllers

import (
"net/http"
"github.com/gin-gonic/gin"
"Task_manager/data"
"Task_manager/models"
)

func RegisterTaskRoutes(rg *gin.RouterGroup) {
rg.GET("/tasks", getAll)
rg.GET("/tasks/:id", getByID)
rg.POST("/tasks", createTask)
rg.PUT("/tasks/:id", updateTask)
rg.DELETE("/tasks/:id", deleteTask)
}


func getAll(c *gin.Context) {
all := data.GetAll()
c.JSON(http.StatusOK, gin.H{"tasks": all})
}


func getByID(c *gin.Context) {
id := c.Param("id")
t, err := data.GetByID(id)
if err != nil {
c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
return
}
c.JSON(http.StatusOK, t)
}


func createTask(c *gin.Context) {
var in models.Task
if err := c.ShouldBindJSON(&in); err != nil {
c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
return
}
// basic validation
if in.Title == "" {
c.JSON(http.StatusBadRequest, gin.H{"error": "title is required"})
return
}
if in.Status == "" {
in.Status = "pending"
}
created := data.Create(in)
c.JSON(http.StatusCreated, created)
}


func updateTask(c *gin.Context) {
id := c.Param("id")
var in models.Task
if err := c.ShouldBindJSON(&in); err != nil {
c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
return
}
if err := data.Update(id, in); err != nil {
c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
return
}
c.JSON(http.StatusOK, gin.H{"message": "Task updated"})
}


func deleteTask(c *gin.Context) {
id := c.Param("id")
if err := data.Delete(id); err != nil {
c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
return
}
c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
}
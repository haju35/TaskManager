package controllers

import (
    "context"
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/haju35/Task_manager -API/data"
    "github.com/haju35/Task_manager -API/models"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
)

// CreateTaskHandler handles POST /tasks
func CreateTaskHandler(c *gin.Context) {
    var payload models.Task
    if err := c.ShouldBindJSON(&payload); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request payload", "details": err.Error()})
        return
    }

    // default Completed false if zero value is fine
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    created, err := data.CreateTask(ctx, &payload)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create task", "details": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, created)
}

// GetAllTasksHandler handles GET /tasks
func GetAllTasksHandler(c *gin.Context) {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    tasks, err := data.GetTasks(ctx)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch tasks", "details": err.Error()})
        return
    }
    c.JSON(http.StatusOK, tasks)
}

// GetTaskHandler handles GET /tasks/:id
func GetTaskHandler(c *gin.Context) {
    id := c.Param("id")
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    task, err := data.GetTaskByID(ctx, id)
    if err != nil {
        if err == mongo.ErrNoDocuments {
            c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
            return
        }
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch task", "details": err.Error()})
        return
    }
    c.JSON(http.StatusOK, task)
}

// UpdateTaskHandler handles PUT /tasks/:id
func UpdateTaskHandler(c *gin.Context) {
    id := c.Param("id")

    var payload map[string]interface{} // flexible partial update
    if err := c.ShouldBindJSON(&payload); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request payload", "details": err.Error()})
        return
    }

    // sanitize allowed fields
    allowed := bson.M{}
    if v, ok := payload["title"].(string); ok {
        allowed["title"] = v
    }
    if v, ok := payload["description"].(string); ok {
        allowed["description"] = v
    }
    if v, ok := payload["completed"].(bool); ok {
        allowed["completed"] = v
    }

    if len(allowed) == 0 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "no valid fields provided for update"})
        return
    }

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    updated, err := data.UpdateTask(ctx, id, allowed)
    if err != nil {
        if err == mongo.ErrNoDocuments {
            c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
            return
        }
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update task", "details": err.Error()})
        return
    }
    c.JSON(http.StatusOK, updated)
}

// DeleteTaskHandler handles DELETE /tasks/:id
func DeleteTaskHandler(c *gin.Context) {
    id := c.Param("id")
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    err := data.DeleteTask(ctx, id)
    if err != nil {
        if err == mongo.ErrNoDocuments {
            c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
            return
        }
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete task", "details": err.Error()})
        return
    }
    c.JSON(http.StatusNoContent, gin.H{})
}

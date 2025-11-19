package main

import (
    "context"
    "log"
    "os"
    "os/signal"
    "syscall"
    "time"

    "github.com/yourusername/task_manager/data"
    "Task_manager -API/router"
)

func main() {
    // Load configuration from env vars
    mongoURI := os.Getenv("MONGO_URI")
    if mongoURI == "" {
        // fallback to local default for convenience, but recommend setting env var in production
        mongoURI = "mongodb://localhost:27017"
    }
    dbName := os.Getenv("MONGO_DB")
    if dbName == "" {
        dbName = "taskdb"
    }
    collName := os.Getenv("MONGO_COLLECTION")
    if collName == "" {
        collName = "tasks"
    }
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    // init mongo
    log.Printf("Connecting to MongoDB: %s (db=%s, collection=%s)\n", mongoURI, dbName, collName)
    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()
    if err := data.InitMongo(ctx, mongoURI, dbName, collName); err != nil {
        log.Fatalf("Failed to connect to MongoDB: %v\n", err)
    }
    defer func() {
        // graceful disconnect
        if err := data.DisconnectMongo(context.Background()); err != nil {
            log.Printf("Error disconnecting MongoDB: %v\n", err)
        }
    }()

    r := router.SetupRouter()

    // run server in goroutine
    srvAddr := ":" + port
    go func() {
        if err := r.Run(srvAddr); err != nil {
            log.Fatalf("Failed to run server: %v\n", err)
        }
    }()
    log.Printf("Server running on %s\n", srvAddr)

    // wait for termination signals for graceful shutdown
    quit := make(chan os.Signal, 1)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    <-quit
    log.Println("Shutting down server...")

    // optional: give some time for graceful shutdown
    timeoutCtx, timeoutCancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer timeoutCancel()
    _ = timeoutCtx
    // since gin.Run uses http.ListenAndServe, a full graceful shutdown needs extra wiring if desired.
    // For now we ensure the app disconnects from mongo in defer and exit.
}

package main

import "Task_manager/router"

func main() {
r := router.SetupRouter()
r.Run(":8080")
}
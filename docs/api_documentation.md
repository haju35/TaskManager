# Task Manager API Documentation

Base URL: `http://localhost:8080/api`

## Endpoints

1. GET /tasks

Description: Retrieve all tasks.

Request:

GET /tasks

Response: 200 OK

{
  "tasks": [
    {
      "id": "1a2b3c4d-5678-9101-1121-314151617181",
      "title": "Finish project",
      "description": "Complete the Go Task API",
      "due_date": "2025-11-20",
      "status": "pending"
    }
  ]
}
2. GET /tasks/:id

Description: Retrieve details of a specific task by ID.

Request:

GET /tasks/{id}

Response (200 OK):

{
  "id": "1a2b3c4d-5678-9101-1121-314151617181",
  "title": "Finish project",
  "description": "Complete the Go Task API",
  "due_date": "2025-11-20",
  "status": "pending"
}

Response (404 Not Found):

{
  "error": "Task not found"
}
3. POST /tasks

Description: Create a new task.

Request:

POST /tasks
Content-Type: application/json

Body Example:

{
  "title": "New task title",
  "description": "Task description",
  "due_date": "2025-11-22",
  "status": "pending"
}

Response (201 Created):

{
  "id": "generated-uuid",
  "title": "New task title",
  "description": "Task description",
  "due_date": "2025-11-22",
  "status": "pending"
}

Validation Errors (400 Bad Request):

{
  "error": "title is required"
}
4. PUT /tasks/:id

Description: Update an existing task (partial updates allowed).

Request:

PUT /tasks/{id}
Content-Type: application/json

Body Example:

{
  "title": "Updated task title",
  "status": "in progress"
}

Response (200 OK):

{
  "message": "Task updated"
}

Response (404 Not Found):

{
  "error": "Task not found"
}
5. DELETE /tasks/:id

Description: Delete a task by ID.

Request:

DELETE /tasks/{id}

Response (200 OK):

{
  "message": "Task deleted"
}

Response (404 Not Found):

{
  "error": "Task not found"
}

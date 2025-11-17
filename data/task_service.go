package data

import (
    "errors"
    "sync"

    "github.com/google/uuid"
    "Task_manager/models"
)

var (
    tasks []models.Task
    mu    sync.RWMutex
)

// GetAll returns all tasks.
func GetAll() []models.Task {
    mu.RLock()
    defer mu.RUnlock()
    out := make([]models.Task, len(tasks))
    copy(out, tasks)
    return out
}

// GetByID returns a task by ID.
func GetByID(id string) (*models.Task, error) {
    mu.RLock()
    defer mu.RUnlock()

    for _, t := range tasks {
        if t.ID == id {
            copy := t
            return &copy, nil
        }
    }
    return nil, errors.New("not found")
}

// Create adds a new task.
func Create(t models.Task) models.Task {
    mu.Lock()
    defer mu.Unlock()

    if t.ID == "" {
        t.ID = uuid.New().String()   // <-- requires github.com/google/uuid
    }

    tasks = append(tasks, t)
    return t
}

// Update modifies a task.
func Update(id string, updated models.Task) error {
    mu.Lock()
    defer mu.Unlock()

    for i := range tasks {
        if tasks[i].ID == id {
            if updated.Title != "" {
                tasks[i].Title = updated.Title
            }
            if updated.Description != "" {
                tasks[i].Description = updated.Description
            }
            if updated.DueDate != "" {
                tasks[i].DueDate = updated.DueDate
            }
            if updated.Status != "" {
                tasks[i].Status = updated.Status
            }
            return nil
        }
    }
    return errors.New("not found")
}

// Delete removes a task.
func Delete(id string) error {
    mu.Lock()
    defer mu.Unlock()

    for i := range tasks {
        if tasks[i].ID == id {
            tasks = append(tasks[:i], tasks[i+1:]...)
            return nil
        }
    }
    return errors.New("not found")
}

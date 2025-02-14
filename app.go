package main

import (
	"context"
	"dmark-test/internal/config"
	"dmark-test/internal/models"
	"dmark-test/internal/repository"
	"dmark-test/internal/service"
	"log"
)

// App struct
type App struct {
	ctx     context.Context
	taskSvc *service.TaskService
}

// NewApp creates a new App application struct
func NewApp() *App {
	// Initialize database
	db, err := config.NewDatabase()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Initialize repository and service
	taskRepo := repository.NewTaskRepository(db)
	taskSvc := service.NewTaskService(taskRepo)

	return &App{
		taskSvc: taskSvc,
	}
}

// startup is called when the app starts
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Frontend-facing methods

func (a *App) CreateTask(input service.CreateTaskInput) error {
	return a.taskSvc.CreateTask(input)
}

func (a *App) GetTasksByStatus() (service.TaskList, error) {
	return a.taskSvc.GetTasksByStatus()
}

func (a *App) ToggleTask(id uint) error {
	return a.taskSvc.ToggleTask(id)
}

func (a *App) DeleteTask(id uint) error {
	return a.taskSvc.DeleteTask(id)
}

func (a *App) GetTasksByPriority() (map[int][]models.Task, error) {
	return a.taskSvc.GetTasksByPriority()
}

func (a *App) GetOverdueTasks() ([]models.Task, error) {
	return a.taskSvc.GetOverdueTasks()
}

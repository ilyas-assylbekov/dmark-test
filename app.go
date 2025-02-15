package main

import (
	"context"
	"dmark-test/internal/config"
	"dmark-test/internal/models"
	"dmark-test/internal/repository"
	"dmark-test/internal/service"
	"log"
)

// App представляет собой основную структуру приложения
// Содержит контекст и сервис для работы с задачами
type App struct {
	ctx     context.Context      // Контекст приложения
	taskSvc *service.TaskService // Сервис для работы с задачами
}

// NewApp создает и инициализирует новый экземпляр приложения
// Выполняет следующие шаги:
// 1. Инициализация подключения к базе данных
// 2. Создание репозитория задач
// 3. Инициализация сервиса задач
func NewApp() *App {
	// Инициализация базы данных
	db, err := config.NewDatabase()
	if err != nil {
		log.Fatalf("Не удалось инициализировать базу данных: %v", err)
	}

	// Инициализация репозитория и сервиса
	taskRepo := repository.NewTaskRepository(db)
	taskSvc := service.NewTaskService(taskRepo)

	return &App{
		taskSvc: taskSvc,
	}
}

// startup вызывается при запуске приложения
// Сохраняет контекст приложения
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Методы для работы с фронтендом

// CreateTask создает новую задачу
// Принимает структуру CreateTaskInput с данными задачи
func (a *App) CreateTask(input service.CreateTaskInput) error {
	return a.taskSvc.CreateTask(input)
}

// GetTasksByStatus возвращает список задач, сгруппированных по статусу
// Возвращает TaskList - специальную структуру для отображения задач
func (a *App) GetTasksByStatus() (service.TaskList, error) {
	return a.taskSvc.GetTasksByStatus()
}

// ToggleTask переключает статус задачи (выполнена/не выполнена)
// Принимает ID задачи
func (a *App) ToggleTask(id uint) error {
	return a.taskSvc.ToggleTask(id)
}

// DeleteTask удаляет задачу по её ID
func (a *App) DeleteTask(id uint) error {
	return a.taskSvc.DeleteTask(id)
}

// GetTasksByPriority возвращает задачи, сгруппированные по приоритету
// Возвращает map[int][]models.Task, где ключ - приоритет
func (a *App) GetTasksByPriority() (map[int][]models.Task, error) {
	return a.taskSvc.GetTasksByPriority()
}

// GetOverdueTasks возвращает список просроченных задач
// Возвращает []models.Task - массив задач
func (a *App) GetOverdueTasks() ([]models.Task, error) {
	return a.taskSvc.GetOverdueTasks()
}

package service

import (
	"dmark-test/internal/models"
	"dmark-test/internal/repository"
	"time"
)

type TaskService struct {
	repo *repository.TaskRepository
}

func NewTaskService(repo *repository.TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

type CreateTaskInput struct {
	Title    string
	Priority int
	DueDate  *time.Time
}

type TaskList struct {
	Active    []models.Task
	Completed []models.Task
}

func (s *TaskService) CreateTask(input CreateTaskInput) error {
	task := &models.Task{
		Title:    input.Title,
		Priority: input.Priority,
		DueDate:  input.DueDate,
	}
	return s.repo.Create(task)
}

func (s *TaskService) GetTasksByStatus() (TaskList, error) {
	tasks, err := s.repo.GetAll()
	if err != nil {
		return TaskList{}, err
	}

	var result TaskList
	for _, task := range tasks {
		if task.Completed {
			result.Completed = append(result.Completed, task)
		} else {
			result.Active = append(result.Active, task)
		}
	}
	return result, nil
}

func (s *TaskService) ToggleTask(id uint) error {
	return s.repo.ToggleComplete(id)
}

func (s *TaskService) DeleteTask(id uint) error {
	return s.repo.Delete(id)
}

// Additional methods for handling priorities and due dates
func (s *TaskService) GetTasksByPriority() (map[int][]models.Task, error) {
	tasks, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	result := make(map[int][]models.Task)
	for _, task := range tasks {
		result[task.Priority] = append(result[task.Priority], task)
	}
	return result, nil
}

func (s *TaskService) GetOverdueTasks() ([]models.Task, error) {
	tasks, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	var overdue []models.Task
	now := time.Now()
	for _, task := range tasks {
		if task.DueDate != nil && task.DueDate.Before(now) && !task.Completed {
			overdue = append(overdue, task)
		}
	}
	return overdue, nil
}

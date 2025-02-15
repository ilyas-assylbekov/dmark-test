package service

import (
	"dmark-test/internal/models"
	"dmark-test/internal/repository"
	"fmt"
	"time"
)

// TaskService предоставляет бизнес-логику для работы с задачами
// repo: репозиторий для взаимодействия с базой данных
type TaskService struct {
	repo *repository.TaskRepository
}

// NewTaskService создает новый экземпляр сервиса задач
// repo: инициализированный репозиторий задач
// returns: указатель на новый экземпляр TaskService
func NewTaskService(repo *repository.TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

// CreateTaskInput определяет структуру входных данных для создания задачи
type CreateTaskInput struct {
	Title    string `json:"title"`             // Заголовок задачи
	Priority int    `json:"priority"`          // Приоритет задачи
	DueDate  string `json:"dueDate,omitempty"` // Срок выполнения (опционально)
}

// TaskList группирует задачи по их статусу
type TaskList struct {
	Active    []models.Task // Активные задачи
	Completed []models.Task // Завершенные задачи
}

// CreateTask создает новую задачу в системе
// input: данные для создания задачи
// returns: ошибка, если создание не удалось
func (s *TaskService) CreateTask(input CreateTaskInput) error {
	var dueDate *time.Time
	if input.DueDate != "" {
		// Парсим строку даты в формате ISO 8601
		parsedTime, err := time.Parse(time.RFC3339, input.DueDate)
		if err != nil {
			return fmt.Errorf("неверный формат даты: %w", err)
		}
		dueDate = &parsedTime
	}

	task := &models.Task{
		Title:    input.Title,
		Priority: input.Priority,
		DueDate:  dueDate,
	}
	return s.repo.Create(task)
}

// GetTasksByStatus возвращает задачи, сгруппированные по статусу выполнения
// returns:
// - TaskList: структура с активными и завершенными задачами
// - error: ошибка при получении данных
func (s *TaskService) GetTasksByStatus() (TaskList, error) {
	tasks, err := s.repo.GetAll()
	if err != nil {
		return TaskList{
			Active:    []models.Task{},
			Completed: []models.Task{},
		}, err
	}

	result := TaskList{
		Active:    []models.Task{},
		Completed: []models.Task{},
	}

	// Распределяем задачи по спискам в зависимости от статуса
	for _, task := range tasks {
		if task.Completed {
			result.Completed = append(result.Completed, task)
		} else {
			result.Active = append(result.Active, task)
		}
	}
	return result, nil
}

// ToggleTask переключает статус выполнения задачи
// id: идентификатор задачи
// returns: ошибка при обновлении статуса
func (s *TaskService) ToggleTask(id uint) error {
	return s.repo.ToggleComplete(id)
}

// DeleteTask удаляет задачу из системы
// id: идентификатор задачи для удаления
// returns: ошибка при удалении
func (s *TaskService) DeleteTask(id uint) error {
	return s.repo.Delete(id)
}

// GetTasksByPriority группирует задачи по уровню приоритета
// returns:
// - map[int][]models.Task: карта задач, где ключ - уровень приоритета
// - error: ошибка при получении данных
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

// GetOverdueTasks возвращает список просроченных активных задач
// returns:
// - []models.Task: список просроченных задач
// - error: ошибка при получении данных
func (s *TaskService) GetOverdueTasks() ([]models.Task, error) {
	tasks, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	var overdue []models.Task
	now := time.Now()
	// Фильтруем задачи, выбирая просроченные и незавершенные
	for _, task := range tasks {
		if task.DueDate != nil && task.DueDate.Before(now) && !task.Completed {
			overdue = append(overdue, task)
		}
	}
	return overdue, nil
}

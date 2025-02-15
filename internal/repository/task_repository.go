package repository

import (
	"dmark-test/internal/models"

	"gorm.io/gorm"
)

// TaskRepository реализует слой доступа к данным для задач
// Использует GORM для взаимодействия с базой данных
type TaskRepository struct {
	db *gorm.DB // Экземпляр подключения к базе данных
}

// NewTaskRepository создает новый экземпляр репозитория задач
// db: инициализированное подключение к базе данных
// returns: указатель на новый экземпляр TaskRepository
func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

// Create сохраняет новую задачу в базе данных
// task: указатель на объект задачи для сохранения
// returns: ошибка, если сохранение не удалось
func (r *TaskRepository) Create(task *models.Task) error {
	return r.db.Create(task).Error
}

// GetAll получает все задачи из базы данных
// returns:
// - []models.Task: список всех задач
// - error: ошибка при получении данных
func (r *TaskRepository) GetAll() ([]models.Task, error) {
	var tasks []models.Task
	err := r.db.Find(&tasks).Error
	return tasks, err
}

// Update обновляет существующую задачу в базе данных
// task: указатель на объект задачи с обновленными данными
// returns: ошибка, если обновление не удалось
func (r *TaskRepository) Update(task *models.Task) error {
	return r.db.Save(task).Error
}

// Delete удаляет задачу из базы данных по идентификатору
// id: идентификатор задачи для удаления
// returns: ошибка, если удаление не удалось
func (r *TaskRepository) Delete(id uint) error {
	return r.db.Delete(&models.Task{}, id).Error
}

// ToggleComplete переключает статус выполнения задачи
// id: идентификатор задачи
// returns: ошибка при обновлении статуса
func (r *TaskRepository) ToggleComplete(id uint) error {
	// Получаем текущую задачу из базы данных
	var task models.Task
	err := r.db.First(&task, id).Error
	if err != nil {
		return err
	}

	// Инвертируем статус выполнения и сохраняем изменения
	task.Completed = !task.Completed
	return r.db.Save(&task).Error
}

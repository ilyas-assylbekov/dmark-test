package config

import (
	"dmark-test/internal/models"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// NewDatabase инициализирует и настраивает подключение к базе данных PostgreSQL
// Выполняет следующие шаги:
// 1. Загружает переменные окружения из .env файла
// 2. Формирует строку подключения из переменных окружения
// 3. Устанавливает соединение с базой данных
// 4. Выполняет автоматическую миграцию схемы
//
// Returns:
// - *gorm.DB: настроенное подключение к базе данных
// - error: ошибка в случае проблем с подключением или миграцией
func NewDatabase() (*gorm.DB, error) {
	// Загрузка переменных окружения из .env файла
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("ошибка загрузки .env файла: %w", err)
	}

	// Формирование строки подключения (DSN) из переменных окружения
	// Требуемые переменные окружения:
	// - DB_HOST: хост базы данных
	// - DB_USER: имя пользователя
	// - DB_PASSWORD: пароль
	// - DB_NAME: имя базы данных
	// - DB_PORT: порт
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	// Установка соединения с базой данных
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("не удалось подключиться к базе данных: %w", err)
	}

	// Автоматическая миграция схемы базы данных
	// Создает или обновляет таблицы на основе структуры models.Task
	err = db.AutoMigrate(&models.Task{})
	if err != nil {
		return nil, fmt.Errorf("не удалось выполнить миграцию базы данных: %w", err)
	}

	return db, nil
}

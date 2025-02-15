package main

import (
	"embed"
	"log"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

// Директива go:embed для встраивания фронтенд-ресурсов
// all:frontend/dist указывает на директорию с собранными файлами фронтенда
//
//go:embed all:frontend/dist
var assets embed.FS

// main является точкой входа в приложение
// Инициализирует и запускает Wails приложение с настроенными параметрами
func main() {
	// Создание экземпляра структуры приложения
	app := NewApp()

	// Создание и запуск приложения с настройками
	err := wails.Run(&options.App{
		// Основные параметры окна приложения
		Title:  "Task Manager", // Заголовок окна
		Width:  1024,           // Ширина окна в пикселях
		Height: 768,            // Высота окна в пикселях

		// Настройка сервера статических ресурсов
		AssetServer: &assetserver.Options{
			Assets: assets, // Встроенные ресурсы фронтенда
		},

		// Настройка цвета фона (в формате RGBA)
		// R:27, G:38, B:54 - тёмно-синий цвет
		// A:1 - полная непрозрачность
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},

		// Обработчик события запуска приложения
		OnStartup: app.startup,

		// Привязка методов приложения к фронтенду
		// Делает методы app доступными для вызова из JavaScript
		Bind: []interface{}{
			app,
		},
	})

	// Обработка ошибок запуска приложения
	if err != nil {
		log.Fatal("Ошибка:", err.Error())
	}
}

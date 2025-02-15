// Основной интерфейс для взаимодействия с бэкендом через Wails
export interface WailsApp {
    // Создание новой задачи
    CreateTask: (input: CreateTaskInput) => Promise<void>;

    // Получение списка задач, сгруппированных по статусу
    GetTasksByStatus: () => Promise<TaskList>;

    // Переключение статуса задачи (выполнено/не выполнено)
    ToggleTask: (id: number) => Promise<void>;

    // Удаление задачи по ID
    DeleteTask: (id: number) => Promise<void>;

    // Получение задач, сгруппированных по приоритету
    GetTasksByPriority: () => Promise<Record<number, Task[]>>;

    // Получение списка просроченных задач
    GetOverdueTasks: () => Promise<Task[]>;
}

// Интерфейс для создания новой задачи
export interface CreateTaskInput {
    title: string;              // Заголовок задачи
    priority: number;           // Приоритет
    dueDate: string | null;     // Срок выполнения в формате ISO 8601
}

// Интерфейс для группировки задач по статусу
export interface TaskList {
    Active: Task[];             // Активные задачи
    Completed: Task[];          // Завершенные задачи
}

// Расширение глобального объекта Window
// Добавление типизации для API Wails
declare global {
    interface Window {
        go: {
            main: {
                App: WailsApp; // Доступ к методам бэкенда через window.go.main.App
            };
        };
    }
}

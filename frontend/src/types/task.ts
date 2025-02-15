// Интерфейс Task описывает структуру задачи в приложении
export interface Task {
    id: number;                 // Уникальный идентификатор задачи
    title: string;              // Заголовок задачи
    completed: boolean;         // Статус выполнения
    priority: number;           // Приоритет задачи
    dueDate: string | null;     // Срок выполнения (в формате ISO 8601 или null)
    createdAt: string;          // Дата создания (в формате ISO 8601)
    updatedAt: string;          // Дата последнего обновления (в формате ISO 8601)
}

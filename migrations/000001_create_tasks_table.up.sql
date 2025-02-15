-- 000001_create_tasks_table.up.sql

-- Создание основной таблицы задач с необходимыми полями
CREATE TABLE IF NOT EXISTS tasks (
    id SERIAL PRIMARY KEY,                              -- Уникальный идентификатор задачи
    title VARCHAR(255) NOT NULL,                        -- Заголовок задачи (обязательное поле)
    completed BOOLEAN DEFAULT FALSE,                    -- Статус выполнения (по умолчанию: не выполнено)
    priority INTEGER DEFAULT 0,                         -- Приоритет задачи (по умолчанию: 0)
    due_date TIMESTAMP,                                 -- Срок выполнения (опционально)
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,     -- Дата создания записи
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP      -- Дата последнего обновления
);

-- Функция для автоматического обновления поля updated_at
-- Вызывается при каждом обновлении записи в таблице
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;     -- Устанавливаем текущее время
    RETURN NEW;
END;
$$ language 'plpgsql';

-- Создание триггера, который автоматически обновляет поле updated_at
-- при любом изменении записи в таблице tasks
CREATE TRIGGER update_tasks_updated_at
    BEFORE UPDATE ON tasks
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();

-- 000001_create_tasks_table.down.sql

-- Удаление триггера обновления времени
DROP TRIGGER IF EXISTS update_tasks_updated_at ON tasks;

-- Удаление функции обновления времени
DROP FUNCTION IF EXISTS update_updated_at_column();

-- Удаление основной таблицы задач
DROP TABLE IF EXISTS tasks;

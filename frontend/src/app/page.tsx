'use client' // Указание на клиентский компонент

// Импорты компонентов и утилит
import React, { useEffect, useState } from 'react'
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card"
import { Input } from "@/components/ui/input"
import { Button } from "@/components/ui/button"
import { Calendar } from "@/components/ui/calendar"
import { Popover, PopoverContent, PopoverTrigger } from "@/components/ui/popover"
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "@/components/ui/select"
import { CalendarIcon, CheckCircle, Circle, Trash2 } from 'lucide-react'
import { format } from 'date-fns'
import { Task } from '@/types/task'
import type { CreateTaskInput } from '@/types/wails'

export default function Home() {
  // Состояния компонента
  const [tasks, setTasks] = useState<Task[]>([])              // Список задач
  const [newTask, setNewTask] = useState('')                  // Новая задача
  const [selectedDate, setSelectedDate] = useState<Date>()    // Выбранная дата
  const [priority, setPriority] = useState('0')               // Приоритет
  const [loading, setLoading] = useState(true)                // Состояние загрузки

  // Загрузка задач при монтировании компонента
  useEffect(() => {
    loadTasks()
  }, [])

  // Функция загрузки задач с бэкенда
  const loadTasks = async () => {
    try {
      const result = await window.go.main.App.GetTasksByStatus()
      setTasks([
        ...(Array.isArray(result.Active) ? result.Active : []),
        ...(Array.isArray(result.Completed) ? result.Completed : [])
      ])
    } catch (error) {
      console.error('Ошибка загрузки задач:', error)
    } finally {
      setLoading(false)
    }
  }

  // Обработчики действий с задачами
  const handleAddTask = async () => {
    if (!newTask.trim()) return

    try {
      const input: CreateTaskInput = {
        title: newTask,
        priority: parseInt(priority),
        // Форматировать дату в формате ISO 8601
        dueDate: selectedDate ? selectedDate.toISOString() : null
      }
      await window.go.main.App.CreateTask(input)
      setNewTask('')
      setSelectedDate(undefined)
      setPriority('0')
      await loadTasks()
    } catch (error) {
      console.error('Ошибка добавления задачи:', error)
    }
  }

  const handleToggleTask = async (id: number) => {
    try {
      await window.go.main.App.ToggleTask(id)
      await loadTasks()
    } catch (error) {
      console.error('Ошибка выбора задачи:', error)
    }
  }

  const handleDeleteTask = async (id: number) => {
    if (!confirm('Вы уверены, что хотите удалить эту задачу?')) return

    try {
      await window.go.main.App.DeleteTask(id)
      await loadTasks()
    } catch (error) {
      console.error('Ошибка удаления задачи:', error)
    }
  }

  // Отображение загрузки
  if (loading) {
    return (
      <main className="container mx-auto p-4 max-w-4xl">
        <Card>
            <CardContent className="p-8 text-center">
            Загрузка задач...
            </CardContent>
        </Card>
      </main>
    )
  }

  // Фильтрация задач по статусу
  const activeTasks = tasks.filter(task => !task.completed)
  const completedTasks = tasks.filter(task => task.completed)

  // Рендер интерфейса
  return (
    <main className="container mx-auto p-4 max-w-4xl">
      <Card className="mb-6">
        <CardHeader>
          <CardTitle>Добавить Новую Задачу</CardTitle>
        </CardHeader>
        <CardContent className="flex gap-4">
          <Input
            placeholder="Введите название задачи..."
            value={newTask}
            onChange={(e) => setNewTask(e.target.value)}
          />
          <Select value={priority} onValueChange={setPriority}>
            <SelectTrigger className="w-32">
              <SelectValue placeholder="Приоритет" />
            </SelectTrigger>
            <SelectContent>
              <SelectItem value="0">Низкий</SelectItem>
              <SelectItem value="1">Средний</SelectItem>
              <SelectItem value="2">Высокий</SelectItem>
            </SelectContent>
          </Select>
          <Popover>
            <PopoverTrigger asChild>
              <Button variant="outline" className="w-[200px]">
          <CalendarIcon className="mr-2 h-4 w-4" />
          {selectedDate ? format(selectedDate, 'PPP') : 'Выберите дату'}
              </Button>
            </PopoverTrigger>
            <PopoverContent className="w-auto p-0">
              <Calendar
          mode="single"
          selected={selectedDate}
          onSelect={setSelectedDate}
              />
            </PopoverContent>
          </Popover>
          <Button onClick={handleAddTask}>Добавить задачу</Button>
        </CardContent>
      </Card>

      <div className="grid gap-6">
        <Card>
          <CardHeader>
            <CardTitle>Активные Задачи ({activeTasks.length})</CardTitle>
          </CardHeader>
          <CardContent>
            <div className="space-y-4">
              {activeTasks.map(task => (
                <div
                  key={task.id}
                  className="flex items-center justify-between p-4 border rounded-lg"
                >
                  <div className="flex items-center gap-4">
                    <button onClick={() => handleToggleTask(task.id)}>
                      <Circle className="h-6 w-6 text-gray-400" />
                    </button>
                    <div>
                      <p className={`font-medium ${getPriorityColor(task.priority)}`}>
                        {task.title}
                      </p>
                      {task.dueDate && (
                        <p className="text-sm text-gray-500">
                          Due: {format(new Date(task.dueDate), 'PPP')}
                        </p>
                      )}
                    </div>
                  </div>
                  <button onClick={() => handleDeleteTask(task.id)}>
                    <Trash2 className="h-5 w-5 text-red-500" />
                  </button>
                </div>
              ))}
            </div>
          </CardContent>
        </Card>

        {completedTasks.length > 0 && (
          <Card>
            <CardHeader>
              <CardTitle>Завершенные Задачи ({completedTasks.length})</CardTitle>
            </CardHeader>
            <CardContent>
              <div className="space-y-4">
                {completedTasks.map(task => (
                  <div
                    key={task.id}
                    className="flex items-center justify-between p-4 border rounded-lg bg-gray-50"
                  >
                    <div className="flex items-center gap-4">
                      <button onClick={() => handleToggleTask(task.id)}>
                        <CheckCircle className="h-6 w-6 text-green-500" />
                      </button>
                      <div>
                        <p className="font-medium line-through text-gray-500">
                          {task.title}
                        </p>
                        {task.dueDate && (
                          <p className="text-sm text-gray-400 line-through">
                            Due: {format(new Date(task.dueDate), 'PPP')}
                          </p>
                        )}
                      </div>
                    </div>
                    <button onClick={() => handleDeleteTask(task.id)}>
                      <Trash2 className="h-5 w-5 text-red-500" />
                    </button>
                  </div>
                ))}
              </div>
            </CardContent>
          </Card>
        )}
      </div>
    </main>
  )
}

// Вспомогательная функция для определения цвета приоритета
function getPriorityColor(priority: number): string {
  switch (priority) {
    case 2:
      return 'text-red-600'
    case 1:
      return 'text-yellow-600'
    default:
      return 'text-gray-900'
  }
}

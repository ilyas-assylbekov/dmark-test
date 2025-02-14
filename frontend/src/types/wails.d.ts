export interface WailsApp {
    CreateTask: (input: CreateTaskInput) => Promise<void>;
    GetTasksByStatus: () => Promise<TaskList>;
    ToggleTask: (id: number) => Promise<void>;
    DeleteTask: (id: number) => Promise<void>;
    GetTasksByPriority: () => Promise<Record<number, Task[]>>;
    GetOverdueTasks: () => Promise<Task[]>;
}

export interface CreateTaskInput {
    title: string;
    priority: number;
    dueDate: string | null;
}

export interface TaskList {
    Active: Task[];
    Completed: Task[];
}

declare global {
    interface Window {
        go: {
            main: {
                App: WailsApp;
            };
        };
    }
}

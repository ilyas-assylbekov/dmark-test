export interface Task {
    id: number;
    title: string;
    completed: boolean;
    priority: number;
    dueDate: string | null;
    createdAt: string;
    updatedAt: string;
}

export interface Todo {
  id: string;
  title: string;
  description?: string;
  completed: boolean;
  tasks: TodoTask[];
}

export interface TodoTask {
  id: string;
  todoId: string;
  title: string;
  description: string;
  completed: boolean;
}

export interface CreateTodoRequest {
  title: string;
  description?: string;
}

export interface CreateTodoResponse {
  success: boolean;
}

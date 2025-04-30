import type { CreateTodoRequest, CreateTodoResponse, Todo } from '../types/todo';

const API_BASE = 'http://localhost:8081'; // 动态化基础 URL

export const todoApi = {
  async create(data: CreateTodoRequest): Promise<CreateTodoResponse> {
    const response = await fetch(`${API_BASE}/todos`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(data),
    });
    const result = await response.json();
    if (result.code !== 0) {
      throw new Error(result.message || '创建待办事项失败');
    }
    return result.data;
  },

  async list(): Promise<Todo[]> {
    const response = await fetch(`${API_BASE}/todos`);
    const result = await response.json();
    if (result.code !== 0) {
      throw new Error(result.message || '获取待办事项列表失败');
    }
    return result.data;
  },

  get: async (id: string): Promise<Todo> => {
    const response = await fetch(`${API_BASE}/todos/${id}`); // 替换为 API_BASE
    const result = await response.json();
    if (result.code !== 0) {
      throw new Error(result.message || '获取待办事项失败');
    }
    return result.data;
  },

  async update(id: string, data: Partial<Todo>): Promise<void> {
    await fetch(`${API_BASE}/todos/${id}`, { // 替换为 API_BASE
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(data),
    });
  },

  async delete(id: string): Promise<void> {
    await fetch(`${API_BASE}/todos/${id}`, { // 替换为 API_BASE
      method: 'DELETE',
    });
  },

  async addTask(data: { todoId: string; title: string; description: string }): Promise<void> {
    const response = await fetch(`${API_BASE}/todos/task`, { // 替换为 API_BASE
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        Accept: 'application/json',
      },
      body: JSON.stringify(data),
    });
    const result = await response.json();
    if (result.code !== 0) {
      throw new Error(result.message || '添加任务失败');
    }
  },

  async markAsCompleted(data: { taskId: string; todoId: string }): Promise<void> {
    const response = await fetch(`${API_BASE}/todos/completed`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        Accept: 'application/json',
      },
      body: JSON.stringify(data),
    });
    const result = await response.json();
    if (result.code !== 0) {
      throw new Error(result.message || '标记任务完成失败');
    }
  },
};

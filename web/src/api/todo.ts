import type { CreateTodoRequest, CreateTodoResponse, Todo } from '../types/todo';

const API_BASE = 'http://localhost:8081';

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
    const response = await fetch(`http://localhost:8081/todos/${id}`);
    const result = await response.json();
    if (result.code !== 0) {
      throw new Error(result.message || '获取待办事项失败');
    }
    return result.data;
  },

  async update(id: string, data: Partial<Todo>): Promise<void> {
    await fetch(`http://localhost:8081/todos/${id}`, {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(data),
    });
  },

  async delete(id: string): Promise<void> {
    await fetch(`http://localhost:8081/todos/${id}`, {
      method: 'DELETE',
    });
  },

  // async addTask(id: string): Promise<Todo> {
  //   const response = await axios.get<Todo>(`${API_BASE}/todos/${id}`);
  //   return response.data;
  // },

};

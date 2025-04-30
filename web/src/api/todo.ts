import axios from 'axios';
import type { CreateTodoRequest, CreateTodoResponse, Todo } from '../types/todo';

const API_BASE = 'http://localhost:8081';

export const todoApi = {
  async create(data: CreateTodoRequest): Promise<CreateTodoResponse> {
    const response = await axios.post<CreateTodoResponse>(`${API_BASE}/todos`, data);
    return response.data;
  },

  async list(): Promise<Todo[]> {
    const response = await axios.get<Todo[]>(`${API_BASE}/todos`);
    return response.data;
  },

  async get(id: string): Promise<Todo> {
    const response = await axios.get<Todo>(`${API_BASE}/todos/${id}`);
    return response.data;
  },

  async update(id: string, data: Partial<Todo>): Promise<void> {
    await axios.put(`${API_BASE}/todos/${id}`, data);
  },

  async delete(id: string): Promise<void> {
    await axios.delete(`${API_BASE}/todos/${id}`);
  },
}

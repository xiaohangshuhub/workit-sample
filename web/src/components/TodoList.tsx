import { useState } from 'react';
import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query';
import { Input, Button, Card, Space, Badge, Typography, List, Empty, Tag, message as antMessage } from 'antd';
import { PlusOutlined } from '@ant-design/icons';
import { todoApi } from '../api/todo';
import type { CreateTodoRequest, Todo } from '../types/todo';

const { Title, Paragraph } = Typography;

export default function TodoList() {
  const [title, setTitle] = useState('');
  const queryClient = useQueryClient();

  const { data: todos = [], isLoading, error } = useQuery({
    queryKey: ['todos'],
    queryFn: todoApi.list,
    initialData: [], // 确保初始值为空数组
    onError: (err) => {
      antMessage.error('加载待办事项失败');
      console.error(err);
    },
  });

  const createMutation = useMutation({
    mutationFn: (data: CreateTodoRequest) => todoApi.create(data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['todos'] });
      setTitle('');
      antMessage.success('创建成功');
    },
    onError: (error: Error) => {
      antMessage.error(error.message || '创建失败');
    },
  });

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    if (!title.trim()) {
      antMessage.warning('标题不能为空');
      return;
    }
    createMutation.mutate({ title });
  };

  return (
    <div style={{ maxWidth: '800px', margin: '0 auto', padding: '20px' }}>
      <Space.Compact style={{ width: '100%', marginBottom: '20px' }}>
        <Input
          value={title}
          onChange={(e) => setTitle(e.target.value)}
          placeholder="添加新的待办事项..."
          onPressEnter={handleSubmit}
        />
        <Button
          type="primary"
          icon={<PlusOutlined />}
          onClick={handleSubmit}
          loading={createMutation.isPending}
        >
          添加
        </Button>
      </Space.Compact>

      <List
        loading={isLoading}
        dataSource={Array.isArray(todos) ? todos : []} // 确保数据源为数组
        locale={{
          emptyText: error ? <Empty description="加载失败" /> : <Empty description="暂无待办事项" />,
        }}
        renderItem={(todo) => <TodoItem key={todo.id} todo={todo} />}
        grid={{ gutter: 16, column: 1 }}
      />
    </div>
  );
}

function TodoItem({ todo }: { todo: Todo }) {
  return (
    <List.Item>
      <Card hoverable>
        <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
          <Title level={4} style={{ margin: 0 }}>{todo.title}</Title>
          <Badge count={todo.tasks.length} overflowCount={99} style={{ backgroundColor: '#52c41a' }}>
            <span style={{ marginRight: '24px' }}>事项</span>
          </Badge>
        </div>
        {todo.description && (
          <Paragraph type="secondary" style={{ marginTop: '8px', marginBottom: 0 }}>
            {todo.description}
          </Paragraph>
        )}
      </Card>
    </List.Item>
  );
}

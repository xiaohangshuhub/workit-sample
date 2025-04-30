import { useState } from 'react';
import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query';
import { Input, Button, Badge, Typography, List, Empty, Row, Col, Modal, Form, message as antMessage, Card, Checkbox } from 'antd';
import { todoApi } from '../api/todo';
import type { CreateTodoRequest, Todo } from '../types/todo';

const { Title } = Typography;

export default function TodoList() {
  const [isTaskModalOpen, setIsTaskModalOpen] = useState(false);
  const [taskForm] = Form.useForm();
  const [isTodoModalOpen, setIsTodoModalOpen] = useState(false);
  const [todoForm] = Form.useForm();
  const [selectedTodo, setSelectedTodo] = useState<Todo | null>(null);
  const queryClient = useQueryClient();

  const { data: todos = [], isLoading, error } = useQuery({
    queryKey: ['todos'],
    queryFn: todoApi.list,
    initialData: [],
  });

  const createTodoMutation = useMutation({
    mutationFn: (data: CreateTodoRequest) => todoApi.create(data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['todos'] });
      todoForm.resetFields();
      setIsTodoModalOpen(false);
      antMessage.success('待办事项创建成功');
    },
    onError: (error: any) => {
      antMessage.error(error?.response?.data?.message || '创建失败');
    },
  });

  const addTaskMutation = useMutation({
    mutationFn: ({ todoId, title, description }: { todoId: string; title: string; description: string }) =>
      todoApi.addTask({ todoId, title, description }), // 使用封装的 API 方法
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['todos'] });
      if (selectedTodo) {
        fetchTodoMutation.mutate(selectedTodo.id); // 重新加载选中的待办事项
      }
      taskForm.resetFields();
      setIsTaskModalOpen(false);
      antMessage.success('任务添加成功');
    },
    onError: (error: any) => {
      antMessage.error(error.message || '添加失败');
    },
  });

  const fetchTodoMutation = useMutation({
    mutationFn: (todoId: string) => todoApi.get(todoId),
    onSuccess: (data) => {
      setSelectedTodo(data);
    },
    onError: (error: any) => {
      antMessage.error(error?.response?.data?.message || '加载待办事项失败');
    },
  });

  const markTaskAsCompletedMutation = useMutation({
    mutationFn: ({ taskId, todoId }: { taskId: string; todoId: string }) =>
      todoApi.markAsCompleted({ taskId, todoId }), // 调用封装的 API 方法
    onSuccess: () => {
      if (selectedTodo) {
        fetchTodoMutation.mutate(selectedTodo.id); // 重新加载选中的待办事项
      }
      antMessage.success('任务已标记为完成');
    },
    onError: (error: any) => {
      antMessage.error(error.message || '标记任务完成失败');
    },
  });

  // 添加标记任务完成的处理函数
  const handleMarkTaskAsCompleted = (taskId: string) => {
    if (!selectedTodo) return;
    markTaskAsCompletedMutation.mutate({ taskId, todoId: selectedTodo.id });
  };

  const handleAddTodo = () => {
    todoForm.validateFields().then((values) => {
      createTodoMutation.mutate({
        title: values.title,
        description: values.description, // 确保传递 description
      });
    });
  };

  const handleAddTask = () => {
    if (!selectedTodo) return;
    taskForm.validateFields().then((values) => {
      addTaskMutation.mutate({
        todoId: selectedTodo.id,
        title: values.title,
        description: values.description, // 确保传递 description
      });
    });
  };

  // 修改点击待办事项的逻辑
  const handleSelectTodo = (todo: Todo) => {
    fetchTodoMutation.mutate(todo.id);
  };

  return (
    <Row style={{ height: '100vh', width: '100%' }} gutter={0} wrap={false}>
      {/* 左侧菜单栏 */}
      <Col flex="250px" style={{ height: '100%', padding: '24px', backgroundColor: '#fff', borderRight: '1px solid #f0f0f0' }}>
        <Title level={4} style={{ marginBottom: 24 }}>待办事项管理</Title>
        <Button
          type="primary"
          style={{
            width: '100%',
            marginBottom: '24px',
            height: '32px', // 调整按钮高度
            fontSize: '14px', // 调整字体大小
          }}
          onClick={() => setIsTodoModalOpen(true)}
        >
          添加待办事项
        </Button>
        <List
          style={{ height: 'calc(100vh - 180px)', overflowY: 'auto' }}
          loading={isLoading}
          dataSource={todos}
          renderItem={(todo) => (
            <List.Item
              onClick={() => handleSelectTodo(todo)} // 调用新的点击逻辑
              style={{
                padding: '12px 16px',
                cursor: 'pointer',
                borderRadius: 6,
                backgroundColor: selectedTodo?.id === todo.id ? '#e6f4ff' : 'transparent',
                border: selectedTodo?.id === todo.id ? '1px solid #1677ff' : '1px solid transparent',
                marginBottom: 8,
              }}
            >
              <div style={{ width: '100%', display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
                <span style={{ fontWeight: 'bold', color: todo.completed ? '#52c41a' : '#000' }}>
                  {todo.title}
                </span>
                <Badge
                  status={todo.completed ? 'success' : 'default'} // 使用 Badge 显示完成状态
                  text={todo.completed ? '已完成' : '未完成'}
                />
              </div>
            </List.Item>
          )}
          locale={{ emptyText: error ? '加载失败' : '暂无待办事项' }}
        />
      </Col>

      {/* 右侧内容区域 */}
      <Col flex="auto" style={{ height: '100%', backgroundColor: '#f5f5f5', padding: '24px', overflow: 'hidden' }}>
        {selectedTodo ? (
          <div style={{ height: '100%', display: 'flex', flexDirection: 'column' }}>
            <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', marginBottom: '24px' }}>
              <Title level={3} style={{ margin: 0 }}>{selectedTodo.title}</Title>
              <Button type="primary" style={{ height: '32px', padding: '0 16px', fontSize: '14px' }} onClick={() => setIsTaskModalOpen(true)}>
                添加任务
              </Button>
            </div>
            <div style={{
  flex: 1,
  overflowY: 'auto',
  display: 'grid',
  gridTemplateColumns: 'repeat(auto-fill, minmax(240px, 1fr))', // 调整卡片宽度
  gap: '12px', // 缩小卡片间距
}}>
  {selectedTodo.tasks.map((task) => (
    <Card
      key={task.id}
      style={{
        borderRadius: '6px',
        boxShadow: '0 1px 4px rgba(0, 0, 0, 0.1)',
        backgroundColor: task.completed ? '#fffaf5' : '#fff', // 已完成任务使用浅灰色背景
        opacity: task.completed ? 1 : 1, // 保持透明度一致
        padding: '12px',
      }}
      title={
        <span
          style={{
            fontWeight: 'bold',
            fontSize: '14px',
            color: task.completed ? '#8c8c8c' : '#000', // 已完成任务使用灰色字体
          }}
        >
          {task.title}
        </span>
      }
      extra={
        <Checkbox
          checked={task.completed}
          disabled={task.completed} // 禁用已完成的任务
          onChange={() => handleMarkTaskAsCompleted(task.id)} // 调用标记完成逻辑
        >
          完成
        </Checkbox>
      }
    >
      <p
        style={{
          color: task.completed ? '#8c8c8c' : '#595959', // 已完成任务使用灰色字体
          fontSize: '12px',
          marginBottom: 0,
        }}
      >
        {task.description || '暂无详情'}
      </p>
    </Card>
  ))}
</div>
          </div>
        ) : (
          <div style={{ height: '100%', width: '100%', display: 'flex', alignItems: 'center', justifyContent: 'center' }}>
            <Empty description="请选择一个待办事项" style={{ transform: 'scale(1.3)' }} />
          </div>
        )}
      </Col>

      {/* 添加待办事项的弹窗 */}
      <Modal title="添加待办事项" visible={isTodoModalOpen} onCancel={() => setIsTodoModalOpen(false)} onOk={handleAddTodo}>
        <Form form={todoForm} layout="vertical">
          <Form.Item
            name="title"
            label="标题"
            rules={[{ required: true, message: '请输入标题' }]}
          >
            <Input placeholder="请输入待办事项标题" />
          </Form.Item>
          <Form.Item
            name="description"
            label="描述"
          >
            <Input.TextArea placeholder="请输入待办事项描述" rows={4} />
          </Form.Item>
        </Form>
      </Modal>

      {/* 添加任务的弹窗 */}
      <Modal title="添加任务" visible={isTaskModalOpen} onCancel={() => setIsTaskModalOpen(false)} onOk={handleAddTask}>
        <Form form={taskForm} layout="vertical">
          <Form.Item
            name="title"
            label="任务标题"
            rules={[{ required: true, message: '请输入任务标题' }]}
          >
            <Input placeholder="请输入任务标题" />
          </Form.Item>
          <Form.Item
            name="description"
            label="任务详情"
          >
            <Input.TextArea placeholder="请输入任务详情" rows={4} />
          </Form.Item>
        </Form>
      </Modal>
    </Row>
  );
}

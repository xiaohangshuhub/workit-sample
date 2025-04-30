import { useState } from 'react';
import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query';
import { Input, Button, Badge, Typography, List, Empty, Row, Col, Modal, Form, message as antMessage } from 'antd';
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
    mutationFn: ({ todoId, title }: { todoId: string; title: string }) =>
      todoApi.addTask(todoId, { title }),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['todos'] });
      taskForm.resetFields();
      setIsTaskModalOpen(false);
      antMessage.success('任务添加成功');
    },
    onError: (error: any) => {
      antMessage.error(error?.response?.data?.message || '添加失败');
    },
  });

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

  return (
    <Row
      style={{ height: '100vh', width: '100%' }}
      gutter={0}
      wrap={false}
    >
      {/* 左侧菜单栏 */}
      <Col
        flex="250px"
        style={{
          height: '100%',
          padding: '24px',
          backgroundColor: '#fff',
          borderRight: '1px solid #f0f0f0'
        }}
      >
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
              onClick={() => setSelectedTodo(todo)}
              style={{
                padding: '12px 16px',
                cursor: 'pointer',
                borderRadius: 6,
                backgroundColor: selectedTodo?.id === todo.id ? '#e6f4ff' : 'transparent',
                border: selectedTodo?.id === todo.id ? '1px solid #1677ff' : '1px solid transparent',
                marginBottom: 8
              }}
            >
              <div style={{ width: '100%', display: 'flex', justifyContent: 'space-between' }}>
                <span>{todo.title}</span>
                <Badge count={todo.tasks.length} size="small" />
              </div>
            </List.Item>
          )}
          locale={{ emptyText: error ? '加载失败' : '暂无待办事项' }}
        />
      </Col>

      {/* 右侧铺满内容区域 */}
      <Col
        flex="auto"
        style={{
          height: '100%',
          backgroundColor: '#fff',
          padding: '24px',
          overflow: 'hidden'
        }}
      >
        {selectedTodo ? (
          <div style={{
            height: '100%',
            display: 'flex',
            flexDirection: 'column'
          }}>
            <div style={{ 
              display: 'flex', 
              justifyContent: 'space-between', 
              alignItems: 'center',
              marginBottom: '24px'
            }}>
              <Title level={3} style={{ margin: 0 }}>{selectedTodo.title}</Title>
              <Button
                type="primary"
                style={{
                  height: '32px', // 调整按钮高度
                  padding: '0 16px', // 调整按钮内边距
                  fontSize: '14px', // 调整字体大小
                }}
                onClick={() => setIsTaskModalOpen(true)}
              >
                添加任务
              </Button>
            </div>
            <div style={{ flex: 1, overflowY: 'auto' }}>
              <List
                dataSource={selectedTodo.tasks}
                renderItem={(task) => (
                  <List.Item style={{ padding: '12px 0', borderBottom: '1px solid #f0f0f0' }}>
                    <span>{task.title}</span>
                  </List.Item>
                )}
                locale={{ emptyText: '暂无任务' }}
              />
            </div>
          </div>
        ) : (
          <div style={{
            height: '100%',
            width: '100%',
            display: 'flex',
            alignItems: 'center',
            justifyContent: 'center'
          }}>
            <Empty description="请选择一个待办事项" style={{ transform: 'scale(1.3)' }} />
          </div>
        )}
      </Col>

      {/* 添加待办事项的弹窗 */}
      <Modal
        title="添加待办事项"
        visible={isTodoModalOpen}
        onCancel={() => setIsTodoModalOpen(false)}
        onOk={handleAddTodo}
        confirmLoading={createTodoMutation.isLoading}
      >
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
      <Modal
        title="添加任务"
        visible={isTaskModalOpen}
        onCancel={() => setIsTaskModalOpen(false)}
        onOk={handleAddTask}
        confirmLoading={addTaskMutation.isLoading}
      >
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

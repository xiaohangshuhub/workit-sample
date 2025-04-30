import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import { Layout, Typography } from 'antd';
import TodoList from './components/TodoList';

const { Header, Content } = Layout;
const { Title } = Typography;
const queryClient = new QueryClient();

function App() {
  return (
    <QueryClientProvider client={queryClient}>
      <Layout style={{ minHeight: '100vh' }}>
        <Header style={{ background: '#fff', padding: '0 20px', borderBottom: '1px solid #f0f0f0' }}>
          <Title level={3} style={{ margin: '16px 0' }}>
            待办事项管理
          </Title>
        </Header>
        <Content style={{ background: '#f5f5f5' }}>
          <TodoList />
        </Content>
      </Layout>
    </QueryClientProvider>
  );
}

export default App;

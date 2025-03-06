import { Outlet } from "react-router-dom";
import { Layout } from 'antd';
const { Header, Content, Footer, Sider } = Layout;

const AppLayout = () => {

    const headerStyle = {
        textAlign: 'center',
        height: '10%',
        backgroundColor: 'rgb(148, 208, 176)'
    };
    const contentStyle = {
        textAlign: 'center',
        height: '80%',
        lineHeight: '120px',
        backgroundColor: 'rgba(208, 234, 201, 0.88)'
    };
    const footerStyle = {
        textAlign: 'center',
        height: '10%',
        backgroundColor: 'rgb(203, 207, 208)'
    };
    const layoutStyle = {
        height: '100vh',
        overflow: 'hidden',
        width: '100%',
        maxWidth: '100%',
    };
  return (
      <>
      <Layout style={layoutStyle}>
      <Header style={headerStyle}>Header</Header>
      <Content style={contentStyle}>
        <Outlet /> {/* This is where the child components will be rendered */}
      </Content>
      <Footer style={footerStyle}>Footer</Footer>
      </Layout>  
      </>
  )

};

export default AppLayout;


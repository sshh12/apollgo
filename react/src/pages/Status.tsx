import React from 'react';
import { PageContainer } from '@ant-design/pro-layout';
import {
  Card,
  Row,
  Col,
  Alert,
  Typography,
  Form,
  Button,
  Input,
  Select,
  InputNumber,
  Steps,
} from 'antd';

const { Step } = Steps;

export default (): React.ReactNode => (
  <PageContainer>
    <Card>
      <Steps progressDot current={2}>
        <Step description=":2000, :1080" title="Listening" />
        <Step description="mixed, socks4" title="Forwarding" />
        <Step description="107.12.53.155" title="Internet" />
      </Steps>
    </Card>
    <br />
    <Row gutter={16}>
      <Col span={8}>
        <StatsCard title="Download" value="5 MB/s" />
      </Col>
      <Col span={8}>
        <StatsCard title="Upload" value="2 MB/s" />
      </Col>
      <Col span={8}>
        <StatsCard title="Latency" value="232 ms" />
      </Col>
    </Row>
  </PageContainer>
);

let StatsCard = ({ title, value }) => {
  return (
    <Card bodyStyle={{ fontWeight: 'bold', fontSize: '1.5rem' }} title={title} bordered={false}>
      {value}
    </Card>
  );
};

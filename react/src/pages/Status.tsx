import React, { useEffect, useState } from 'react';
import { PageContainer } from '@ant-design/pro-layout';
import { Card, Row, Col, Steps, Spin } from 'antd';
import api from '../api';

const { Step } = Steps;

export default (): React.ReactNode => {
  let [cfg, setCfg] = useState(null);
  let [status, setStatus] = useState(null);
  useEffect(() => {
    api.get('/api/config').then((cfg) => {
      cfg.listeners.map((list) => (list._id = Math.random().toString(32)));
      setCfg(cfg);
    });
  }, []);
  useEffect(() => {
    api.get('/api/status').then((status) => setStatus(status));
    let statusInterval = setInterval(() => {
      api.get('/api/status').then((status) => setStatus(status));
    }, 10 * 1000);
    return () => {
      clearInterval(statusInterval);
    };
  }, []);
  if (!cfg || !status) {
    return (
      <PageContainer>
        <Spin />
      </PageContainer>
    );
  }
  console.log(status);
  let listDesc = cfg.listeners.map((list) => list.uri).join(', ');
  let forwDesc = cfg.listeners.map((list) => list.forwarders.join(', ')).join(', ');
  return (
    <PageContainer>
      <Card>
        <Steps progressDot current={2}>
          <Step description={listDesc} title="Listening" />
          <Step description={forwDesc} title="Forwarding" />
          <Step description={status.ip} title="Internet" />
        </Steps>
      </Card>
      <br />
      <Row gutter={16}>
        <Col span={8}>
          <StatsCard title="Download" value={`${Math.round(status.dlSpeed * 100) / 100} MB/s`} />
        </Col>
        <Col span={8}>
          <StatsCard title="Upload" value={`${Math.round(status.ulSpeed * 100) / 100} MB/s`} />
        </Col>
        <Col span={8}>
          <StatsCard title="Latency" value={`${Math.round(status.latency)} ms`} />
        </Col>
      </Row>
    </PageContainer>
  );
};

let StatsCard = ({ title, value }) => {
  return (
    <Card bodyStyle={{ fontWeight: 'bold', fontSize: '1.5rem' }} title={title} bordered={false}>
      {value}
    </Card>
  );
};

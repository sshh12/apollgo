import React, { useEffect, useState } from 'react';
import { PageContainer } from '@ant-design/pro-layout';
import { Card, Alert, Form, Button, Input, Select, InputNumber, Spin } from 'antd';
import { MinusCircleOutlined, PlusOutlined } from '@ant-design/icons';
import api from '../api';

const { Option } = Select;

export default (): React.ReactNode => {
  let [cfg, setCfg] = useState(null);
  useEffect(() => {
    api.get('/api/config').then((cfg) => {
      cfg.listeners.map((list) => (list._id = Math.random().toString(32)));
      setCfg(cfg);
    });
  }, []);
  let addListener = () => {
    let newCfg = JSON.parse(JSON.stringify(cfg));
    let newList = JSON.parse(JSON.stringify(newCfg.listeners[0]));
    newList._id = Math.random().toString(32);
    newCfg.listeners.push(newList);
    setCfg(newCfg);
  };
  let removeListener = (id) => {
    let newCfg = JSON.parse(JSON.stringify(cfg));
    if (newCfg.listeners.length <= 1) {
      return;
    }
    newCfg.listeners = newCfg.listeners.filter(({ _id }) => _id != id);
    setCfg(newCfg);
  };
  if (!cfg) {
    return (
      <PageContainer>
        <Spin />
      </PageContainer>
    );
  }
  return (
    <PageContainer>
      <Alert
        message={
          <a target="_blank" href="https://github.com/nadoo/glider#protocols">
            View supported configuration values and schemas
          </a>
        }
        type="info"
        showIcon
      />
      <br />
      {cfg.listeners.map((list) => {
        return (
          <div key={list._id}>
            <ListenerConfig
              listener={list}
              add={addListener}
              remove={() => removeListener(list._id)}
            />
            <br />
          </div>
        );
      })}
    </PageContainer>
  );
};

const formItemProps = {
  style: { width: '90%' },
  labelCol: {
    xs: { span: 24 },
    sm: { span: 4 },
  },
  wrapperCol: {
    xs: { span: 24 },
    sm: { span: 20 },
  },
};
const formItemPropsWithOutLabel = {
  wrapperCol: {
    xs: { span: 24, offset: 0 },
    sm: { span: 20, offset: 4 },
  },
};

const ListenerConfig = ({ listener, add, remove }) => {
  return (
    <Card>
      <Form requiredMark={false} {...formItemPropsWithOutLabel} onFinish={console.log}>
        <Form.Item label="Listener" required={true} {...formItemProps}>
          <Input
            placeholder="SCHEME://[USER|METHOD:PASSWORD@][HOST]:PORT?PARAMS"
            defaultValue={listener.uri}
          />
        </Form.Item>
        <Form.List name="forwarders">
          {(fields, { add, remove }, { errors }) => (
            <>
              {fields.map((field, index) => (
                <Form.Item
                  {...(index === 0 ? formItemProps : formItemPropsWithOutLabel)}
                  label={index === 0 ? 'Fowarders' : ''}
                  required={true}
                  key={field.key}
                >
                  <Form.Item {...field} noStyle>
                    <Input placeholder="SCHEME://[USER|METHOD:PASSWORD@][HOST]:PORT?PARAMS[,SCHEME://[USER|METHOD:PASSWORD@][HOST]:PORT?PARAMS]" />
                  </Form.Item>
                  {fields.length > 1 ? (
                    <MinusCircleOutlined onClick={() => remove(field.name)} />
                  ) : null}
                </Form.Item>
              ))}
              <Form.Item>
                <Button type="dashed" onClick={() => add()} icon={<PlusOutlined />}>
                  Add forwarder
                </Button>
                <Form.ErrorList errors={errors} />
              </Form.Item>
            </>
          )}
        </Form.List>
        <Form.Item required={true} name="strategy" label="Forward Strategy" {...formItemProps}>
          <Select
            placeholder="Select a forward strategy"
            onChange={console.log}
            allowClear
            defaultValue="rr"
          >
            <Option value="rr">Round Robin</Option>
            <Option value="ha">High Availability</Option>
            <Option value="lha">Latency-Based High Availability</Option>
            <Option value="dh">Destination Hashing</Option>
          </Select>
        </Form.Item>
        <Form.Item label="Connection Check" required={true} {...formItemProps}>
          <Input defaultValue="www.google.com" />
          <InputNumber
            min={1}
            style={{ width: '10rem' }}
            max={99999}
            defaultValue={120}
            formatter={(v) => `every ${v} seconds`}
          />
          <InputNumber
            min={1}
            style={{ width: '12rem' }}
            max={99999}
            defaultValue={120}
            formatter={(v) => `timeout in ${v} seconds`}
          />
        </Form.Item>
        <Form.Item label="DNS" required={true} {...formItemProps}>
          <Select mode="multiple" allowClear defaultValue={['8.8.8.8:53']}>
            <Option value={'8.8.8.8:53'}>8.8.8.8:53</Option>
          </Select>
        </Form.Item>
        <br />
        <Form.Item {...formItemProps}>
          <Button type="primary">Apply All</Button>
          <Button type="default" onClick={remove}>
            Delete
          </Button>
          <Button type="link" onClick={add}>
            Add Listener
          </Button>
        </Form.Item>
      </Form>
    </Card>
  );
};

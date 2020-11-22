import React from 'react';
import { PageContainer } from '@ant-design/pro-layout';
import { Card, Alert, Form, Button, Input, Select, InputNumber } from 'antd';
import { MinusCircleOutlined, PlusOutlined } from '@ant-design/icons';

const { Option } = Select;

export default (): React.ReactNode => (
  <PageContainer>
    <Alert
      message={
        <a target="_blank" href="https://github.com/nadoo/glider">
          View supported configuration values and schemas
        </a>
      }
      type="info"
      showIcon
    />
    <br />
    <InstanceConfig />
  </PageContainer>
);

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

const InstanceConfig = () => {
  return (
    <Card>
      <Form requiredMark={false} {...formItemPropsWithOutLabel} onFinish={console.log}>
        <Form.Item label="Listener" required={true} {...formItemProps}>
          <Input placeholder="SCHEME://[USER|METHOD:PASSWORD@][HOST]:PORT?PARAMS" />
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
          <Button type="primary">Save All</Button>
          <Button type="default">Delete</Button>
          <Button type="link">Add Listener</Button>
        </Form.Item>
      </Form>
    </Card>
  );
};

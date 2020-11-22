export default [
  {
    path: '/status',
    name: 'status',
    icon: 'CloudServerOutlined',
    component: './Status',
  },
  {
    path: '/config',
    name: 'config',
    icon: 'SettingOutlined',
    component: './Config',
  },
  {
    path: '/',
    redirect: '/status',
  },
];

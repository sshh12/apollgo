import './App.css';
import React, { useEffect, useState } from 'react';
import { ThemeProvider } from 'theme-ui';
import theme from './theme';
import { Box, Text, Flex, Link } from 'rebass';
import Glider from './components/Glider';
import Status from './components/Status';
import api from './api';

const TABS = {
  status: Status,
  glider: Glider
};

function getTab() {
  let path = window.location.pathname.replace('/', '');
  if (Object.keys(TABS).includes(path)) {
    return path;
  }
  return 'status';
}

function App() {
  let [tab, setTab] = useState(getTab());
  let [status, setStatus] = useState(null);
  let [config, setConfig] = useState(null);
  let [loading, setLoading] = useState(false);
  useEffect(() => {
    api.get('/api/status').then(setStatus);
    api.get('/api/config').then(setConfig);
  }, []);
  let applyConfig = (newCfg) => {
    setLoading(true);
    setConfig(newCfg);
    setLoading(false);
  };
  let TabView = TABS[tab];
  return (
    <ThemeProvider theme={theme}>
      <div className="App">
        <Flex px={2} color="white" bg="black" alignItems="center">
          <Text p={2} fontWeight="bold">
            apollgo
          </Text>
          {loading && <Text p={2}>*loading*</Text>}
          <Box mx="auto" />
          {Object.keys(TABS).map((tb) => (
            <Link
              key={tb}
              href="#"
              onClick={() => setTab(tb)}
              sx={{
                display: 'inline-block',
                fontWeight: 'bold',
                px: 2,
                py: 1,
                color: 'inherit'
              }}
            >
              {tb}
            </Link>
          ))}
        </Flex>
        <TabView status={status} config={config} setConfig={applyConfig} />
      </div>
    </ThemeProvider>
  );
}

export default App;

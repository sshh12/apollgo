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
  let [tab, _] = useState(getTab());
  let TabView = TABS[tab];
  return (
    <ThemeProvider theme={theme}>
      <div className="App">
        <Flex px={2} color="white" bg="black" alignItems="center">
          <Text p={2} fontWeight="bold">
            Apollgo
          </Text>
          <Box mx="auto" />
          <Link
            sx={{
              display: 'inline-block',
              fontWeight: 'bold',
              px: 2,
              py: 1,
              color: 'inherit'
            }}
          >
            Status
          </Link>
          <Link
            sx={{
              display: 'inline-block',
              fontWeight: 'bold',
              px: 2,
              py: 1,
              color: 'inherit'
            }}
          >
            Glider
          </Link>
        </Flex>
        <TabView />
      </div>
    </ThemeProvider>
  );
}

export default App;

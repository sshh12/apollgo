import React from 'react';
import { Box, Card, Flex, Heading, Text } from 'rebass';

export default function Status({ status }) {
  return (
    <Box>
      <Flex mt={10}>
        <Box width={1 / 2}>
          <Card>
            <Heading>Connection</Heading>
            <Text>IPv4 {status?.ip}</Text>
          </Card>
        </Box>
        <Box width={1 / 2}>
          <Card>
            <Heading>Metrics</Heading>
            <Text>{Math.round(status?.dlSpeed * 100) / 100} Mb/s &darr;</Text>
            <Text>{Math.round(status?.ulSpeed * 100) / 100} Mb/s &uarr;</Text>
            <Text>{status?.latency} ms</Text>
          </Card>
        </Box>
      </Flex>
    </Box>
  );
}

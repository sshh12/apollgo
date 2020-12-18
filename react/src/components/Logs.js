import React from 'react';
import { Box } from 'rebass';

export default function Logs({ status }) {
  let logs = status?.logs || [];
  return (
    <Box>
      <Box ml={4} mt={3} textAlign={'left'}>
        <pre>
          {logs
            .map((l) => `[${new Date(l.time * 1000).toISOString()}] ${l.text}`)
            .join('\n')}
        </pre>
      </Box>
    </Box>
  );
}

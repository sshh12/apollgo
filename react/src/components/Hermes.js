import React, { useState, useEffect } from 'react';
import { Box, Card, Heading, Flex, Button, Text } from 'rebass';
import { Label, Input, Select, Checkbox, Switch } from '@rebass/forms';
import TextInput from './TextInput';
import CSVInput from './CSVInput';

export default function Hermes({ config, setConfig }) {
  let [editCfg, setEditCfg] = useState(config);
  useEffect(() => {
    setEditCfg(config);
  }, [config]);
  let cfgChanged = JSON.stringify(config) != JSON.stringify(editCfg);
  return (
    <Box>
      <Box>
        <a href="https://github.com/sshh12/hermes">
          <p>See hermes docs</p>
        </a>
        {editCfg && (
          <Switch
            checked={editCfg.enableHermes}
            onClick={(e) =>
              setEditCfg({ ...editCfg, enableHermes: !editCfg.enableHermes })
            }
          />
        )}
      </Box>
      {editCfg && (
        <Box ml={4} mr={4}>
          <Flex>
            <Box width={7 / 8}>
              <TextInput
                placeholder="hermes.example.com"
                key={'hermeshost'}
                value={editCfg.hermesConfig.server}
                label={'use hermes server'}
                onChange={(v) =>
                  setEditCfg({
                    ...editCfg,
                    hermesConfig: { ...editCfg.hermesConfig, server: v }
                  })
                }
              />
            </Box>
            <Box ml={2} width={1 / 8}>
              <TextInput
                type={'number'}
                placeholder="4000"
                key={'check'}
                value={editCfg.hermesConfig.port}
                label={'on port'}
                onChange={(v) =>
                  setEditCfg({
                    ...editCfg,
                    hermesConfig: { ...editCfg.hermesConfig, port: v }
                  })
                }
              />
            </Box>
          </Flex>
          <TextInput
            placeholder="password"
            key={'hermespass'}
            value={editCfg.hermesConfig.password}
            label={'with password (if used by server)'}
            onChange={(v) =>
              setEditCfg({
                ...editCfg,
                hermesConfig: { ...editCfg.hermesConfig, password: v }
              })
            }
          />
          <CSVInput
            placeholder="localPort/remotePort[,localPort/remotePort]"
            key={'hermespairs'}
            label={'forward TCP/HTTP on these ports (as local/remote)'}
            value={editCfg.hermesConfig.forwardPairs}
            validate={(v) => !!v.match(/^\d+\/\d+$/)}
            onChange={(v) =>
              setEditCfg({
                ...editCfg,
                hermesConfig: { ...editCfg.hermesConfig, forwardPairs: v }
              })
            }
          />
        </Box>
      )}
      <Box mt={4}>
        {cfgChanged && (
          <Button onClick={() => setEditCfg(config)} ml={2}>
            Reset
          </Button>
        )}
        {cfgChanged && (
          <Button onClick={() => setConfig(editCfg)} bg={'secondary'} ml={2}>
            Apply
          </Button>
        )}
      </Box>
    </Box>
  );
}

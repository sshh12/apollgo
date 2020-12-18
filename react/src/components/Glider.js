import React, { useState, useEffect } from 'react';
import { Box, Card, Heading, Flex, Button, Text } from 'rebass';
import { Label, Select, Checkbox, Switch } from '@rebass/forms';
import TextInput from './TextInput';
import CSVInput from './CSVInput';

export default function Glider({ config, setConfig, defaults }) {
  let [editCfg, setEditCfg] = useState(config);
  useEffect(() => {
    setEditCfg(config);
  }, [config]);
  let editListener = (idx, newList) => {
    let newCfg = JSON.parse(JSON.stringify(editCfg));
    newCfg.listeners = Object.assign([], newCfg.listeners, { [idx]: newList });
    setEditCfg(newCfg);
  };
  let addListener = () => {
    let newCfg = JSON.parse(JSON.stringify(editCfg));
    newCfg.listeners = newCfg.listeners.concat([defaults.listeners[0]]);
    setEditCfg(newCfg);
  };
  let deleteListener = (idx) => {
    if (editCfg.listeners.length == 1) {
      return;
    }
    let newCfg = JSON.parse(JSON.stringify(editCfg));
    newCfg.listeners = newCfg.listeners.filter((_, i) => i != idx);
    setEditCfg(newCfg);
  };
  let cfgChanged = JSON.stringify(config) != JSON.stringify(editCfg);
  return (
    <Box>
      <Box>
        <a href="https://github.com/nadoo/glider#protocols">
          <p>See glider docs and schemes</p>
        </a>
        {editCfg && (
          <Switch
            checked={editCfg.enableGlider}
            onClick={(e) =>
              setEditCfg({ ...editCfg, enableGlider: !editCfg.enableGlider })
            }
          />
        )}
      </Box>
      {editCfg && (
        <Box mt={10} ml={4} mr={4} textAlign={'left'}>
          {editCfg.listeners.map((list, listIdx) => (
            <>
              <ListenerSettingsCard
                list={list}
                listIdx={listIdx}
                edit={(newList) => editListener(listIdx, newList)}
                del={() => deleteListener(listIdx)}
              />
              <hr />
            </>
          ))}
        </Box>
      )}
      <Box mt={4}>
        <Button onClick={addListener}>Add Glider Instance</Button>
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

function ListenerSettingsCard({ list, listIdx, edit, del }) {
  let editURI = (uriIdx, newURI) => {
    let newURIs = Object.assign([], list.uris, {
      [uriIdx]: newURI
    });
    while (newURIs[newURIs.length - 1] == '') {
      newURIs.pop();
    }
    newURIs.push('');
    edit({
      ...list,
      uris: newURIs
    });
  };
  let editFoward = (forwardIdx, newForward) => {
    let newForwards = Object.assign([], list.forwarders, {
      [forwardIdx]: newForward
    });
    while (newForwards[newForwards.length - 1] == '') {
      newForwards.pop();
    }
    if (newForwards.length > 0) newForwards.push('');
    edit({
      ...list,
      forwarders: newForwards
    });
  };
  let hasForward = list.forwarders.length > 0;
  return (
    <Card key={listIdx} width={1} mb={3}>
      <Heading>
        {list.uris?.filter((v) => !!v).join(', ') || 'Listener ' + listIdx}
      </Heading>
      <>
        {list.uris.map((uri, uriIdx) => (
          <TextInput
            placeholder="SCHEME://[USER|METHOD:PASSWORD@][HOST]:PORT?PARAMS"
            key={'uri' + uriIdx}
            value={uri}
            label={uri || uriIdx == 0 ? 'listen on' : '(and listen on)'}
            onChange={(v) => editURI(uriIdx, v)}
          />
        ))}
      </>
      {!hasForward ? (
        <>
          <TextInput
            placeholder="SCHEME://[USER|METHOD:PASSWORD@][HOST]:PORT?PARAMS[,SCHEME..."
            key={'forwardTemp'}
            value={''}
            label={'(forward to)'}
            onChange={(v) => {
              edit({ ...list, forwarders: [v, ''] });
            }}
          />
        </>
      ) : (
        <>
          {list.forwarders.map((forward, forwardIdx) => (
            <TextInput
              placeholder="SCHEME://[USER|METHOD:PASSWORD@][HOST]:PORT?PARAMS[,SCHEME..."
              key={'forward' + forwardIdx}
              value={forward}
              label={forward ? 'forward to' : '(forward to)'}
              onChange={(v) => editFoward(forwardIdx, v)}
            />
          ))}
        </>
      )}
      {hasForward && (
        <Box mt={2}>
          <Label htmlFor={'strat'}>with forward strategy</Label>
          <Select
            value={list.strategy}
            id="strat"
            name="strat"
            onChange={(v) => edit({ ...list, strategy: v })}
          >
            <option value={'rr'}>round robin</option>
            <option value={'ha'}>high availability</option>
            <option value={'lha'}>latency based high availability</option>
            <option value={'dh'}>destination hashing</option>
          </Select>
        </Box>
      )}
      <Flex>
        <Box width={7 / 8}>
          <TextInput
            placeholder="http://www.msftconnecttest.com/connecttest.txt#expect=200"
            key={'check'}
            value={list.check}
            label={'check internet connection by polling'}
            onChange={(v) => edit({ ...list, check: v })}
          />
        </Box>
        <Box ml={2} width={1 / 8}>
          <TextInput
            type={'number'}
            placeholder="300"
            key={'check'}
            value={list.checkInterval}
            label={'every X seconds'}
            onChange={(v) => edit({ ...list, checkInterval: v })}
          />
        </Box>
      </Flex>
      <Box mt={2}>
        <Label>
          <Checkbox
            id="dodns"
            name="dodns"
            checked={list.dns != ''}
            onChange={(e) => {
              if (e.target.checked) edit({ ...list, dns: ':53' });
              else edit({ ...list, dns: '' });
            }}
          />
          Custom DNS
        </Label>
      </Box>
      {list.dns != '' && (
        <Box ml={3} mr={3}>
          <TextInput
            placeholder="[host]:port"
            key={'dns'}
            label={'listen for dns on'}
            value={list.dns}
            onChange={(v) => {
              edit({ ...list, dns: v });
            }}
          />
          <CSVInput
            placeholder="host:port,[host:port]"
            key={'dnsremote'}
            label={'use these remote dns servers'}
            value={list.dnsServers}
            validate={(v) => !!v.match(/^\d+\.\d+\.\d+\.\d+:\d+$/)}
            onChange={(v) => {
              edit({ ...list, dnsServers: v });
            }}
          />
          <CSVInput
            placeholder="domain/ip,[domain/ip]"
            key={'dnsrecords'}
            label={'include these records'}
            value={list.dnsRecords}
            validate={(v) => !!v.match(/^[\w\.]+\/[A-Za-z0-9\.\.]+$/)}
            onChange={(v) => {
              edit({ ...list, dnsRecords: v });
            }}
          />
        </Box>
      )}
      <Text mt={3} color="secondary" onClick={del}>
        [delete]
      </Text>
    </Card>
  );
}

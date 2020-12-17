import React, { useState, useEffect } from 'react';
import { Box, Card, Heading, Flex, Button, Text } from 'rebass';
import { Label, Input, Select } from '@rebass/forms';

export default function Glider({ config, setConfig }) {
  let [editCfg, setEditCfg] = useState(null);
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
    newCfg.listeners = newCfg.listeners.concat([newCfg.listeners[0]]);
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
      </Box>
      {editCfg && (
        <Box mt={10} ml={4} mr={4} textAlign={'left'}>
          {editCfg.listeners.map((list, listIdx) => (
            <>
              <ListenerSettingsCard
                key={listIdx}
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
        <Button onClick={addListener}>Add Listener</Button>
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
      <Heading>{list.uri || 'Listener ' + listIdx}</Heading>
      <TextInput
        placeholder="SCHEME://[USER|METHOD:PASSWORD@][HOST]:PORT?PARAMS"
        key={'uri'}
        label={'listen on'}
        value={list.uri}
        onChange={(v) => {
          edit({ ...list, uri: v });
        }}
      />
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
          <Select id="strat" name="strat">
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
      <Text mt={3} color="secondary" onClick={del}>
        [delete]
      </Text>
    </Card>
  );
}

function TextInput({ value, placeholder, label, key, onChange, type }) {
  return (
    <Box mt={2}>
      <Label htmlFor={key}>{label}</Label>
      <Input
        id={key}
        name={key}
        type={type || 'text'}
        value={value}
        onChange={(e) => onChange(e.target.value)}
        placeholder={placeholder}
      />
    </Box>
  );
}

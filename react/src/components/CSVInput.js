import React, { useState } from 'react';
import { Text } from 'rebass';
import TextInput from './TextInput';

export default function CSVInput({
  value,
  placeholder,
  label,
  key,
  onChange,
  validate
}) {
  let [text, setText] = useState(value.join(', '));
  let [values, setValues] = useState(value);
  return (
    <>
      <TextInput
        placeholder={placeholder}
        key={key}
        label={label}
        value={text}
        onChange={(v) => {
          setText(v);
          let vals = v
            .split(',')
            .map((item) => item.trim())
            .filter(validate);
          vals = [...new Set(vals)];
          setValues(vals);
          onChange(vals);
        }}
      />
      <Text fontSize={2} fontWeight="bold" color="primary">
        {values.map((v) => `(${v})`).join(',')}
      </Text>
    </>
  );
}

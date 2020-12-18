import React from 'react';
import { Box } from 'rebass';
import { Label, Input } from '@rebass/forms';

export default function TextInput({
  value,
  placeholder,
  label,
  key,
  onChange,
  type
}) {
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

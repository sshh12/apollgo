const glob = require('glob');
const fs = require('fs');
const path = require('path');

glob('../react/build/**/*.*', {}, (err, files) => {
  if (err) {
    console.warn(err);
    return;
  }
  files.map((fn) => {
    let newFn = path.join(
      '../cmd/android/assets',
      fn.replace(/[\\/]/g, '__').replace('..__react__build__', '')
    );
    console.log(newFn);
    fs.mkdir(path.dirname(newFn), { recursive: true }, (err) => {
      if (err) {
        console.warn(err);
        return;
      }
      fs.rename(fn, newFn, (err) => {
        if (err) console.warn(err);
      });
    });
  });
});

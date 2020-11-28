const glob = require('glob');
const fs = require('fs');
const path = require('path');

glob('../react/dist/**/*.*', {}, (err, files) => {
  if (err) {
    console.warn(err);
    return;
  }
  files.map((fn) => {
    let newFn = path.join(
      '../assets',
      fn.replace('/', '__').replace('\\', '__').replace('..__react/dist/', ''),
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

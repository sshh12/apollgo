const BASE_URL =
  window.location.origin == 'http://localhost:8000'
    ? 'http://localhost:8888'
    : window.location.origin;

let get = async (path, data = null) => {
  if (data) {
    path +=
      '?' +
      Object.entries(data)
        .map((kv) => kv.map(encodeURIComponent).join('='))
        .join('&');
  }
  let resp = await fetch(BASE_URL + path);
  return await resp.json();
};

let post = async (path, data = {}) => {
  let resp = await fetch(BASE_URL + path, {
    method: 'POST',
    body: JSON.stringify(data),
    headers: {
      'Content-Type': 'application/json',
    },
  });
  return await resp.json();
};

let put = async (path, data = {}) => {
  let resp = await fetch(BASE_URL + path, {
    method: 'PUT',
    body: JSON.stringify(data),
    headers: {
      'Content-Type': 'application/json',
    },
  });
  return await resp.json();
};

export default { get, post, put };

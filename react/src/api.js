const BASE_URL =
  window.location.origin == 'http://localhost:3000'
    ? 'http://localhost:5000'
    : window.location.origin;

let get = async (path) => {
  let resp = await fetch(BASE_URL + path);
  return await resp.json();
};

let put = async (path, data) => {
  let resp = await fetch(BASE_URL + path, {
    method: 'PUT',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(data)
  });
  return await resp.json();
};

export default { get, put };

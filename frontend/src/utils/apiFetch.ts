export const apiFetch = async (url, method = 'GET', data = null) => {
  const res = await fetch(url, {
    method,
    headers: { 'Content-Type': 'application/json' },
    body: data ? JSON.stringify(data) : null,
  });

  const json = await res.json();
  if (!res.ok) throw json;
  return json;
};

export const apiFetch = async (url: string, method: string = 'GET', data: any = null) => {
  // Use RequestInit type for fetch options
  const options: RequestInit = {
    method,
    headers: {
      'Content-Type': 'application/json',
    },
  };

  if (data) {
    options.body = JSON.stringify(data); // now TS knows body exists
  }

  const res = await fetch(url, options);

  // Parse JSON safely
  const json = await res.json().catch(() => null);

  if (!res.ok) {
    throw json || { error: 'API request failed' };
  }

  return json;
};

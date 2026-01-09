// src/api/report.js
import { apiFetch } from '../utils/apiFetch.js';

export const getStockReport = (from, to) => {
  const url = `http://localhost:8080/api/stock/report?from=${from}&to=${to}`;
  return apiFetch(url);
};

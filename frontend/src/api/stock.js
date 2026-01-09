// src/api/stock.js
import { apiFetch } from '../utils/apiFetch.js';

const BASE_URL = 'http://localhost:8080/api/stock';

// Payload: { product_id: string, sub_variant_id: string, quantity: string }
export const addStock = (payload) => {
  return apiFetch(`${BASE_URL}/in`, 'POST', payload);
};

export const removeStock = (payload) => {
  return apiFetch(`${BASE_URL}/out`, 'POST', payload);
};

import { apiFetch } from '../utils/apiFetch.js';

const BASE_URL = 'http://localhost:8080/api/products';

// Create new product
export const createProduct = (payload) => {
  return apiFetch(BASE_URL, 'POST', payload);
};

// List products with pagination
export const listProducts = (page = 1, limit = 10) => {
  const url = `${BASE_URL}?page=${page}&limit=${limit}`;
  return apiFetch(url);
};

// Get product by ID
export const getProductById = (id) => {
  return apiFetch(`${BASE_URL}/${id}`);
};

// Update product
export const updateProduct = (id, payload) => {
  return apiFetch(`${BASE_URL}/${id}`, 'PUT', payload);
};

// Delete product
export const deleteProduct = (id) => {
  return apiFetch(`${BASE_URL}/${id}`, 'DELETE');
};

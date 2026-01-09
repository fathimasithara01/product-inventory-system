import { writable } from 'svelte/store';
import { apiFetch } from '../utils/apiFetch.js';

export const products = writable([]);
export const totalProducts = writable(0);

export async function fetchProducts(page = 1, limit = 5) {
  try {
    const res = await apiFetch(`http://localhost:8080/api/products?page=${page}&limit=${limit}`);
    products.set(res.products || []);
    totalProducts.set(res.total || 0);
  } catch (err) {
    console.error(err);
  }
}

export async function addProduct(product) {
  return await apiFetch('http://localhost:8080/api/products', 'POST', product);
}

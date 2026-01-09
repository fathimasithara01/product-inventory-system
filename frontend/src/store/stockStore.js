// src/store/stockStore.js
import { writable } from 'svelte/store';
import { addStock, removeStock } from '../api/stock.js';

export const stockLoading = writable(false);
export const stockError = writable(null);

// Add stock
export const stockIn = async (payload) => {
  stockLoading.set(true);
  stockError.set(null);
  try {
    const res = await addStock(payload);
    return res;
  } catch (err) {
    stockError.set(err);
    throw err;
  } finally {
    stockLoading.set(false);
  }
};

// Remove stock
export const stockOut = async (payload) => {
  stockLoading.set(true);
  stockError.set(null);
  try {
    const res = await removeStock(payload);
    return res;
  } catch (err) {
    stockError.set(err);
    throw err;
  } finally {
    stockLoading.set(false);
  }
};

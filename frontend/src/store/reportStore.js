// src/store/reportStore.js
import { writable } from 'svelte/store';
import { getStockReport } from '../api/report.js';

export const reportLoading = writable(false);
export const reportError = writable(null);
export const stockReport = writable([]);

// Fetch stock report
export const fetchStockReport = async (from, to) => {
  reportLoading.set(true);
  reportError.set(null);

  try {
    const res = await getStockReport(from, to);
    stockReport.set(res.transactions || []);
  } catch (err) {
    reportError.set(err);
  } finally {
    reportLoading.set(false);
  }
};

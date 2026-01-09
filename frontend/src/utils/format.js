// src/utils/format.js

export const formatDate = (dateStr) => {
  if (!dateStr) return '';
  const date = new Date(dateStr);
  return date.toLocaleDateString();
};

export const formatDecimal = (num) => {
  return Number(num).toFixed(2);
};

export const formatCurrency = (num, symbol = '₹') => {
  return symbol + Number(num).toFixed(2);
};

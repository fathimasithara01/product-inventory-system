// src/utils/validator.js

export const isRequired = (value) => {
  return value !== null && value !== undefined && value.toString().trim() !== '';
};

export const isPositiveNumber = (value) => {
  return !isNaN(value) && Number(value) > 0;
};

export const isValidDate = (value) => {
  return !isNaN(new Date(value).getTime());
};

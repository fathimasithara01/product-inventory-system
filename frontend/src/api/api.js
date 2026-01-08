const BASE_URL = "http://localhost:8080/api";

/**
 * Helper function to handle API responses
 */
async function handleResponse(response) {
  const data = await response.json();
  if (!response.ok) {
    throw new Error(data.error || data.message || "Something went wrong");
  }
  return data;
}

/* =========================
   PRODUCT APIs
========================= */

/**
 * Create Product with variants & sub-variants
 */
export async function createProduct(payload) {
  const response = await fetch(`${BASE_URL}/products`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(payload),
  });

  return handleResponse(response);
}

/**
 * Get Products with pagination
 */
export async function getProducts(page = 1, limit = 10) {
  const response = await fetch(
    `${BASE_URL}/products?page=${page}&limit=${limit}`
  );

  return handleResponse(response);
}

/**
 * Get SubVariants (for stock operations)
 */
export async function getSubVariants() {
  const response = await fetch(`${BASE_URL}/subvariants`);
  return handleResponse(response);
}

/* =========================
   STOCK APIs
========================= */

/**
 * Add stock (Stock In)
 */
export async function addStock(payload) {
  const response = await fetch(`${BASE_URL}/stock/in`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(payload),
  });

  return handleResponse(response);
}

/**
 * Remove stock (Stock Out)
 */
export async function removeStock(payload) {
  const response = await fetch(`${BASE_URL}/stock/out`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(payload),
  });

  return handleResponse(response);
}

/* =========================
   REPORT APIs
========================= */

/**
 * Get Stock Report by date range
 */
export async function getStockReport(fromDate, toDate) {
  const response = await fetch(
    `${BASE_URL}/reports/stock?from=${fromDate}&to=${toDate}`
  );

  return handleResponse(response);
}

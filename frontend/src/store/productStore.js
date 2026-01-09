// src/store/productStore.js
import { writable } from 'svelte/store';
import { listProducts, createProduct } from '../api/product.js';

export const products = writable([]);
export const totalProducts = writable(0);
export const productLoading = writable(false);
export const productError = writable(null);

// Fetch products with pagination
export const fetchProducts = async (page = 1, limit = 10) => {
  productLoading.set(true);
  productError.set(null);

  try {
    const res = await listProducts(page, limit);
    products.set(res.products || []);
    totalProducts.set(res.total || res.products.length);
  } catch (err) {
    productError.set(err);
  } finally {
    productLoading.set(false);
  }
};

// Transform Svelte product to backend DTO
const transformProductForAPI = (product) => ({
  product_code: product.product_code,
  product_name: product.product_name,
  product_image: product.product_image || "",
  created_user: product.created_user,
  hsn_code: product.hsn_code || "",
  variants: product.variants.map(v => ({ name: v.name })),
  options: product.variants.flatMap((v, idx) =>
    v.options.map(opt => ({ variant_index: idx, value: opt }))
  ),
  sub_variants: product.sub_variants.map(s => ({
    option_indices: s.optionIndices,
    sku: s.sku,
    stock: s.stock || "0"
  }))
});

// Add a new product
export const addProduct = async (product) => {
  productLoading.set(true);
  productError.set(null);

  try {
    const payload = transformProductForAPI(product);
    const res = await createProduct(payload);

    // Optionally update products store
    products.update((prev) => [res.product, ...prev]);

    return res;
  } catch (err) {
    productError.set(err);
    throw err;
  } finally {
    productLoading.set(false);
  }
};

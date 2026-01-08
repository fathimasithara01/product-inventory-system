<script>
  import { onMount } from "svelte";
  import { getProducts } from "../api/api.js";
  import Pagination from "./Pagination.svelte";

  let products = [], loading = false, error = "";
  let page = 1, limit = 10, total = 0;

  async function loadProducts() {
    try {
      loading = true; error = "";
      const res = await getProducts(page, limit);
      products = res.data;
      total = res.total;
    } catch(e) {
      error = "Failed to load products";
    } finally {
      loading = false;
    }
  }

  onMount(loadProducts);
</script>

<h2>Product List</h2>
{#if loading}<p>Loading...</p>
{:else if error}<p class="error">{error}</p>
{:else}
  <table>
    <thead><tr><th>Name</th><th>Code</th><th>HSN</th><th>Total Stock</th><th>Status</th></tr></thead>
    <tbody>{#each products as p}
      <tr>
        <td>{p.product_name}</td>
        <td>{p.product_code}</td>
        <td>{p.hsn_code}</td>
        <td>{p.total_stock}</td>
        <td>{p.active ? "Active" : "Inactive"}</td>
      </tr>
    {/each}</tbody>
  </table>
  <Pagination {page} {limit} {total} on:change={(e)=>{ page=e.detail.page; loadProducts(); }} />
{/if}

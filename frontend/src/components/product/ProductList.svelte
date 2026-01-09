<!-- src/components/product/ProductList.svelte -->
<script>
  import { products, fetchProducts, totalProducts } from '../../store/productStore.js';
  import Pagination from '../common/Pagination.svelte';

  let page = 1;
  let limit = 5;

  // Reactive statement: fetch products whenever page or limit changes
  $: fetchProducts(page, limit);

  const handlePageChange = (p) => {
    page = p;
    fetchProducts(page, limit);
  };
</script>

<div class="p-4 border rounded shadow">
  <h2 class="text-lg font-bold mb-2">Product List</h2>

  {#if $products.length === 0}
    <p>No products found.</p>
  {:else}
    <table class="border w-full">
      <thead class="bg-gray-200">
        <tr>
          <th class="border p-1">Name</th>
          <th class="border p-1">Code</th>
          <th class="border p-1">Image</th>
          <th class="border p-1">Variants</th>
        </tr>
      </thead>
      <tbody>
        {#each $products as product}
          <tr>
            <td class="border p-1">{product.product_name}</td>
            <td class="border p-1">{product.product_code}</td>
            <td class="border p-1">
              <img src={product.product_image} alt={product.product_name} class="w-12 h-12 object-cover" />
            </td>
            <td class="border p-1">
              {#each product.variants || [] as v}
                <div>{v.name}: {v.options.join(', ')}</div>
              {/each}
            </td>
          </tr>
        {/each}
      </tbody>
    </table>

    <!-- Pagination -->
    <Pagination
      currentPage={page}
      totalPages={Math.ceil($totalProducts / limit)}
      onPageChange={handlePageChange}
    />
  {/if}
</div>

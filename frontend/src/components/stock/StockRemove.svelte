<!-- src/components/stock/StockRemove.svelte -->
<script>
  import { stockOut } from '../../store/stockStore.js';
  import { products } from '../../store/productStore.js';
  let selectedProduct = '';
  let selectedSubVariant = '';
  let quantity = '';

  const handleStockRemove = async () => {
    if (!selectedProduct || !selectedSubVariant || !quantity) {
      alert('Please fill all fields');
      return;
    }

    try {
      await stockOut({
        product_id: selectedProduct,
        sub_variant_id: selectedSubVariant,
        quantity
      });
      alert('Stock removed successfully!');
      quantity = '';
    } catch (err) {
      alert('Error: ' + (err.message || JSON.stringify(err)));
    }
  };
</script>

<div class="p-4 border rounded shadow mb-4">
  <h2 class="text-lg font-bold mb-2">Remove Stock</h2>

  <div class="mb-2">
    <label>Product</label>
    <select bind:value={selectedProduct} class="border p-1 rounded w-full">
      <option value="">Select Product</option>
      {#each $products as product}
        <option value={product.ID}>{product.product_name}</option>
      {/each}
    </select>
  </div>

  <div class="mb-2">
    <label>Sub Variant</label>
    <select bind:value={selectedSubVariant} class="border p-1 rounded w-full">
      <option value="">Select Sub Variant</option>
      {#each $products.find(p => p.ID === selectedProduct)?.sub_variants || [] as sub}
        <option value={sub.ID}>{sub.sku} (Stock: {sub.stock})</option>
      {/each}
    </select>
  </div>

  <div class="mb-2">
    <label>Quantity</label>
    <input type="number" min="1" bind:value={quantity} class="border p-1 rounded w-full" />
  </div>

  <button class="bg-red-600 text-white px-4 py-2 rounded" on:click={handleStockRemove}>
    Remove Stock
  </button>
</div>

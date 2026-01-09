<!-- src/components/product/SubVariantInput.svelte -->
<script>
  export let subVariant = { optionIndices: [], sku: '', stock: '' };
  export let variantOptions = []; // array of arrays: [["S", "M", "L"], ["Red", "Blue"]]
  export let onUpdate = (sv) => {};

  // When a dropdown changes, update the selected option index
  const handleOptionChange = (variantIndex, valueIndex) => {
    subVariant.optionIndices[variantIndex] = Number(valueIndex);
    subVariant = { ...subVariant }; // trigger reactivity
    onUpdate(subVariant);
  };

  const handleInputChange = () => {
    subVariant = { ...subVariant }; // trigger reactivity
    onUpdate(subVariant);
  };
</script>

<div class="border p-2 rounded mb-2">
  {#each variantOptions as options, i}
    <select
      class="border p-1 rounded mr-2 mb-1"
      on:change={(e) => handleOptionChange(i, e.target.value)}
    >
      <option value="">Select Option</option>
      {#each options as opt, idx}
        <option value={idx} selected={subVariant.optionIndices[i] == idx}>
          {opt}
        </option>
      {/each}
    </select>
  {/each}

  <input
    type="text"
    placeholder="SKU"
    bind:value={subVariant.sku}
    class="border p-1 rounded w-full mb-1"
    on:input={handleInputChange}
  />

  <input
    type="number"
    placeholder="Stock"
    bind:value={subVariant.stock}
    class="border p-1 rounded w-full"
    on:input={handleInputChange}
  />
</div>

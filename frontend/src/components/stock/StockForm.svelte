<script>
  import { addStock, removeStock } from "../api/api.js";

  export let subVariantOptions = []; // Array of available sub-variants (id + SKU)
  export let type = "IN"; // "IN" for add stock, "OUT" for remove stock

  let selectedSubVariant = "";
  let quantity = "";
  let loading = false;
  let error = "";
  let success = "";

  async function handleSubmit() {
    error = success = "";
    if (!selectedSubVariant || !quantity || Number(quantity) <= 0) {
      error = "Select a variant and enter a valid quantity";
      return;
    }

    try {
      loading = true;
      if (type === "IN") await addStock({ sub_variant_id: selectedSubVariant, quantity });
      else await removeStock({ sub_variant_id: selectedSubVariant, quantity });
      success = `Stock ${type === "IN" ? "added" : "removed"} successfully`;
      selectedSubVariant = "";
      quantity = "";
    } catch (err) {
      error = `Failed to ${type === "IN" ? "add" : "remove"} stock`;
    } finally {
      loading = false;
    }
  }
</script>

<h3>{type === "IN" ? "Add Stock" : "Remove Stock"}</h3>
<form on:submit|preventDefault={handleSubmit}>
  <div class="field">
    <label>Sub-Variant</label>
    <select bind:value={selectedSubVariant}>
      <option value="">Select</option>
      {#each subVariantOptions as sv}
        <option value={sv.id}>{sv.sku}</option>
      {/each}
    </select>
  </div>

  <div class="field">
    <label>Quantity</label>
    <input type="number" bind:value={quantity} min="0" step="0.01" />
  </div>

  <button disabled={loading}>{loading ? "Processing..." : (type==="IN"?"Add Stock":"Remove Stock")}</button>

  {#if error}<p class="error">{error}</p>{/if}
  {#if success}<p class="success">{success}</p>{/if}
</form>

<style>
  .field { margin-bottom: 12px; display: flex; flex-direction: column; }
  select, input { padding: 6px; font-size: 14px; }
  button { padding: 8px; margin-top: 10px; cursor: pointer; }
  .error { color: red; margin-top: 8px; }
  .success { color: green; margin-top: 8px; }
</style>

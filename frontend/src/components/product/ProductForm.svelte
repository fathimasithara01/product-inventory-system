<script>
  import { createProduct } from "../api/api.js";

  let productName = "";
  let productCode = "";
  let hsnCode = "";
  let active = true;

  let loading = false;
  let error = "";
  let success = "";

  async function handleSubmit() {
    error = success = "";
    if (!productName || !productCode) { error = "Name & code required"; return; }

    try {
      loading = true;
      await createProduct({ product_name: productName, product_code: productCode, hsn_code: hsnCode, active });
      success = "Product created successfully";
      productName = productCode = hsnCode = ""; active = true;
    } catch (err) { error = "Failed to create product"; }
    finally { loading = false; }
  }
</script>

<h2>Create Product</h2>
<form on:submit|preventDefault={handleSubmit}>
  <div class="field"><label>Name</label><input bind:value={productName} /></div>
  <div class="field"><label>Code</label><input bind:value={productCode} /></div>
  <div class="field"><label>HSN</label><input bind:value={hsnCode} /></div>
  <div class="field checkbox"><input type="checkbox" bind:checked={active} /> Active</div>
  <button disabled={loading}>{loading ? "Saving..." : "Create"}</button>
  {#if error}<p class="error">{error}</p>{/if}
  {#if success}<p class="success">{success}</p>{/if}
</form>

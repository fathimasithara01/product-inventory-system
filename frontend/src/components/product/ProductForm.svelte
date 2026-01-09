<!-- src/components/product/ProductForm.svelte -->
<script>
  import { addProduct } from '../../store/productStore.js';
  import VariantInput from './VariantInput.svelte';
  import SubVariantInput from './SubVariantInput.svelte';

  let product = {
    product_id: 0,
    product_code: '',
    product_name: '',
    product_image: '',
    hsn_code: '',
    created_user: 'c8bcdcb4-2113-4f1f-b861-5f57fe0dcad7', // example user ID
    variants: [],
    sub_variants: []
  };

  // Add a new variant
  const addVariant = () => {
    product.variants = [...product.variants, { name: '', options: [] }];
  };

  // Add a new sub-variant
  const addSubVariant = () => {
    product.sub_variants = [
      ...product.sub_variants,
      { optionIndices: [], sku: '', stock: '' }
    ];
  };

  // Update a variant
  const handleVariantUpdate = (variant, index) => {
    product.variants[index] = variant;
    product = { ...product }; // trigger reactivity
  };

  // Update a sub-variant
  const handleSubVariantUpdate = (subVariant, index) => {
    product.sub_variants[index] = subVariant;
    product = { ...product }; // trigger reactivity
  };

  // Submit product to backend
  const handleSubmit = async () => {
    try {
      await addProduct(product);
      alert('Product added successfully!');

      // Reset form
      product = {
        product_id: 0,
        product_code: '',
        product_name: '',
        product_image: '',
        hsn_code: '',
        created_user: product.created_user,
        variants: [],
        sub_variants: []
      };
    } catch (err) {
      alert(err?.message || 'Failed to add product');
    }
  };
</script>

<div class="bg-white p-6 rounded shadow-md mb-6">
  <h2 class="text-xl font-bold mb-4">Add Product</h2>

  <!-- Product Basic Info -->
  <div class="space-y-2 mb-4">
    <input type="text" placeholder="Product Name" class="border p-2 w-full rounded" bind:value={product.product_name} />
    <input type="text" placeholder="Product Code" class="border p-2 w-full rounded" bind:value={product.product_code} />
    <input type="text" placeholder="Product Image URL" class="border p-2 w-full rounded" bind:value={product.product_image} />
    <input type="text" placeholder="HSN Code" class="border p-2 w-full rounded" bind:value={product.hsn_code} />
  </div>

  <!-- Variants Section -->
  <div class="mb-4">
    <h3 class="font-semibold mb-2">Variants</h3>
    {#each product.variants as variant, i}
      <VariantInput {variant} onUpdate={(v) => handleVariantUpdate(v, i)} />
    {/each}
    <button type="button" class="bg-green-600 text-white px-3 py-1 rounded mt-2" on:click={addVariant}>
      Add Variant
    </button>
  </div>

  <!-- Sub-Variants Section -->
  <div class="mb-4">
    <h3 class="font-semibold mb-2">Sub Variants</h3>
    {#each product.sub_variants as sub, i}
      <SubVariantInput 
        subVariant={sub} 
        variantOptions={product.variants.map(v => v.name)} 
        onUpdate={(s) => handleSubVariantUpdate(s, i)}
      />
    {/each}
    <button type="button" class="bg-blue-600 text-white px-3 py-1 rounded mt-2" on:click={addSubVariant}>
      Add Sub Variant
    </button>
  </div>

  <!-- Submit Button -->
  <button class="bg-black text-white px-4 py-2 rounded mt-4" on:click={handleSubmit}>
    Save Product
  </button>
</div>

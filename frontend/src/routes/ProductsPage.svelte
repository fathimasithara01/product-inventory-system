<!-- src/routes/ProductsPage.svelte -->
<script>
  import Header from '../components/common/Header.svelte';
  import ProductForm from '../components/product/ProductForm.svelte';
  import ProductList from '../components/product/ProductList.svelte';
</script>

<Header title="Products" />

<div class="p-4 space-y-6">
  <ProductForm />
  <ProductList />
</div>

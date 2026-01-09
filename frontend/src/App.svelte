<script>
  import Router from 'svelte-spa-router';

  import ProductsPage from './routes/ProductsPage.svelte';
  import StockPage from './routes/StockPage.svelte';
  import StockReportPage from './routes/StockReportPage.svelte';

  const routes = {
    '/': ProductsPage,
    '/stock': StockPage,
    '/report': StockReportPage
  };
</script>

<!-- Navigation Bar -->
<nav class="bg-blue-600 text-white p-4 flex gap-4">
  <a href="#/" class="hover:underline">Products</a>
  <a href="#/stock" class="hover:underline">Stock</a>
  <a href="#/report" class="hover:underline">Stock Report</a>
</nav>

<!-- Router -->
<main class="p-4">
  <Router {routes} />
</main>

<style>
  nav a {
    font-weight: bold;
  }
</style>

<!-- src/routes/StockPage.svelte -->
<script>
  import Header from '../components/common/Header.svelte';
  import StockAdd from '../components/stock/StockAdd.svelte';
  import StockRemove from '../components/stock/StockRemove.svelte';
</script>

<Header title="Stock Management" />

<div class="p-4 grid gap-4 md:grid-cols-2">
  <StockAdd />
  <StockRemove />
</div>

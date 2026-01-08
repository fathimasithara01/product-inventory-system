<script>
  import StockIn from "./StockIn.svelte";
  import StockOut from "./StockOut.svelte";
</script>

<h2>Stock Management</h2>
<div class="stock-management">
  <div class="column"><StockIn /></div>
  <div class="column"><StockOut /></div>
</div>

<style>
  .stock-management { display: flex; gap: 40px; flex-wrap: wrap; }
  .column { flex: 1 1 300px; }
</style>

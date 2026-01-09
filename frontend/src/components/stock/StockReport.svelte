<!-- src/components/stock/StockReport.svelte -->
<script>
  import { fetchStockReport, stockReport } from '../../store/reportStore.js';
  import { onMount } from 'svelte';

  let fromDate = '';
  let toDate = '';

  const handleFetchReport = async () => {
    if (!fromDate || !toDate) {
      alert('Please select both From and To dates');
      return;
    }
    await fetchStockReport(fromDate, toDate);
  };

  // Optional: fetch default last 7 days on mount
  onMount(() => {
    const today = new Date();
    const lastWeek = new Date();
    lastWeek.setDate(today.getDate() - 7);
    fromDate = lastWeek.toISOString().split('T')[0];
    toDate = today.toISOString().split('T')[0];
    fetchStockReport(fromDate, toDate);
  });

  const calculateNet = (txn) => {
    return Number(txn.stock_in || 0) - Number(txn.stock_out || 0);
  };
</script>

<div class="p-4 border rounded shadow">
  <h2 class="text-lg font-bold mb-2">Stock Report</h2>

  <div class="flex gap-2 mb-4">
    <div>
      <label>From</label>
      <input type="date" bind:value={fromDate} class="border p-1 rounded" />
    </div>
    <div>
      <label>To</label>
      <input type="date" bind:value={toDate} class="border p-1 rounded" />
    </div>
    <button class="bg-blue-600 text-white px-3 py-1 rounded mt-5" on:click={handleFetchReport}>
      Fetch Report
    </button>
  </div>

  {#if $stockReport.length === 0}
    <p>No transactions found.</p>
  {:else}
    <table class="border w-full">
      <thead class="bg-gray-200">
        <tr>
          <th class="border p-1">Date</th>
          <th class="border p-1">Product</th>
          <th class="border p-1">Sub Variant (SKU)</th>
          <th class="border p-1">Stock IN</th>
          <th class="border p-1">Stock OUT</th>
          <th class="border p-1">Net Stock</th>
        </tr>
      </thead>
      <tbody>
        {#each $stockReport as txn}
          <tr>
            <td class="border p-1">{new Date(txn.transaction_date).toLocaleDateString()}</td>
            <td class="border p-1">{txn.product_name}</td>
            <td class="border p-1">{txn.sku}</td>
            <td class="border p-1">{txn.transaction_type === 'IN' ? txn.quantity : 0}</td>
            <td class="border p-1">{txn.transaction_type === 'OUT' ? txn.quantity : 0}</td>
            <td class="border p-1">{calculateNet(txn)}</td>
          </tr>
        {/each}
      </tbody>
    </table>
  {/if}
</div>

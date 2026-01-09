<!-- src/components/stock/StockReport.svelte -->
<script>
  import { fetchStockReport, stockReport } from '../../store/reportStore.js';
  import { onMount } from 'svelte';

  let fromDate = '';
  let toDate = '';

  // Fetch report for selected dates
  const handleFetchReport = async () => {
    if (!fromDate || !toDate) {
      alert('Please select both From and To dates');
      return;
    }
    await fetchStockReport(fromDate, toDate);
  };

  // Default: last 7 days
  onMount(() => {
    const today = new Date();
    const lastWeek = new Date();
    lastWeek.setDate(today.getDate() - 7);

    fromDate = lastWeek.toISOString().split('T')[0];
    toDate = today.toISOString().split('T')[0];

    fetchStockReport(fromDate, toDate);
  });

  // Aggregate stock per product-subvariant
  const getAggregatedStock = () => {
    const agg = {};
    $stockReport.forEach(txn => {
      const key = txn.product_id + '-' + txn.sub_variant_id;
      if (!agg[key]) {
        agg[key] = {
          product_name: txn.product_name,
          sku: txn.sku,
          stock_in: 0,
          stock_out: 0
        };
      }
      if (txn.transaction_type === 'IN') agg[key].stock_in += Number(txn.quantity);
      if (txn.transaction_type === 'OUT') agg[key].stock_out += Number(txn.quantity);
    });
    return Object.values(agg).map(item => ({
      ...item,
      net_stock: item.stock_in - item.stock_out
    }));
  };
</script>

<div class="p-4 border rounded shadow">
  <h2 class="text-lg font-bold mb-2">Stock Report</h2>

  <!-- Date Filter -->
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

  <!-- Stock Table -->
  {#if $stockReport.length === 0}
    <p>No transactions found.</p>
  {:else}
    <table class="border w-full">
      <thead class="bg-gray-200">
        <tr>
          <th class="border p-1">Product</th>
          <th class="border p-1">Sub Variant (SKU)</th>
          <th class="border p-1">Stock IN</th>
          <th class="border p-1">Stock OUT</th>
          <th class="border p-1">Net Stock</th>
        </tr>
      </thead>
      <tbody>
        {#each getAggregatedStock() as txn}
          <tr>
            <td class="border p-1">{txn.product_name}</td>
            <td class="border p-1">{txn.sku}</td>
            <td class="border p-1">{txn.stock_in}</td>
            <td class="border p-1">{txn.stock_out}</td>
            <td class="border p-1">{txn.net_stock}</td>
          </tr>
        {/each}
      </tbody>
    </table>
  {/if}
</div>

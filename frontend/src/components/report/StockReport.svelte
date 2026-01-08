<script>
  import { getStockReport } from "../api/api.js";

  let fromDate="", toDate="", report=[], loading=false, error="";
  let totalIn=0, totalOut=0;

  async function loadReport() {
    loading=true; error=""; totalIn=totalOut=0;
    try {
      const res = await getStockReport(fromDate, toDate);
      report = res.data || [];
      report.forEach(r => r.transaction_type==="IN"? totalIn+=Number(r.quantity): totalOut+=Number(r.quantity));
    } catch(e){ error="Failed to load report"; }
    finally{ loading=false; }
  }
</script>

<h2>Stock Report</h2>
<div class="filters">
  <input type="date" bind:value={fromDate} />
  <input type="date" bind:value={toDate} />
  <button on:click={loadReport}>Apply</button>
</div>

{#if loading}<p>Loading...</p>
{:else if error}<p class="error">{error}</p>
{:else if report.length===0}<p>No data</p>
{:else}
  <table>
    <thead><tr><th>Date</th><th>Product</th><th>Type</th><th>Qty</th></tr></thead>
    <tbody>{#each report as r}
      <tr>
        <td>{r.transaction_date}</td><td>{r.product_name || r.product_id}</td>
        <td>{r.transaction_type}</td><td>{r.quantity}</td>
      </tr>
    {/each}</tbody>
  </table>
  <div class="summary">
    <p><strong>Total IN:</strong> {totalIn}</p>
    <p><strong>Total OUT:</strong> {totalOut}</p>
    <p><strong>Net:</strong> {totalIn-totalOut}</p>
  </div>
{/if}

<style>
  .filters { margin-bottom: 12px; display: flex; gap: 8px; }
  table { width: 100%; border-collapse: collapse; margin-top: 12px; }
  th, td { padding: 8px; border: 1px solid #ddd; }
  .summary { margin-top: 12px; font-size: 14px; }
  .error { color: red; }
</style>

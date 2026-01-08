<script>
  import { createEventDispatcher } from "svelte";

  export let page = 1;
  export let limit = 10;
  export let total = 0;

  const dispatch = createEventDispatcher();
  $: totalPages = Math.ceil(total / limit);

  function prev() { if (page > 1) dispatch('change', { page: page - 1 }); }
  function next() { if (page < totalPages) dispatch('change', { page: page + 1 }); }
</script>

{#if totalPages > 1}
  <div class="pagination">
    <button on:click={prev} disabled={page === 1}>Prev</button>
    <span>Page {page} of {totalPages}</span>
    <button on:click={next} disabled={page === totalPages}>Next</button>
  </div>
{/if}

<style>
  .pagination { display: flex; justify-content: center; gap: 12px; margin-top: 20px; }
  button { padding: 6px 12px; cursor: pointer; }
  button:disabled { opacity: 0.5; cursor: not-allowed; }
</style>

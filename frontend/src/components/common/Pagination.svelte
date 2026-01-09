<!-- src/components/common/Pagination.svelte -->
<script>
  export let currentPage = 1;
  export let totalPages = 1;
  export let onPageChange = (page) => {};

  // Reactive statement ensures pages update if totalPages changes
  $: pages = Array.from({ length: totalPages }, (_, i) => i + 1);
</script>

<div class="flex space-x-2 mt-4">
  {#each pages as page}
    <button
      class="px-3 py-1 border rounded {page === currentPage ? 'bg-blue-600 text-white' : 'bg-white'}"
      on:click={() => onPageChange(page)}
    >
      {page}
    </button>
  {/each}
</div>
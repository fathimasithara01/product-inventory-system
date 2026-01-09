<!-- src/components/common/Modal.svelte -->
<script>
  export let show = false;
  export let onClose = () => {};
  export let title = ""; // optional title for the modal
</script>

{#if show}
<div class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50" role="dialog" aria-modal="true">
  <div class="bg-white p-6 rounded shadow-lg w-full max-w-md">
    {#if title}
      <h2 class="text-lg font-bold mb-4">{title}</h2>
    {/if}
    <slot></slot>
    <div class="mt-4 text-right">
      <button 
        class="bg-red-600 text-white px-4 py-2 rounded hover:bg-red-700"
        on:click={onClose}
      >
        Close
      </button>
    </div>
  </div>
</div>
{/if}

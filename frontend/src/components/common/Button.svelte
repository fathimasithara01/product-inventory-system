<!-- src/components/common/Button.svelte -->
<script>
  export let type = "button";
  export let disabled = false;
  export let onClick = () => {};
  export let label = "Button";
  export let className = ""; // allow custom classes
</script>

<button
  type={type}
  class={`bg-blue-600 hover:bg-blue-700 text-white px-4 py-2 rounded disabled:opacity-50 disabled:cursor-not-allowed ${className}`}
  disabled={disabled}
  on:click={onClick}
>
  {label}
</button>

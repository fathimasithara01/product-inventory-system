<!-- src/components/common/Header.svelte -->
<script>
  export let title = "Product Inventory";
  export let className = ""; // allow extra Tailwind classes if needed
</script>

<header class={`bg-blue-600 text-white p-4 shadow-md ${className}`}>
  <h1 class="text-xl font-bold">{title}</h1>
</header>

<!-- src/components/product/VariantInput.svelte -->
<script>
  export let variant = { name: '', options: [] };
  export let onUpdate = (v) => {};

  let optionValue = '';

  const addOption = () => {
    const val = optionValue.trim();
    if (val) {
      variant.options = [...variant.options, val]; // trigger reactivity
      optionValue = '';
      onUpdate(variant);
    }
  };

  const removeOption = (index) => {
    variant.options = variant.options.filter((_, i) => i !== index); // reactive removal
    onUpdate(variant);
  };
</script>

<div class="border p-2 rounded mb-2">
  <!-- Variant Name -->
  <input
    type="text"
    placeholder="Variant Name (e.g., Size, Color)"
    bind:value={variant.name}
    class="border p-1 rounded w-full mb-2"
    on:input={() => onUpdate({ ...variant })}
  />

  <!-- Option Input -->
  <div class="flex gap-2 mb-2">
    <input
      type="text"
      placeholder="Option Value"
      bind:value={optionValue}
      class="border p-1 rounded flex-1"
      on:keydown={(e) => e.key === 'Enter' && addOption()}
    />
    <button
      type="button"
      class="bg-green-600 text-white px-3 py-1 rounded"
      on:click={addOption}
    >
      Add Option
    </button>
  </div>

  <!-- Option List -->
  <div class="flex flex-wrap gap-2">
    {#each variant.options as opt, i}
      <span class="bg-gray-200 px-2 py-1 rounded flex items-center gap-1">
        {opt}
        <button
          type="button"
          class="text-red-500 font-bold"
          on:click={() => removeOption(i)}
        >
          ×
        </button>
      </span>
    {/each}
  </div>
</div>

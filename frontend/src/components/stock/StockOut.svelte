<script>
  import StockForm from "./StockForm.svelte";
  import { getSubVariants } from "../api/api.js";

  let subVariants = [];

  async function loadSubVariants() {
    try { subVariants = await getSubVariants(); }
    catch(e){ console.error("Failed to load sub-variants"); }
  }

  loadSubVariants();
</script>

<StockForm {subVariants} type="OUT" />

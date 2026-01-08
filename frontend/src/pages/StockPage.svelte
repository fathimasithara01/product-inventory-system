<script>
  import Header from "../components/common/Header.svelte";
  import Footer from "../components/common/Footer.svelte";
  import StockManagement from "../components/stock/StockManagement.svelte";
</script>

<Header />

<main>
  <StockManagement />
</main>

<Footer />

<style>
  main { padding: 20px; max-width: 1200px; margin: auto; }
</style>

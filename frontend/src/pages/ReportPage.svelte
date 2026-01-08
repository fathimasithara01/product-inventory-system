<script>
  import Header from "../components/common/Header.svelte";
  import Footer from "../components/common/Footer.svelte";
  import StockReport from "../components/report/StockReport.svelte";
</script>

<Header />

<main>
  <StockReport />
</main>

<Footer />

<style>
  main { padding: 20px; max-width: 1200px; margin: auto; }
</style>

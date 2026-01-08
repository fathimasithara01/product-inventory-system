<script>
  import Header from "../components/common/Header.svelte";
  import Footer from "../components/common/Footer.svelte";
  import ProductForm from "../components/product/ProductForm.svelte";
  import ProductList from "../components/product/ProductList.svelte";
</script>

<Header />

<main>
  <section>
    <ProductForm />
  </section>

  <section style="margin-top: 40px;">
    <ProductList />
  </section>
</main>

<Footer />

<style>
  main { padding: 20px; max-width: 1200px; margin: auto; }
  section { margin-bottom: 40px; }
</style>

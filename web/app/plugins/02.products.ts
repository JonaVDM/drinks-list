export default defineNuxtPlugin(async () => {
  const productStore = useProductStore()
  await productStore.fetchProducts()
});

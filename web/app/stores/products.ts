export const useProductStore = defineStore('products', () => {
  const products = ref<Product[]>([])
  const loading = ref(false)
  const error = ref()

  const byId = computed(() => {
    return products.value.reduce<Record<string, Product>>((prev, curr) => {
      prev[curr.id] = curr
      return prev
    }, {})
  })

  const fetchProducts = async () => {
    loading.value = true
    error.value = null;

    const pb = usePocketbase()

    try {
      products.value = await pb.collection('products').getFullList();
    } catch (err) {
      error.value = err;
    } finally {
      loading.value = false;
    }
  }

  return {
    products,
    loading,
    error,

    byId,

    fetchProducts,
  }
})

export const useBasketStore = defineStore('basket', () => {
  const productsStore = useProductStore()
  const products = ref<Record<string, number>>({})

  const total = computed(() => {
    return Object.entries(products.value).reduce<number>((prev, [id, amount]) => {
      return prev + amount * (productsStore.byId[id]?.price ?? 0)
    }, 0);
  })

  const add = (id: string) => {
    if (!products.value[id]) {
      products.value[id] = 1
    } else {
      products.value[id] += 1
    }
  }

  const remove = (id: string) => {
    if (!products.value[id]) {
      return
    } else if (products.value[id] == 1) {
      delete products.value[id]
    } else {
      products.value[id] -= 1
    }
  }

  const clear = () => {
    products.value = {};
  }

  return {
    products,

    total,

    add,
    remove,
    clear,
  }
}, {
  persist: true,
})

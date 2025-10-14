export const useBasketStore = defineStore('basket', () => {
  const products = ref<Record<string, number>>({})

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

    add,
    remove,
    clear,
  }
}, {
  persist: true,
})

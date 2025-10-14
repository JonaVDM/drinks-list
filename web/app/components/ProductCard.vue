<script setup lang="ts">
const props = defineProps<{
  product: Product
}>()

const basketStore = useBasketStore()
const pb = usePocketbase()
const img = computed(() => pb.files.getURL(props.product, props.product.photo, { thumb: '500x300' }))
const price = computed(() => {
  return new Intl.NumberFormat('nl-NL', {
    style: 'currency',
    currency: 'EUR',
  }).format(props.product.price);
})
</script>

<template>
  <div class="card bg-base-300 shadow-lg">
    <figure>
      <img :src="img" alt="Shoes" class="w-full" />
    </figure>
    <div class="card-body">
      <h3 class="card-title">{{ product.name }}</h3>

      <p>{{ price }}</p>
      <div class="card-actions justify-center">
        <div class="flex gap-2 items-center">
          <button class="btn btn-primary" @click="basketStore.remove(product.id)">-</button>
          <p>
            {{ basketStore.products[product.id] }}
          </p>
          <button class="btn btn-primary" @click="basketStore.add(product.id)">+</button>
        </div>
      </div>
    </div>
  </div>
</template>

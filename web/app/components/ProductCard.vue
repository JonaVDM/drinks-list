<script setup lang="ts">
const props = defineProps<{
  drink: Drink
}>()

const pb = usePocketbase()
const img = computed(() => pb.files.getURL(props.drink, props.drink.photo, { thumb: '500x300' }))
const price = computed(() => {
  return new Intl.NumberFormat('nl-NL', {
    style: 'currency',
    currency: 'EUR',
  }).format(props.drink.price);
})
</script>

<template>
  <div class="card bg-base-300 shadow-lg">
    <figure>
      <img :src="img" alt="Shoes" class="w-full" />
    </figure>
    <div class="card-body">
      <h3 class="card-title">{{ props.drink.drink }}</h3>

      <p>{{ price }}</p>
      <div class="card-actions justify-end">
        <button class="btn btn-primary">Buy Now</button>
      </div>
    </div>
  </div>
</template>

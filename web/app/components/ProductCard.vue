<script setup lang="ts">
const props = defineProps<{
  drink: Drink
}>()

const basketStore = useBasketStore()
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
      <div class="card-actions justify-center">
        <div class="flex gap-2 items-center">
          <button class="btn btn-primary" @click="basketStore.remove(props.drink.id)">-</button>
          <p>
            {{ basketStore.products[drink.id] }}
          </p>
          <button class="btn btn-primary" @click="basketStore.add(props.drink.id)">+</button>
        </div>
      </div>
    </div>
  </div>
</template>

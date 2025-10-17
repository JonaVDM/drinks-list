<script lang="ts" setup>
definePageMeta({
  layout: 'kiosk',
});

const { loading, products, error } = useProductStore()
</script>

<template>
  <div class="alert alert-error" v-if="error">
    <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6 shrink-0 stroke-current" fill="none" viewBox="0 0 24 24">
      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
        d="M10 14l2-2m0 0l2-2m-2 2l-2-2m2 2l2 2m7-2a9 9 0 11-18 0 9 9 0 0118 0z" />
    </svg>

    <span>Failed to load products, {{ error.message }}</span>
  </div>


  <div v-if="loading">Loading</div>

  <div v-else>
    <NuxtLink to="/kiosk/checkout" class="md:hidden my-2 btn btn-primary w-full btn-xl">Betaal</NuxtLink>

    <div class="md:flex items-start">
      <div class="grid md:grid-cols-2 xl:grid-cols-3 gap-4">
        <ProductCard v-for="product in products" :product></ProductCard>
      </div>

      <div class="flex-1/3 md:flex-2/3 px-3 max-md:hidden">
        <CheckoutPanel />
      </div>
    </div>
  </div>
</template>

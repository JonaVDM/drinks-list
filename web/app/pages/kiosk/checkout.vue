<script setup lang="ts">
definePageMeta({
  layout: 'kiosk',
})
const router = useRouter()
</script>

<template>
  <h2 class="text-lg font-bold pt-10">Bestel lijst</h2>
  <NuxtLink to="/kiosk" class="link link-error">Terug</NuxtLink>

  <CheckoutPanel @submit="() => router.push('/kiosk')"></CheckoutPanel>
</template>

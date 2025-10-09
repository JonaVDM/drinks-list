<script setup lang="ts">
definePageMeta({
  middleware: 'auth'
})

const userStore = useUserStore()
</script>

<template>
  <pre>{{ JSON.stringify(userStore.user, null, 2) }}</pre>
</template>

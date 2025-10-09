<script setup lang="ts">
const pb = usePocketbase()

const { data } = await useAsyncData(() => {
  return pb.collection('drinks').getFullList()
})
</script>

<template>
  <div>
    <h1 class="text-2xl font-bold text-center">Hello World</h1>

    <pre>{{ JSON.stringify(data, null, 2) }}</pre>
  </div>
</template>

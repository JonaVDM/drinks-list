<script setup lang="ts">
const userStore = useUserStore()
const pb = usePocketbase()

function logout() {
  pb.authStore.clear()
  reloadNuxtApp()
}
</script>

<template>
  <div v-if="userStore.authed" class="p-2 text-right">
    <p>Loged in as {{ userStore.user?.name }}</p>
    <button class="btn btn-sm btn-error" @click="logout">Logout</button>
  </div>
  <div class="max-w-6xl mx-auto">
    <slot />
  </div>
</template>

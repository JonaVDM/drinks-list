<script setup lang="ts">
import type { RecordModel } from 'pocketbase';

const emit = defineEmits(['submit'])

const basketStore = useBasketStore()

const loading = ref(false)
const error = ref<string>()

const user = ref<RecordModel>();

async function pay() {
  if (!user.value || basketStore.itemCount() == 0) {
    return
  }
  loading.value = true
  error.value = undefined

  try {
    const pb = usePocketbase()
    pb.send("/api/create-order", {
      body: {
        user: user.value?.id,
        products: Object.entries(basketStore.products).map(([id, amount]) => {
          return {
            id,
            amount,
          }
        }),
      },
      method: 'POST',
    })
    basketStore.clear()
    user.value = undefined
    emit('submit')
  } catch (e) {
    error.value = 'Error!!!!!!!!!!'
  } finally {
    loading.value = false
  }
}
</script>

<template>
  <div>
    <div class="alert alert-error" v-if="error">
      <p>Er ging iets fout!!!</p>
    </div>

    <TotalTable :products="basketStore.products" :total="basketStore.total"></TotalTable>

    <div class="py-10">
      <UserSelect v-model="user" />
    </div>

    <button class="btn btn-success w-full" @click="pay" :disabled="!user || loading">Betaal</button>
  </div>
</template>

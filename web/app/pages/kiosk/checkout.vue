<script setup lang="ts">
import type { RecordModel } from 'pocketbase';

definePageMeta({
  layout: 'kiosk',
})

const basketStore = useBasketStore()
const productStore = useProductStore()
const router = useRouter()

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
    router.push('/kiosk')
  } catch (e) {
    error.value = 'Error!!!!!!!!!!'
  } finally {
    loading.value = false
  }
}
</script>

<template>
  <div>
    <h2 class="text-lg font-bold pt-10">Bestel lijst</h2>

    <div class="alert alert-error" v-if="error">
      <p>Er ging iets fout!!!</p>
    </div>

    <div class="overflow-x-auto rounded-box border border-base-content/5 bg-base-100">
      <table class="table table-zebra table-fixed">
        <tbody>
          <tr v-for="amount, id in basketStore.products">
            <td class="w-20">{{ amount }}</td>
            <td>{{ productStore.byId[id]?.name ?? 'idk /shrug' }}</td>
            <td class="text-right">{{ formatCurrency((productStore.byId[id]?.price ?? 0) * amount) }}</td>
          </tr>

          <tr>
            <td></td>
            <td class="font-bold">Totaal</td>
            <td class="text-right">{{ formatCurrency(basketStore.total) }}</td>
          </tr>
        </tbody>
      </table>
    </div>

    <div class="py-10 md:grid grid-cols-2 gap-5">
      <UserSelect v-model="user" />
    </div>

    <button class="btn btn-success w-full" @click="pay" :disabled="!user || loading">Betaal</button>
  </div>
</template>

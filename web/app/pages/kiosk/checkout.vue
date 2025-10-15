<script setup lang="ts">
import type { RecordModel } from 'pocketbase';

definePageMeta({
  layout: 'kiosk',
})

const basketStore = useBasketStore()
const productStore = useProductStore()

const user = ref<RecordModel>();
</script>

<template>
  <div>
    <h2 class="text-lg font-bold pt-10">Bestel lijst</h2>
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
  </div>
</template>

<script setup lang="ts">
import type { RecordModel } from 'pocketbase'

definePageMeta({
  middleware: 'auth'
})

const userStore = useUserStore()
const pb = usePocketbase()

const { data, error } = await useAsyncData(() => {
  return pb.collection('orders').getFullList({ filter: `user.id='${userStore.user?.id}'`, sort: '-created' })
})

const ordersByMonth = computed(() => data.value?.reduce<Record<string, RecordModel[]>>((prev, curr) => {
  const [year, month] = curr.created.split('-');
  const key = `${year}-${month}`;
  if (!prev[key]) {
    prev[key] = [];
  }
  prev[key].push(curr)
  return prev
}, {}))
</script>

<template>
  <div v-if="error" class="alert alert-error">Something went wrong fetching data</div>
  <div v-for="(orders, month) in ordersByMonth" class="overflow-x-auto">
    <h2 class="text-xl">{{ month }}</h2>

    <table class="table">
      <thead>
        <tr>
          <th>Date</th>
          <th>Drink</th>
          <th>Cost</th>
        </tr>
      </thead>

      <tbody>
        <tr v-for="order in orders">
          <td>{{ new Date(order.created).toLocaleString() }}</td>
          <td>{{ order.expand?.drink.drink }}</td>
          <td>{{ order.expand?.drink.price }}</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

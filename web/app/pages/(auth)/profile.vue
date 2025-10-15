<script setup lang="ts">
definePageMeta({
  middleware: 'auth'
})

const userStore = useUserStore()
const pb = usePocketbase()

const { data, error } = await useAsyncData(() => {
  return pb.collection<Order>('orders').getFullList({ filter: `user.id='${userStore.user?.id}'`, sort: '-created', expand: 'rows' })
})

const ordersByMonth = computed(() => data.value?.reduce<Record<string, Order[]>>((prev, curr) => {
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
  <div v-for="(orders, month) in ordersByMonth">
    <h2 class="text-xl">{{ month }}</h2>

    <div class="flex flex-col gap-4">
      <div v-for="order in orders">
        <h3 class="text-lg font-thin">{{ new Date(order.created).toLocaleString() }}</h3>
        <TotalTable :total="order.total" :products="order.expand.rows.reduce<Record<string, number>>((prev, curr) => {
          prev[curr.product] = curr.amount
          return prev
        }, {})"></TotalTable>
      </div>
    </div>
  </div>
</template>

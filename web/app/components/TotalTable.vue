<script setup lang="ts">
const productStore = useProductStore()

defineProps<{
  products: Record<string, number>
  total: number
}>();
</script>

<template>
  <div class="overflow-x-auto rounded-box border border-base-content/5 bg-base-200">
    <table class="table table-fixed">
      <tbody>
        <tr v-for="amount, id in products">
          <td class="w-20">{{ amount }}</td>
          <td>{{ productStore.byId[id]?.name ?? 'idk /shrug' }}</td>
          <td class="text-right">{{ formatCurrency((productStore.byId[id]?.price ?? 0) * amount) }}</td>
        </tr>

        <tr class="bg-base-300">
          <td></td>
          <td class="font-bold">Totaal</td>
          <td class="text-right">{{ formatCurrency(total) }}</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

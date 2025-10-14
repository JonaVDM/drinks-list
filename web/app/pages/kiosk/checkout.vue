<script setup lang="ts">
definePageMeta({
  layout: 'kiosk',
})

const basketStore = useBasketStore()
</script>

<template>
  <div>
    <h2 class="text-lg font-bold pt-10">Bestel lijst</h2>
    <div class="overflow-x-auto rounded-box border border-base-content/5 bg-base-100">
      <table class="table table-zebra table-fixed">
        <tbody>
          <tr v-for="amount, id in basketStore.products">
            <td class="w-20">{{ amount }}</td>
            <td>{{ id }}</td>
            <td class="text-right">10</td>
          </tr>
          <tr>
            <td></td>
            <td class="font-bold">Totaal</td>
            <td class="text-right">210</td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

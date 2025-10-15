<script setup lang="ts">
import type { RecordModel } from 'pocketbase';

const pb = usePocketbase()

defineProps<{ modelValue?: RecordModel }>()
const emit = defineEmits<{
  (e: 'update:modelValue', value: RecordModel): void
}>()

const filter = ref('');
const users = ref<RecordModel[]>([]);

async function loadUsers() {
  if (filter.value.length < 3) {
    return
  }
  const data = await pb.collection('users').getFullList({ filter: `name~'${filter.value}'`, perPage: 10 })
  users.value = data
}

onMounted(loadUsers)
</script>

<template>
  <div>
    <input class="input input-primary w-full" type="text" v-model="filter" @keyup="loadUsers">

    <div v-if="modelValue" class="rounded-box p-2 my-2 bg-base-300 shadow-lg">
      <UserCard :user="modelValue">
      </UserCard>
    </div>

    <template v-for="user in users">
      <UserCard :user v-if="user.id != modelValue?.id">
        <button class="btn btn-secondary" @click="emit('update:modelValue', user)">Selecteer</button>
      </UserCard>
    </template>
  </div>
</template>

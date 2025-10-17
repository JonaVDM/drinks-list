<script setup lang="ts">
import type { RecordModel } from 'pocketbase';

defineProps<{ user: RecordModel }>()
</script>

<template>
  <div class="py-4 flex gap-2 justify-between items-center">
    <div class="flex gap-2 items-center">
      <div class="avatar avatar-placeholder">
        <div class="bg-neutral text-neutral-content w-12 rounded-full">
          <span class="text-lg">{{ user.name[0] }}</span>
        </div>
      </div>
      <p class="font-thin">{{ user.name }}</p>
    </div>

    <slot />
  </div>
</template>

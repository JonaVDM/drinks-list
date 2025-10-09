<script setup lang="ts">
const pb = usePocketbase()
const error = ref(false)
const loading = ref(false)
const userStore = useUserStore()

const formdata = ref({
  email: '',
  password: '',
});

async function login() {
  loading.value = true
  error.value = false
  try {
    const { email, password } = formdata.value;
    await pb.collection('users').authWithPassword(email.toLowerCase(), password)
  } catch (e) {
    console.log('error', e)
    error.value = true
  }
  loading.value = false
}
</script>

<template>
  <form @submit.prevent="login" class="flex flex-col gap-2 max-w-lg mx-auto py-12">
    <div v-if="userStore.authed">
      <pre>{{ JSON.stringify(userStore.user, null, 2) }}</pre>
    </div>
    <input type="email" placeholder="email" class="input w-full" v-model="formdata.email" />
    <input type="password" placeholder="password" class="input w-full" v-model="formdata.password" />
    <p v-if="error" class="text-error">failed to log in</p>
    <button class="btn btn-primary" :disabled="loading">Login</button>
  </form>
</template>

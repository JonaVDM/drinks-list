import type { AuthRecord } from "pocketbase";

export const useUserStore = defineStore('users', () => {
  const user = ref<AuthRecord>()
  const authed = ref(false)

  return { user, authed }
});

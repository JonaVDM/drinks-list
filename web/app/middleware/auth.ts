export default defineNuxtRouteMiddleware((to) => {
  const userStore = useUserStore()

  if (userStore.authed) {
    return
  }

  if (to.path != '/login') {
    return navigateTo(`/login?to=${to.fullPath}`)
  }
})

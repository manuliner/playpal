
export default defineNuxtRouteMiddleware((to) => {

  const { loggedIn } = useUserSession()

  if (!loggedIn.value) {
    if (to.name !== "index") {
      return navigateTo('/')
    }
  }


});
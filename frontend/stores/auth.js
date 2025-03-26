import { defineStore } from 'pinia'

export const useAuthStore = defineStore('auth', {
  state: () => ({
    isAuthenticated: false,
    currentUser: null,
  }),
  actions: {
    login(user) {
      this.isAuthenticated = true
      this.currentUser = user
    },
    logout() {
      this.isAuthenticated = false
      this.currentUser = null
    },
  },
})
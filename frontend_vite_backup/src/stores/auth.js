import { defineStore } from 'pinia'

export const useAuthStore = defineStore('auth', {
  state: () => ({
    token: localStorage.getItem('jwt') || null,
  }),
  getters: {
    isAuthenticated: (state) => !!state.token,
    currentUser: (state) => {
      if (!state.token) return null
      try {
        return JSON.parse(atob(state.token.split('.')[1]))
      } catch (e) {
        return null
      }
    }
  },
  actions: {
    setToken(token) {
      this.token = token
      localStorage.setItem('jwt', token)
    },
    logout() {
      this.token = null
      localStorage.removeItem('jwt')
    }
  }
})

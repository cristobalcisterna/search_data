<template>

  <header class="header">
    <a
      v-if="auth.isAuthenticated"
      class="menu-btn"
      @click="menuExpanded = !menuExpanded"
    >
      <img class="menu-icon" src="/img/menu-ico.svg" />
    </a>

    <div @click="$router.push('/')" class="logo">
      <img src="/img/Usach SB.png" />
    </div>

    <nav class="topnav" :aria-expanded="menuExpanded">
      <ul>
        <div v-if="auth.isAuthenticated" class="mobile-only">
          <span style="color:white; margin-left: 1vw;">
            {{ auth.currentUser?.email }}
          </span>
        </div>
      </ul>
    </nav>

    <nav
      v-if="auth.isAuthenticated"
      class="desktop-only usernav"
      :aria-expanded="userExpanded"
    >
      <button class="btn btn-link" @click="userExpanded = !userExpanded">
        Usuario<svg class="svg-icon"><use xlink:href="#chevron-down" /></svg>
      </button>

      <span style="color:white;">{{ auth.currentUser?.email }}</span>
    </nav>
  </header>

  <NuxtPage />
</template>

<script setup>
import { ref } from 'vue'
import { useAuthStore } from '@/stores/auth'


const auth = useAuthStore()
const menuExpanded = ref(false)
const userExpanded = ref(false)
</script>


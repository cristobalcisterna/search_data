<script setup>
import { RouterLink, RouterView } from 'vue-router'
import HelloWorld from './components/HelloWorld.vue'
</script>

<template>
      <header class="header">
      <a
        v-if="isAuthenticated"
        class="menu-btn"
        @click="menuExpanded = !menuExpanded"
      >
        <img class="menu-icon" svg-inline src="@/assets/img/menu-ico.svg" />
      </a>
      <div @click="$router.push('/')" class="logo">
        <img svg-inline src="@/assets/img/Usach SB.png" />
      </div>

      <nav class="topnav" :aria-expanded="menuExpanded">
        <ul>
          <div v-if="isAuthenticated" class="mobile-only">
            <span style="color:white; margin-left: 1vw;">{{
            this.$store.getters.getCurrentUser.email
          }}</span>
          </div>
          <li v-if="isAuthenticated && this.$store.getters.getActivation">
            <router-link to="/resumen">Resumen</router-link>
          </li>

          <li
            v-if="
              isAuthenticated &&
                (this.$store.getters.getCurrentUser.rol === 'admin')
            "
          >
            <router-link to="/usuarios">Usuarios</router-link>
          </li>
          <li
            v-if="
              isAuthenticated &&
                (this.$store.getters.getCurrentUser.rol === 'admin')
            "
          >
            <router-link to="/offices">Oficinas</router-link>
          </li>
          <li
            v-if="
              isAuthenticated &&
                (this.$store.getters.getCurrentUser.rol === 'admin')
            "
          >
            <router-link to="/periodo">Convocatorias</router-link>
          </li>
          <li v-if="isAuthenticated && this.$store.getters.getActivation">
            <router-link to="/" exact>Buzón de reportes</router-link>
          </li>
          <li v-if="isAuthenticated">
            <router-link to="/emails" emails>Correos</router-link>
          </li>
        </ul>

        <user-menu-opts
          v-if="isAuthenticated"
          class="user-menu"
        ></user-menu-opts>
      </nav>

      <nav
        v-if="isAuthenticated"
        class="desktop-only usernav"
        :aria-expanded="userExpanded"
      >
        <button class="btn btn-link" @click="userExpanded = !userExpanded">
          Usuario<svg class="svg-icon"><use xlink:href="#chevron-down" /></svg>
        </button>

        <user-menu-opts @close="userExpanded = false"></user-menu-opts>
          <span style="color:white;">{{
            this.$store.getters.getCurrentUser.email
          }}</span>
      </nav>
      
    </header>

  <RouterView />
</template>


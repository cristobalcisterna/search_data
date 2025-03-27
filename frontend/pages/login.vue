<template>
  <main class="login">
    <form class="form login-form" @submit.prevent="login">
      <h3>Inicio de sesión</h3>
      <label for="login-user">Correo</label>
      <input
        type="text"
        id="login-user"
        v-model="email"
        @keyup.enter="login"
      />
      <label for="login-password">Contraseña</label>
      <input
        type="password"
        id="login-password"
        v-model="password"
        @keyup.enter="login"
      />
      <div class="buttons-container">
        <button type="submit" class="btn btn-primary">
          Iniciar sesión
        </button>
      </div>
    </form>
  </main>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '~/stores/auth'

const email = ref('')
const password = ref('')
const auth = useAuthStore()
const router = useRouter()

const login = async () => {
  try {
    const loginData = {
      email: email.value,
      password: password.value,
    }

    const response = await fetch('http://localhost:8001/api/login', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(loginData),
    })

    if (!response.ok) throw new Error('Credenciales incorrectas')

    const data = await response.json()
    localStorage.setItem('jwt', data.token)

    const payload = JSON.parse(atob(data.token.split('.')[1]))

    if (payload.rol !== 'admin' && payload.rol !== 'evaluador') {
      alert('No puedes acceder con este usuario')
      localStorage.removeItem('jwt')
      return
    }

    auth.login({ email: loginData.email, rol: payload.rol }) // Guarda en el store

    router.push('/search-data')
  } catch (error) {
    console.error(error)
    alert('Credenciales incorrectas')
  }
}
</script>
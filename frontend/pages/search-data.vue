<template>
    <div>
      <h1>Buscador de datos</h1>
  
      <div class="convocatoria-container">
        <h3 style="margin: 0;">Rut:</h3>
        <input class="convocatoria-input" type="text" v-model="auxRut" />
      </div>
  
      <div class="text-center" style="margin-bottom: 1%;">
        <button @click="buscarPersona" class="btn-primary spaced-button">
          Buscar
        </button>
      </div>
  
      <div v-if="data">
        <h2>Datos:</h2>
        <pre>{{ data }}</pre>
      </div>
    </div>
  </template>
  
  <script setup>
  import { ref } from 'vue'
  
  const auxRut = ref('')
  const data = ref(null)
  
  const buscarPersona = async () => {
    if (!auxRut.value) {
      alert('Por favor, ingresa un RUT')
      return
    }
  
    try {
      const response = await fetch(`http://localhost:8001/api/personal-data/${auxRut.value}`, {
        headers: {
          Authorization: `Bearer ${localStorage.getItem('jwt')}`,
          'Content-Type': 'application/json',
        },
      })
  
      if (!response.ok) throw new Error('Datos no encontrados')
  
      const result = await response.json()
      data.value = result.data
    } catch (error) {
      console.error('Error al obtener los datos:', error)
      alert('No se encontraron datos para el RUT proporcionado.')
    }
  }
  </script>


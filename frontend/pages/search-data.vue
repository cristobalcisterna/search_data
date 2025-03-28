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

    <div v-if="PersonalDataDisponible" class="data-card">
  <h2>Datos:</h2>
  <ul>
    <li v-for="(valor, llave) in PersonalDataDisponible" :key="llave">
      <strong>{{ llave }}:</strong> {{ valor }}
    </li>
  </ul>
</div>
  </div>
</template>
  
<script setup>
import { ref } from 'vue'

const auxRut = ref('')

// Datos falsos para pruebas
/*

const PersonalDataDisponible = ref({
  full_name: "Alice Johnson",
  id_number: "98.765.432-1",
  position: "Software Engineer",
  email: "alice.johnson@example.com",
  department: "Engineering"
})
*/
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
    PersonalDataDisponible.value = result.PersonalDataDisponible
  } catch (error) {
    console.error('Error al obtener los datos:', error)
    alert('No se encontraron datos para el RUT proporcionado.')
  }
}
</script>

<style>
.data-card {
  border: 1px solid #ccc;
  background-color: #f9f9f9;
  padding: 16px;
  margin-top: 20px;
  border-radius: 8px;
  max-width: 500px;
  box-shadow: 0 2px 6px rgba(0, 0, 0, 0.1);
}

.data-card ul {
  list-style: none;
  padding: 0;
}

.data-card li {
  margin-bottom: 8px;
}
</style>
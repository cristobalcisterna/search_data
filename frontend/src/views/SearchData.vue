<template>
  <h1>Buscador de datos</h1>

  <!-- Entrada para el RUT -->
  <div class="convocatoria-container">
      <h3 style="margin: 0;">Rut:</h3>
      <input class="convocatoria-input" type="text" v-model="aux_rut">
  </div>

  <!-- Botón de búsqueda -->
  <div class="text-center" style="margin-bottom: 1%;">
      <button @click="buscarPersona" class="btn-primary spaced-button">Buscar</button>
  </div>

  <!-- Muestra los datos obtenidos -->
  <div v-if="data">
      <h2>Datos:</h2>
      <pre>{{ data }}</pre>
  </div>
</template>

<script>
export default {
name: 'SearchData',
data() {
  return {
    aux_rut: null, // El rut que el usuario introduce
    data: null, // Los datos que obtendremos de la base de datos
  };
},

methods: {
  async buscarPersona() {
    try {
      // Asegúrate de que el RUT no esté vacío
      if (!this.aux_rut) {
        alert("Por favor, ingresa un RUT");
        return;
      }

      // Realizamos la solicitud GET a la API para obtener los datos del RUT
      const response = await fetch(`http://localhost:8001/api/students-personal-data/${this.aux_rut}`, {
        method: "GET",
        headers: {
          "Authorization": `Bearer ${localStorage.getItem("jwt")}`, // Incluye el JWT en la cabecera
          "Content-Type": "application/json",
        },
      });

      // Verificamos si la respuesta fue exitosa
      if (!response.ok) {
        throw new Error("Datos no encontrados");
      }

      // Parseamos los datos JSON que nos devuelve el backend
      const result = await response.json();

      // Asignamos los datos al atributo `data`
      this.data = result.data; // Aquí `result.data` es lo que devuelves desde Go

    } catch (error) {
      console.error("Error al obtener los datos:", error);
      alert("No se encontraron datos para el RUT proporcionado.");
    }
  },
},
};
</script>

<template>

<div class="search-container">
    <h1 class="title">Buscador de datos</h1>

    <!-- Sección para ingresar el RUT -->
    <div class="convocatoria-container">
      <h3 style="margin: 0;">Rut:</h3>
      <input class="convocatoria-input" type="text" v-model="aux_rut" placeholder="Ingrese el RUT" />
    </div>

    <!-- Botón para realizar la búsqueda -->
    <div class="text-center" style="margin-bottom: 1%;">
      <button @click="buscarPersona" class="btn-primary spaced-button">Buscar</button>
    </div>

    <!-- Datos encontrados -->
    <div v-if="data" class="result-container">
      <h2>Datos:</h2>
      <pre>{{ data }}</pre>
    </div>

    <!-- 
    <div v-if="aux_data" class="details-container">
      <h3>Datos Encontrados:</h3>
      <p><strong>RUT:</strong> {{ aux_data.rut }}</p>
      <p><strong>Nombre:</strong> {{ aux_data.name }}</p>
      <p><strong>Correo:</strong> {{ aux_data.email }}</p>
      <p><strong>Teléfono:</strong> {{ aux_data.phone }}</p>
      <p><strong>Cargo:</strong> {{ aux_data.work }}</p>
    </div>
-->

  </div>
</template>

<script>
export default {
  name: 'SearchData',
  data: () => ({

    //aux_data: null,
    aux_rut: null,
    data: null,
  }),

  methods: {


    async buscarPersona() {
      try {
        const token = localStorage.getItem("jwt");
        if (!token) {
          alert("Por favor, inicia sesión primero.");
          return;
        }

        const response = await fetch(`http://localhost:8001/api/personal-data/${this.aux_rut}`, {
          method: 'GET',
          headers: {
            'Authorization': `Bearer ${token}`,
            'Content-Type': 'application/json',
          },
        });

        if (!response.ok) {
          throw new Error('Datos no encontrados');
        }

        const data = await response.json();
        this.data = data;
      } catch (error) {
        console.error('Error al obtener los datos:', error);
        alert(error.message);
      }
    },

    /*
    async buscarPersona() {
    try {
      if (!this.aux_rut) {
        alert("Por favor ingresa un RUT.");
        return;
      }
      // Obtener el token JWT del localStorage
      const token = localStorage.getItem("jwt");
      if (!token) {
        alert("Por favor inicie sesión primero.");
        return;
      }
      // Solicitud GET al backend con el RUT ingresado (GetByRut)
      const response = await fetch(`http://localhost:8001/api/personal-data/${this.aux_rut}`, {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
          "Authorization": `Bearer ${token}` // Enviar el token en la cabecera
        },
      });

      // Si la respuesta no es correcta, lanzamos un error
      if (!response.ok) {
        throw new Error("Datos no encontrados");
      }

      // Parseamos la respuesta JSON y asignamos los datos a aux_data
      const data = await response.json();
      this.aux_data = data; // Almacenamos la información recibida

      // Mostrar en consola los datos recibidos
      console.log("Datos recibidos:", this.aux_data);

    } catch (error) {
      // Mostrar un mensaje de error si la consulta falla
      console.error("Error al obtener los datos:", error);
      alert("Error al obtener los datos. Verifica el RUT.");
    }
  },

  */


  },

  

  mounted() {
    console.log("Usuario ha iniciado sesión.");
  }
};
</script>

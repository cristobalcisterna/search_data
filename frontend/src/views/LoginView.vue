
<template>
    <main class="login">
      <form class="form login-form">
        <h3>Inicio de sesión</h3>
        <label for="login-user">Correo</label>
        <input
          type="text"
          id="login-user"
          v-model="email"
          v-on:keyup.enter="login()"
        />
        <label for="login-password">Contraseña</label>
        <input
          type="password"
          id="login-password"
          v-model="password"
          v-on:keyup.enter="login()"
        />
        <div class="buttons-container">
          <button type="button" class="btn btn-primary" @click="login()">
            Iniciar sesión
          </button>
        </div>
        <div class="text-center">
          <router-link to="/modificar-contraseña"
            >Olvidé la contraseña</router-link
          >
        </div>
        <!--<span>¿No tiene una cuenta? <a href="">Regístrese</a></span>-->
      </form>
    </main>
  </template>
  

  <script>
export default {
  data() {
    return {
      email: "",
      password: "",
      errorMsg: "",
    };
  },
  components: {},
  methods: {
    
    login: async function () {
      try {
        let loginData = {
          email: this.email,
          password: this.password,
        };
  
        const response = await fetch("http://localhost:8000/api/login", {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify(loginData),
        });
  
        if (!response.ok) {
          throw new Error("Credenciales incorrectas");
        }
  
        const data = await response.json();
  
        // Guardar token en localStorage
        localStorage.setItem("jwt", data.token);
  
        // Decodificar el token para obtener rol (opcional)
        const payload = JSON.parse(atob(data.token.split('.')[1]));
  
        // Validar rol
        if (
          payload.rol !== "admin" &&
          payload.rol !== "evaluador"
        ) {
          alert("No puedes acceder con este usuario");
          localStorage.removeItem("jwt");
          return;
        }
  
        // Redirigir
        this.$router.push("/");
      } catch (error) {
        console.error(error);
        this.alertError("Credenciales incorrectas o servidor caído");
      }
    }
  }
};
</script>


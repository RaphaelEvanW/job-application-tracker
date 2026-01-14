<template>
    <div class="login">
        <h2> Login </h2>

    <form @submit.prevent="handleLogin">
        <div>
            <input
                type ="email"
                v-model="email"
                placeholder="Email"
                required
            />
        </div>

        <div>
            <input
                type ="password"
                v-model="password"
                placeholder="Password"
                required
            />
        </div>

        <button type="submit">Login</button>
    </form>

    <p>
        Don't Have an Account?
        <router-link to ="/register">Register</router-link>
    </p>
    </div>
</template>

<script setup>
import { ref } from "vue";
import { useRouter } from "vue-router";
import api from "../services/api";

const router = useRouter();

const email = ref("");
const password = ref("");

const handleLogin = async () => {
  try {
    const response = await api.post("/auth/login", {
      email: email.value,
      password: password.value,
    });

    localStorage.setItem("token", response.data.token);
    router.push("/dashboard");

  } catch (error) {
    alert("Login gagal. Periksa email atau password.");
    console.error(error);
  }
};
</script>

<style scoped>
.login {
  max-width: 400px;
  margin: 40px auto;
}
</style>
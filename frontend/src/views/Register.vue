<template>
  <div class="register">
    <h2>Register</h2>

    <form @submit.prevent="handleRegister">
      <div>
        <input
          type="text"
          v-model="name"
          placeholder="Nama"
          required
        />
      </div>

      <div>
        <input
          type="email"
          v-model="email"
          placeholder="Email"
          required
        />
      </div>

      <div>
        <input
          type="password"
          v-model="password"
          placeholder="Password"
          required
        />
      </div>

      <button type="submit">Register</button>
    </form>

    <p>
      Sudah punya akun?
      <router-link to="/login">Login</router-link>
    </p>
  </div>
</template>

<script setup>
import { ref } from "vue";
import { useRouter } from "vue-router";
import api from "../services/api";

const router = useRouter();

const name = ref("");
const email = ref("");
const password = ref("");

const handleRegister = async () => {
  try {
    await api.post("/auth/register", {
      name: name.value,
      email: email.value,
      password: password.value,
    });

    alert("Register berhasil. Silakan login.");
    router.push("/login");

  } catch (error) {
    alert("Register gagal. Email sudah terdaftar.");
    console.error(error);
  }
};
</script>

<style scoped>
.register {
  max-width: 400px;
  margin: 40px auto;
}
</style>

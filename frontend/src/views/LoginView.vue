<template>
  <div class="flex items-center justify-center min-h-screen p-4 bg-slate-100 font-sans text-slate-800">
    <div class="w-full max-w-md bg-white rounded-xl shadow-lg p-8 border border-slate-200">
      <!-- Brand Header -->
      <div class="text-center mb-6">
        <div class="inline-flex items-center justify-center w-12 h-12 bg-indigo-600 text-white rounded-lg font-bold text-xl mb-3 shadow">
          TA
        </div>
        <h2 class="text-2xl font-bold text-slate-900">Selamat Datang Kembali</h2>
        <p class="text-sm text-slate-500 mt-1">Masuk dengan akun Vue 3 + TailAdmin Anda</p>
      </div>

      <!-- Alert -->
      <div v-if="alert.message" :class="[
        'mb-4 p-3 rounded-lg text-sm border',
        alert.isSuccess ? 'bg-emerald-100 text-emerald-800 border-emerald-300' : 'bg-rose-100 text-rose-800 border-rose-300'
      ]">
        {{ alert.message }}
      </div>

      <!-- Form Login Vue 3 -->
      <form @submit.prevent="handleLogin" class="space-y-4">
        <div>
          <label class="block text-sm font-medium text-slate-700 mb-1">Alamat Email</label>
          <input
            v-model="email"
            type="email"
            required
            placeholder="nama@email.com"
            class="w-full px-4 py-2 border border-slate-300 rounded-lg focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 outline-none text-sm"
          />
        </div>

        <div>
          <label class="block text-sm font-medium text-slate-700 mb-1">Password</label>
          <input
            v-model="password"
            type="password"
            required
            placeholder="••••••••"
            class="w-full px-4 py-2 border border-slate-300 rounded-lg focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 outline-none text-sm"
          />
        </div>

        <button
          type="submit"
          :disabled="isLoading"
          class="w-full py-2.5 bg-indigo-600 hover:bg-indigo-700 disabled:opacity-50 text-white font-semibold rounded-lg shadow-md transition duration-200 text-sm"
        >
          {{ isLoading ? 'Memverifikasi...' : 'Masuk (Login)' }}
        </button>
      </form>

      <div class="mt-6 text-center text-sm text-slate-500">
        Belum punya akun?
        <router-link to="/register" class="text-indigo-600 hover:underline font-semibold">Daftar sekarang</router-link>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive } from 'vue'
import { useRouter } from 'vue-router'
import api from '../services/api'

const router = useRouter()
const email = ref('')
const password = ref('')
const isLoading = ref(false)

const alert = reactive({
  message: '',
  isSuccess: false
})

const handleLogin = async () => {
  isLoading.value = true
  alert.message = ''

  try {
    const res = await api.post('/login', {
      email: email.value,
      password: password.value
    })

    if (res.data && res.data.data && res.data.data.token) {
      localStorage.setItem('token', res.data.data.token)
      alert.isSuccess = true
      alert.message = 'Login berhasil! Mengalihkan ke Dashboard...'

      setTimeout(() => {
        router.push('/dashboard')
      }, 800)
    }
  } catch (err) {
    alert.isSuccess = false
    alert.message = err.response?.data?.message || 'Email atau password salah'
  } finally {
    isLoading.value = false
  }
}
</script>

<template>
  <div class="relative flex items-center justify-center min-h-screen p-4 sm:p-6 bg-slate-100 dark:bg-slate-950 font-sans text-slate-800 dark:text-slate-100 antialiased transition-colors duration-200">
    <!-- Top Right Theme Switcher Button -->
    <button
      @click="toggleTheme"
      class="absolute top-4 right-4 p-2.5 rounded-xl bg-white dark:bg-slate-800 border border-slate-200 dark:border-slate-700 text-slate-600 dark:text-slate-300 hover:bg-slate-50 dark:hover:bg-slate-700 shadow-sm transition duration-200 flex items-center space-x-2 text-xs font-semibold"
      :title="isDarkMode ? 'Beralih ke Mode Terang' : 'Beralih ke Mode Gelap'"
    >
      <svg v-if="isDarkMode" class="w-4 h-4 text-amber-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 3v1m0 16v1m9-9h-1M4 12H3m15.364 6.364l-.707-.707M6.343 6.343l-.707-.707m12.728 0l-.707.707M6.343 17.657l-.707.707M16 12a4 4 0 11-8 0 4 4 0 018 0z" />
      </svg>
      <svg v-else class="w-4 h-4 text-slate-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20.354 15.354A9 9 0 018.646 3.646 9.003 9.003 0 0012 21a9.003 9.003 0 008.354-5.646z" />
      </svg>
      <span class="hidden sm:inline">{{ isDarkMode ? 'Light Mode' : 'Dark Mode' }}</span>
    </button>

    <div class="w-full max-w-md bg-white dark:bg-slate-900 rounded-xl shadow-xl p-6 sm:p-8 border border-slate-200 dark:border-slate-800 transition-colors duration-200">
      <!-- Brand Header -->
      <div class="text-center mb-6">
        <div class="inline-flex items-center justify-center w-12 h-12 bg-indigo-600 text-white rounded-lg font-bold text-xl mb-3 shadow">
          TA
        </div>
        <h2 class="text-xl sm:text-2xl font-bold text-slate-900 dark:text-white">Selamat Datang Kembali</h2>
        <p class="text-xs sm:text-sm text-slate-500 dark:text-slate-400 mt-1">Masuk dengan akun Vue 3 + TailAdmin Anda</p>
      </div>

      <!-- Alert -->
      <div v-if="alert.message" :class="[
        'mb-4 p-3 rounded-lg text-sm border',
        alert.isSuccess ? 'bg-emerald-100 dark:bg-emerald-950/50 text-emerald-800 dark:text-emerald-300 border-emerald-300 dark:border-emerald-800' : 'bg-rose-100 dark:bg-rose-950/50 text-rose-800 dark:text-rose-300 border-rose-300 dark:border-rose-800'
      ]">
        {{ alert.message }}
      </div>

      <!-- Form Login Vue 3 -->
      <form @submit.prevent="handleLogin" class="space-y-4">
        <div>
          <label class="block text-xs font-semibold text-slate-700 dark:text-slate-300 uppercase tracking-wider mb-1">Alamat Email</label>
          <input
            v-model="email"
            type="email"
            required
            placeholder="nama@email.com"
            class="w-full px-4 py-2.5 border border-slate-300 dark:border-slate-700 bg-white dark:bg-slate-800 text-slate-900 dark:text-white dark:placeholder-slate-500 rounded-lg focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 outline-none text-sm transition"
          />
        </div>

        <div>
          <div class="flex items-center justify-between mb-1">
            <label class="block text-xs font-semibold text-slate-700 dark:text-slate-300 uppercase tracking-wider">Password</label>
            <router-link to="/forgot-password" class="text-xs text-indigo-600 dark:text-indigo-400 hover:underline font-semibold">
              Lupa password?
            </router-link>
          </div>
          <div class="relative">
            <input
              v-model="password"
              :type="showPassword ? 'text' : 'password'"
              required
              placeholder="••••••••"
              class="w-full px-4 py-2.5 pr-10 border border-slate-300 dark:border-slate-700 bg-white dark:bg-slate-800 text-slate-900 dark:text-white dark:placeholder-slate-500 rounded-lg focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 outline-none text-sm transition"
            />
            <button
              type="button"
              @click="showPassword = !showPassword"
              class="absolute inset-y-0 right-0 flex items-center pr-3 text-slate-400 hover:text-slate-600 dark:hover:text-slate-200 focus:outline-none"
              :title="showPassword ? 'Sembunyikan password' : 'Lihat password'"
            >
              <svg v-if="showPassword" class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13.875 18.825A10.05 10.05 0 0112 19c-4.478 0-8.268-2.943-9.543-7a9.97 9.97 0 011.563-3.029m5.858-5.908a8.956 8.956 0 013.122-.763c4.478 0 8.268 2.943 9.543 7a10.025 10.025 0 01-4.132 5.411m0 0L21 21M3 3l18 18" />
              </svg>
              <svg v-else class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z" />
              </svg>
            </button>
          </div>
        </div>

        <button
          type="submit"
          :disabled="isLoading"
          class="w-full py-3 bg-indigo-600 hover:bg-indigo-700 disabled:opacity-50 text-white font-semibold rounded-lg shadow-md transition duration-200 text-sm"
        >
          {{ isLoading ? 'Memverifikasi...' : 'Masuk (Login)' }}
        </button>
      </form>

      <div class="mt-6 text-center text-sm text-slate-500 dark:text-slate-400">
        Belum punya akun?
        <router-link to="/register" class="text-indigo-600 dark:text-indigo-400 hover:underline font-semibold">Daftar sekarang</router-link>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive } from 'vue'
import { useRouter } from 'vue-router'
import api from '../services/api'
import { useTheme } from '../utils/theme'

const { isDarkMode, toggleTheme } = useTheme()
const router = useRouter()
const email = ref('')
const password = ref('')
const showPassword = ref(false)
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

    if (res.data && res.data.data) {
      const accessToken = res.data.data.access_token || res.data.data.token
      const refreshToken = res.data.data.refresh_token

      localStorage.setItem('token', accessToken)
      localStorage.setItem('access_token', accessToken)
      if (refreshToken) {
        localStorage.setItem('refresh_token', refreshToken)
      }

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

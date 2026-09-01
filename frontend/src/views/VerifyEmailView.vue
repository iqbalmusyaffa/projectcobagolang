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

    <div class="w-full max-w-md bg-white dark:bg-slate-900 rounded-2xl shadow-xl p-6 sm:p-8 border border-slate-200 dark:border-slate-800 transition-colors duration-200">
      <!-- Brand Header -->
      <div class="text-center mb-6">
        <div class="inline-flex items-center justify-center w-14 h-14 bg-gradient-to-tr from-indigo-600 to-violet-600 text-white rounded-2xl font-bold text-2xl mb-3 shadow-lg shadow-indigo-500/30">
          <svg class="w-7 h-7" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 8l7.89 5.26a2 2 0 002.22 0L21 8M5 19h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z" />
          </svg>
        </div>
        <h2 class="text-xl sm:text-2xl font-bold text-slate-900 dark:text-white">Aktivasi Akun Anda</h2>
        <p class="text-xs sm:text-sm text-slate-500 dark:text-slate-400 mt-1">Masukkan kode OTP 6-digit yang dikirimkan ke email</p>
      </div>

      <!-- Alert Notifikasi -->
      <div v-if="alert.message" :class="[
        'mb-5 p-4 rounded-xl text-sm border flex items-start space-x-3 transition-all',
        alert.isSuccess 
          ? 'bg-emerald-50 dark:bg-emerald-950/40 text-emerald-800 dark:text-emerald-300 border-emerald-200 dark:border-emerald-800/60' 
          : 'bg-rose-50 dark:bg-rose-950/40 text-rose-800 dark:text-rose-300 border-rose-200 dark:border-rose-800/60'
      ]">
        <svg v-if="alert.isSuccess" class="w-5 h-5 text-emerald-500 flex-shrink-0 mt-0.5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
        </svg>
        <svg v-else class="w-5 h-5 text-rose-500 flex-shrink-0 mt-0.5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
        </svg>
        <div class="flex-1 text-xs leading-relaxed font-medium">
          {{ alert.message }}
        </div>
      </div>

      <!-- Form Aktivasi Email -->
      <form @submit.prevent="handleVerifyEmail" class="space-y-4">
        <div>
          <label class="block text-xs font-semibold text-slate-700 dark:text-slate-300 uppercase tracking-wider mb-1.5">Alamat Email Terdaftar</label>
          <div class="relative">
            <input
              v-model="email"
              type="email"
              required
              placeholder="nama@email.com"
              class="w-full px-4 py-2.5 pl-10 border border-slate-300 dark:border-slate-700 bg-white dark:bg-slate-800 text-slate-900 dark:text-white dark:placeholder-slate-500 rounded-xl focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 outline-none text-sm transition"
            />
            <div class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none text-slate-400">
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 12a4 4 0 10-8 0 4 4 0 008 0zm0 0v1.5a2.5 2.5 0 005 0V12a9 9 0 10-9 9m4.5-1.206a8.959 8.959 0 01-4.5 1.206" />
              </svg>
            </div>
          </div>
        </div>

        <div>
          <div class="flex items-center justify-between mb-1.5">
            <label class="block text-xs font-semibold text-slate-700 dark:text-slate-300 uppercase tracking-wider">Kode OTP 6-Digit</label>
            <span class="text-xs text-slate-400">Berlaku 15 menit</span>
          </div>
          <div class="relative">
            <input
              v-model="otp"
              type="text"
              maxlength="6"
              required
              placeholder="123456"
              class="w-full px-4 py-3 text-center tracking-[0.5em] font-mono text-xl font-bold border border-slate-300 dark:border-slate-700 bg-slate-50 dark:bg-slate-800 text-slate-900 dark:text-white dark:placeholder-slate-600 rounded-xl focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 outline-none transition"
            />
          </div>
        </div>

        <button
          type="submit"
          :disabled="isLoading || otp.length !== 6"
          class="w-full py-3 bg-gradient-to-r from-indigo-600 to-violet-600 hover:from-indigo-700 hover:to-violet-700 disabled:opacity-50 text-white font-semibold rounded-xl shadow-md shadow-indigo-500/20 transition duration-200 text-sm flex items-center justify-center space-x-2"
        >
          <svg v-if="isLoading" class="animate-spin -ml-1 mr-2 h-4 w-4 text-white" fill="none" viewBox="0 0 24 24">
            <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
            <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
          </svg>
          <span>{{ isLoading ? 'Memverifikasi Akun...' : 'Aktivasi Akun Sekarang' }}</span>
        </button>
      </form>

      <!-- Kirim Ulang OTP Section -->
      <div class="mt-6 pt-5 border-t border-slate-200 dark:border-slate-800 text-center space-y-3">
        <p class="text-xs text-slate-500 dark:text-slate-400">
          Tidak menerima kode verifikasi di email Anda?
        </p>

        <button
          type="button"
          @click="handleResendOTP"
          :disabled="isResending || countdown > 0 || !email"
          class="inline-flex items-center space-x-1.5 text-xs font-semibold text-indigo-600 dark:text-indigo-400 hover:text-indigo-700 dark:hover:text-indigo-300 disabled:opacity-50 disabled:cursor-not-allowed transition"
        >
          <svg v-if="isResending" class="animate-spin h-3.5 w-3.5 text-indigo-600 dark:text-indigo-400" fill="none" viewBox="0 0 24 24">
            <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
            <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
          </svg>
          <svg v-else class="w-3.5 h-3.5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15" />
          </svg>
          <span v-if="countdown > 0">Kirim ulang dalam {{ countdown }}s</span>
          <span v-else>{{ isResending ? 'Mengirim...' : 'Kirim Ulang Kode OTP' }}</span>
        </button>

        <div class="text-xs text-slate-500 dark:text-slate-400 pt-1">
          Sudah terverifikasi?
          <router-link to="/login" class="text-indigo-600 dark:text-indigo-400 hover:underline font-semibold ml-1">Masuk ke Akun</router-link>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted, onUnmounted } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import api from '../services/api'
import { useTheme } from '../utils/theme'

const { isDarkMode, toggleTheme } = useTheme()
const router = useRouter()
const route = useRoute()

const email = ref('')
const otp = ref('')
const isLoading = ref(false)
const isResending = ref(false)
const countdown = ref(0)
let timerInterval = null

const alert = reactive({
  message: '',
  isSuccess: false
})

const startCountdown = (seconds = 60) => {
  countdown.value = seconds
  if (timerInterval) clearInterval(timerInterval)
  timerInterval = setInterval(() => {
    if (countdown.value > 0) {
      countdown.value--
    } else {
      clearInterval(timerInterval)
    }
  }, 1000)
}

const handleVerifyEmail = async () => {
  if (!email.value || otp.value.length !== 6) return
  isLoading.value = true
  alert.message = ''

  try {
    const res = await api.post('/verify-email', {
      email: email.value.trim(),
      otp: otp.value.trim()
    })

    alert.isSuccess = true
    alert.message = res.data.message || 'Akun Anda berhasil diverifikasi! Mengalihkan ke Login...'

    setTimeout(() => {
      router.push({ path: '/login', query: { email: email.value, verified: 'true' } })
    }, 1500)
  } catch (err) {
    alert.isSuccess = false
    alert.message = err.response?.data?.message || 'Kode OTP tidak valid atau sudah kadaluwarsa'
  } finally {
    isLoading.value = false
  }
}

const handleResendOTP = async () => {
  if (!email.value || countdown.value > 0) return
  isResending.value = true
  alert.message = ''

  try {
    const res = await api.post('/resend-verification', {
      email: email.value.trim()
    })

    alert.isSuccess = true
    alert.message = res.data.message || 'Kode OTP baru berhasil dikirimkan ke email Anda!'
    startCountdown(60)
  } catch (err) {
    alert.isSuccess = false
    alert.message = err.response?.data?.message || 'Gagal mengirim ulang kode OTP'
  } finally {
    isResending.value = false
  }
}

onMounted(() => {
  // Tangkap query param jika ada (misal: ?email=...&otp=...)
  if (route.query.email) {
    email.value = String(route.query.email)
  }
  if (route.query.otp) {
    otp.value = String(route.query.otp)
    if (email.value && otp.value.length === 6) {
      // Auto trigger verifikasi jika email & OTP sudah tertera di URL
      handleVerifyEmail()
    }
  }
})

onUnmounted(() => {
  if (timerInterval) clearInterval(timerInterval)
})
</script>

<template>
  <div class="flex min-h-screen bg-slate-100 font-sans text-slate-800 antialiased">
    <!-- Sidebar TailAdmin -->
    <aside class="w-64 bg-slate-900 text-slate-300 flex flex-col justify-between hidden md:flex min-h-screen">
      <div>
        <div class="h-16 flex items-center px-6 border-b border-slate-800">
          <div class="inline-flex items-center justify-center w-8 h-8 bg-indigo-600 text-white rounded-lg font-bold mr-3">
            TA
          </div>
          <span class="text-lg font-bold text-white tracking-wide">TailAdmin Vue</span>
        </div>

        <nav class="p-4 space-y-1">
          <div class="px-3 text-xs font-semibold text-slate-500 uppercase tracking-wider mb-2">Menu Utama</div>
          
          <a href="#" class="flex items-center space-x-3 px-3 py-2.5 rounded-lg bg-slate-800 text-white font-medium text-sm transition">
            <svg class="w-5 h-5 text-indigo-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 12l2-2m0 0l7-7 7 7M5 10v10a1 1 0 001 1h3m10-11l2 2m-2-2v10a1 1 0 01-1 1h-3m-6 0a1 1 0 001-1v-4a1 1 0 011-1h2a1 1 0 011 1v4a1 1 0 001 1m-6 0h6"/>
            </svg>
            <span>Dashboard</span>
          </a>

          <a href="#section-profile" class="flex items-center space-x-3 px-3 py-2.5 rounded-lg text-slate-400 hover:bg-slate-800 hover:text-white font-medium text-sm transition">
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z"/>
            </svg>
            <span>Edit Profil</span>
          </a>

          <a href="#section-password" class="flex items-center space-x-3 px-3 py-2.5 rounded-lg text-slate-400 hover:bg-slate-800 hover:text-white font-medium text-sm transition">
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 15v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2zm10-10V7a4 4 0 00-8 0v4h8z"/>
            </svg>
            <span>Ganti Password</span>
          </a>
        </nav>
      </div>

      <div class="p-4 border-t border-slate-800">
        <button @click="handleLogout" class="w-full flex items-center justify-center space-x-2 py-2 px-4 bg-rose-600/10 text-rose-400 hover:bg-rose-600 hover:text-white rounded-lg transition text-sm font-semibold">
          <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 16l4-4m0 0l-4-4m4 4H7m6 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h4a3 3 0 013 3v1"/>
          </svg>
          <span>Keluar (Logout)</span>
        </button>
      </div>
    </aside>

    <!-- Main Content Wrapper -->
    <div class="flex-1 flex flex-col min-w-0 overflow-hidden">
      <!-- Top Header -->
      <header class="h-16 bg-white border-b border-slate-200 flex items-center justify-between px-6">
        <div class="flex items-center space-x-4">
          <h1 class="text-xl font-bold text-slate-800">Overview Dashboard Vue 3</h1>
        </div>

        <!-- User Info Header with Interactive Dropdown -->
        <div class="relative">
          <div @click.stop="isDropdownOpen = !isDropdownOpen" class="flex items-center space-x-3 cursor-pointer select-none p-1.5 rounded-lg hover:bg-slate-100 transition">
            <div class="text-right hidden sm:block">
              <div class="text-sm font-semibold text-slate-800">{{ user.name || 'Loading...' }}</div>
              <div class="text-xs text-slate-500">{{ user.email || 'Loading...' }}</div>
            </div>
            <div class="w-10 h-10 rounded-full bg-indigo-600 text-white font-bold flex items-center justify-center shadow">
              {{ avatarInitial }}
            </div>
            <svg class="w-4 h-4 text-slate-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7"/>
            </svg>
          </div>

          <!-- Dropdown Menu Vue 3 -->
          <div v-if="isDropdownOpen" class="absolute right-0 mt-2 w-56 bg-white rounded-xl shadow-lg border border-slate-200 py-2 z-50">
            <div class="px-4 py-2 border-b border-slate-100">
              <p class="text-sm font-bold text-slate-800 truncate">{{ user.name }}</p>
              <p class="text-xs text-slate-500 truncate">{{ user.email }}</p>
            </div>

            <a href="#section-profile" @click="isDropdownOpen = false" class="flex items-center px-4 py-2 text-sm text-slate-700 hover:bg-slate-50 hover:text-indigo-600 transition">
              <svg class="w-4 h-4 mr-2.5 text-slate-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z"/>
              </svg>
              Edit Profil
            </a>

            <a href="#section-password" @click="isDropdownOpen = false" class="flex items-center px-4 py-2 text-sm text-slate-700 hover:bg-slate-50 hover:text-indigo-600 transition">
              <svg class="w-4 h-4 mr-2.5 text-slate-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 15v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2zm10-10V7a4 4 0 00-8 0v4h8z"/>
              </svg>
              Ganti Password
            </a>

            <div class="border-t border-slate-100 my-1"></div>

            <button @click="handleLogout" class="w-full text-left flex items-center px-4 py-2 text-sm text-rose-600 hover:bg-rose-50 font-medium transition">
              <svg class="w-4 h-4 mr-2.5 text-rose-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 16l4-4m0 0l-4-4m4 4H7m6 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h4a3 3 0 013 3v1"/>
              </svg>
              Keluar (Logout)
            </button>
          </div>
        </div>
      </header>

      <!-- Main Body -->
      <main class="flex-1 overflow-y-auto p-6 space-y-6">
        <!-- Reaktif Alert Box -->
        <div v-if="alert.message" :class="[
          'p-4 rounded-xl text-sm font-medium shadow-sm transition',
          alert.isSuccess ? 'bg-emerald-100 text-emerald-800 border border-emerald-300' : 'bg-rose-100 text-rose-800 border border-rose-300'
        ]">
          {{ alert.message }}
        </div>

        <!-- Stat Cards Grid -->
        <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
          <div class="bg-white p-6 rounded-xl border border-slate-200 shadow-sm flex items-center space-x-4">
            <div class="p-3 bg-indigo-50 text-indigo-600 rounded-lg">
              <svg class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4.354a4 4 0 110 5.292M15 21H3v-1a6 6 0 0112 0v1zm0 0h6v-1a6 6 0 00-9-5.197M13 7a4 4 0 11-8 0 4 4 0 018 0z"/>
              </svg>
            </div>
            <div>
              <p class="text-xs font-medium text-slate-500 uppercase tracking-wider">Frontend Engine</p>
              <h3 class="text-xl font-bold text-emerald-600 mt-1">Vue 3 + Vite (npm)</h3>
            </div>
          </div>

          <div class="bg-white p-6 rounded-xl border border-slate-200 shadow-sm flex items-center space-x-4">
            <div class="p-3 bg-emerald-50 text-emerald-600 rounded-lg">
              <svg class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m5.618-4.016A11.955 11.955 0 0112 2.944a11.955 11.955 0 01-8.618 3.04A12.02 12.02 0 003 9c0 5.591 3.824 10.29 9 11.622 5.176-1.332 9-6.03 9-11.622 0-1.042-.133-2.052-.382-3.016z"/>
              </svg>
            </div>
            <div>
              <p class="text-xs font-medium text-slate-500 uppercase tracking-wider">Autentikasi</p>
              <h3 class="text-xl font-bold text-slate-800 mt-1">JWT Bearer Axios</h3>
            </div>
          </div>

          <div class="bg-white p-6 rounded-xl border border-slate-200 shadow-sm flex items-center space-x-4">
            <div class="p-3 bg-blue-50 text-blue-600 rounded-lg">
              <svg class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 7v10c0 2.21 3.582 4 8 4s8-1.79 8-4V7M4 7c0 2.21 3.582 4 8 4s8-1.79 8-4M4 7c0-2.21 3.582-4 8-4s8 1.79 8 4"/>
              </svg>
            </div>
            <div>
              <p class="text-xs font-medium text-slate-500 uppercase tracking-wider">Database</p>
              <h3 class="text-xl font-bold text-slate-800 mt-1">PostgreSQL</h3>
            </div>
          </div>
        </div>

        <!-- User Profile Information Card -->
        <div id="section-profile" class="bg-white rounded-xl border border-slate-200 shadow-sm overflow-hidden">
          <div class="px-6 py-4 border-b border-slate-200 bg-slate-50 flex items-center justify-between">
            <h2 class="text-lg font-bold text-slate-900">Informasi Pengguna (Profile Vue 3)</h2>
            <span class="px-3 py-1 bg-indigo-100 text-indigo-700 text-xs font-semibold rounded-full">Axios REST Data</span>
          </div>

          <div class="p-6">
            <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
              <div>
                <label class="block text-xs font-semibold text-slate-400 uppercase tracking-wider mb-1">User ID</label>
                <p class="text-base font-bold text-slate-800 bg-slate-100 px-4 py-2 rounded-lg border border-slate-200">{{ user.id || '-' }}</p>
              </div>

              <div>
                <label class="block text-xs font-semibold text-slate-400 uppercase tracking-wider mb-1">Nama Lengkap</label>
                <p class="text-base font-bold text-slate-800 bg-slate-100 px-4 py-2 rounded-lg border border-slate-200">{{ user.name || '-' }}</p>
              </div>

              <div>
                <label class="block text-xs font-semibold text-slate-400 uppercase tracking-wider mb-1">Email</label>
                <p class="text-base font-bold text-slate-800 bg-slate-100 px-4 py-2 rounded-lg border border-slate-200">{{ user.email || '-' }}</p>
              </div>

              <div>
                <label class="block text-xs font-semibold text-slate-400 uppercase tracking-wider mb-1">Tanggal Terdaftar</label>
                <p class="text-base font-bold text-slate-800 bg-slate-100 px-4 py-2 rounded-lg border border-slate-200">{{ formattedCreatedAt }}</p>
              </div>
            </div>
          </div>
        </div>

        <!-- Action Forms Section: Edit Profil & Ganti Password -->
        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <!-- Form Edit Profil Card -->
          <div class="bg-white rounded-xl border border-slate-200 shadow-sm overflow-hidden flex flex-col justify-between">
            <div>
              <div class="px-6 py-4 border-b border-slate-200 bg-slate-50">
                <h3 class="text-lg font-bold text-slate-900">Edit Profil (Vue 3)</h3>
                <p class="text-xs text-slate-500 mt-0.5">Perbarui nama dan alamat email reaktif</p>
              </div>

              <form @submit.prevent="handleUpdateProfile" class="p-6 space-y-4">
                <div>
                  <label class="block text-xs font-semibold text-slate-700 uppercase tracking-wider mb-1">Nama Lengkap</label>
                  <input
                    v-model="editForm.name"
                    type="text"
                    required
                    class="w-full px-4 py-2 border border-slate-300 rounded-lg focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 outline-none text-sm"
                  />
                </div>

                <div>
                  <label class="block text-xs font-semibold text-slate-700 uppercase tracking-wider mb-1">Alamat Email</label>
                  <input
                    v-model="editForm.email"
                    type="email"
                    required
                    class="w-full px-4 py-2 border border-slate-300 rounded-lg focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 outline-none text-sm"
                  />
                </div>

                <button
                  type="submit"
                  :disabled="isUpdating"
                  class="w-full py-2.5 bg-indigo-600 hover:bg-indigo-700 disabled:opacity-50 text-white font-semibold rounded-lg shadow-md transition duration-200 text-sm"
                >
                  {{ isUpdating ? 'Menyimpan...' : 'Simpan Perubahan Profil' }}
                </button>
              </form>
            </div>
          </div>

          <!-- Form Ganti Password Card -->
          <div id="section-password" class="bg-white rounded-xl border border-slate-200 shadow-sm overflow-hidden flex flex-col justify-between">
            <div>
              <div class="px-6 py-4 border-b border-slate-200 bg-slate-50">
                <h3 class="text-lg font-bold text-slate-900">Ganti Password (Vue 3)</h3>
                <p class="text-xs text-slate-500 mt-0.5">Perbarui kata sandi akun Anda</p>
              </div>

              <form @submit.prevent="handleChangePassword" class="p-6 space-y-4">
                <div>
                  <label class="block text-xs font-semibold text-slate-700 uppercase tracking-wider mb-1">Password Saat Ini</label>
                  <input
                    v-model="passwordForm.old_password"
                    type="password"
                    required
                    placeholder="••••••••"
                    class="w-full px-4 py-2 border border-slate-300 rounded-lg focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 outline-none text-sm"
                  />
                </div>

                <div>
                  <label class="block text-xs font-semibold text-slate-700 uppercase tracking-wider mb-1">Password Baru</label>
                  <input
                    v-model="passwordForm.new_password"
                    type="password"
                    required
                    placeholder="••••••••"
                    class="w-full px-4 py-2 border border-slate-300 rounded-lg focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 outline-none text-sm"
                  />
                </div>

                <button
                  type="submit"
                  :disabled="isChangingPassword"
                  class="w-full py-2.5 bg-slate-800 hover:bg-slate-900 disabled:opacity-50 text-white font-semibold rounded-lg shadow-md transition duration-200 text-sm"
                >
                  {{ isChangingPassword ? 'Memproses...' : 'Perbarui Password' }}
                </button>
              </form>
            </div>
          </div>
        </div>
      </main>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive, computed, onMounted, onUnmounted } from 'vue'
import { useRouter } from 'vue-router'
import api from '../services/api'

const router = useRouter()
const isDropdownOpen = ref(false)
const isUpdating = ref(false)
const isChangingPassword = ref(false)

const user = reactive({
  id: '',
  name: '',
  email: '',
  created_at: ''
})

const editForm = reactive({
  name: '',
  email: ''
})

const passwordForm = reactive({
  old_password: '',
  new_password: ''
})

const alert = reactive({
  message: '',
  isSuccess: true
})

const avatarInitial = computed(() => {
  return user.name ? user.name.charAt(0).toUpperCase() : 'U'
})

const formattedCreatedAt = computed(() => {
  return user.created_at ? new Date(user.created_at).toLocaleString('id-ID') : '-'
})

const showAlert = (message, isSuccess = true) => {
  alert.message = message
  alert.isSuccess = isSuccess

  setTimeout(() => {
    alert.message = ''
  }, 4000)
}

const loadUserProfile = async () => {
  try {
    const res = await api.get('/profile')
    const data = res.data.data

    user.id = data.id
    user.name = data.name
    user.email = data.email
    user.created_at = data.created_at

    editForm.name = data.name
    editForm.email = data.email
  } catch (err) {
    localStorage.removeItem('token')
    router.push('/login')
  }
}

const handleUpdateProfile = async () => {
  isUpdating.value = true
  try {
    const res = await api.put('/profile', {
      name: editForm.name,
      email: editForm.email
    })

    user.name = res.data.data.name
    user.email = res.data.data.email
    showAlert('Profil Vue 3 Anda berhasil diperbarui!', true)
  } catch (err) {
    showAlert(err.response?.data?.message || 'Gagal memperbarui profil', false)
  } finally {
    isUpdating.value = false
  }
}

const handleChangePassword = async () => {
  isChangingPassword.value = true
  try {
    await api.put('/change-password', {
      old_password: passwordForm.old_password,
      new_password: passwordForm.new_password
    })

    passwordForm.old_password = ''
    passwordForm.new_password = ''
    showAlert('Password berhasil diperbarui!', true)
  } catch (err) {
    showAlert(err.response?.data?.message || 'Gagal mengganti password', false)
  } finally {
    isChangingPassword.value = false
  }
}

const handleLogout = () => {
  localStorage.removeItem('token')
  router.push('/login')
}

const closeDropdownOnOutsideClick = () => {
  isDropdownOpen.value = false
}

onMounted(() => {
  loadUserProfile()
  window.addEventListener('click', closeDropdownOnOutsideClick)
})

onUnmounted(() => {
  window.removeEventListener('click', closeDropdownOnOutsideClick)
})
</script>

<template>
  <div class="flex min-h-screen bg-slate-100 font-sans text-slate-800 antialiased">
    <!-- Desktop Sidebar TailAdmin -->
    <aside class="w-64 bg-slate-900 text-slate-300 flex flex-col justify-between hidden md:flex min-h-screen sticky top-0 h-screen">
      <div>
        <div class="h-16 flex items-center px-6 border-b border-slate-800">
          <div class="inline-flex items-center justify-center w-8 h-8 bg-indigo-600 text-white rounded-lg font-bold mr-3 shadow">
            TA
          </div>
          <span class="text-lg font-bold text-white tracking-wide">TailAdmin Vue</span>
        </div>

        <nav class="p-4 space-y-1">
          <div class="px-3 text-xs font-semibold text-slate-500 uppercase tracking-wider mb-2">Menu Utama</div>
          
          <a href="#" @click.prevent="activeTab = 'dashboard'" :class="[
            'flex items-center space-x-3 px-3 py-2.5 rounded-lg font-medium text-sm transition',
            activeTab === 'dashboard' ? 'bg-slate-800 text-white shadow-sm' : 'text-slate-400 hover:bg-slate-800 hover:text-white'
          ]">
            <svg class="w-5 h-5 text-indigo-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 12l2-2m0 0l7-7 7 7M5 10v10a1 1 0 001 1h3m10-11l2 2m-2-2v10a1 1 0 01-1 1h-3m-6 0a1 1 0 001-1v-4a1 1 0 011-1h2a1 1 0 011 1v4a1 1 0 001 1m-6 0h6"/>
            </svg>
            <span>Dashboard</span>
          </a>

          <!-- Menu Khusus Superadmin & Owner -->
          <a v-if="canManageUsers" href="#" @click.prevent="activeTab = 'users'" :class="[
            'flex items-center space-x-3 px-3 py-2.5 rounded-lg font-medium text-sm transition',
            activeTab === 'users' ? 'bg-slate-800 text-white shadow-sm' : 'text-slate-400 hover:bg-slate-800 hover:text-white'
          ]">
            <svg class="w-5 h-5 text-purple-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4.354a4 4 0 110 5.292M15 21H3v-1a6 6 0 0112 0v1zm0 0h6v-1a6 6 0 00-9-5.197M13 7a4 4 0 11-8 0 4 4 0 018 0z"/>
            </svg>
            <span>Manajemen User</span>
          </a>

          <a href="#section-profile" @click="activeTab = 'dashboard'" class="flex items-center space-x-3 px-3 py-2.5 rounded-lg text-slate-400 hover:bg-slate-800 hover:text-white font-medium text-sm transition">
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z"/>
            </svg>
            <span>Edit Profil</span>
          </a>

          <a href="#section-password" @click="activeTab = 'dashboard'" class="flex items-center space-x-3 px-3 py-2.5 rounded-lg text-slate-400 hover:bg-slate-800 hover:text-white font-medium text-sm transition">
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

    <!-- Mobile Drawer Off-Canvas Sidebar (Tampil di Layar Smartphone/Tablet <768px) -->
    <div v-if="isMobileSidebarOpen" class="fixed inset-0 z-50 flex md:hidden">
      <!-- Backdrop Overlay dengan Blur -->
      <div @click="isMobileSidebarOpen = false" class="fixed inset-0 bg-slate-900/60 backdrop-blur-sm transition-opacity"></div>

      <!-- Drawer Content -->
      <div class="relative flex flex-col w-72 max-w-full bg-slate-900 text-slate-300 shadow-2xl h-full z-10 animate-in slide-in-from-left duration-200">
        <div class="h-16 flex items-center justify-between px-6 border-b border-slate-800">
          <div class="flex items-center space-x-3">
            <div class="inline-flex items-center justify-center w-8 h-8 bg-indigo-600 text-white rounded-lg font-bold shadow">
              TA
            </div>
            <span class="text-lg font-bold text-white tracking-wide">TailAdmin Vue</span>
          </div>
          <button @click="isMobileSidebarOpen = false" class="text-slate-400 hover:text-white p-1 rounded-lg">
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"/>
            </svg>
          </button>
        </div>

        <nav class="p-4 space-y-1 flex-1 overflow-y-auto">
          <div class="px-3 text-xs font-semibold text-slate-500 uppercase tracking-wider mb-2">Menu Utama</div>
          
          <a href="#" @click.prevent="activeTab = 'dashboard'; isMobileSidebarOpen = false" :class="[
            'flex items-center space-x-3 px-3 py-3 rounded-lg font-medium text-sm transition',
            activeTab === 'dashboard' ? 'bg-slate-800 text-white shadow-sm' : 'text-slate-400 hover:bg-slate-800 hover:text-white'
          ]">
            <svg class="w-5 h-5 text-indigo-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 12l2-2m0 0l7-7 7 7M5 10v10a1 1 0 001 1h3m10-11l2 2m-2-2v10a1 1 0 01-1 1h-3m-6 0a1 1 0 001-1v-4a1 1 0 011-1h2a1 1 0 011 1v4a1 1 0 001 1m-6 0h6"/>
            </svg>
            <span>Dashboard</span>
          </a>

          <!-- Menu Khusus Superadmin & Owner -->
          <a v-if="canManageUsers" href="#" @click.prevent="activeTab = 'users'; isMobileSidebarOpen = false" :class="[
            'flex items-center space-x-3 px-3 py-3 rounded-lg font-medium text-sm transition',
            activeTab === 'users' ? 'bg-slate-800 text-white shadow-sm' : 'text-slate-400 hover:bg-slate-800 hover:text-white'
          ]">
            <svg class="w-5 h-5 text-purple-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4.354a4 4 0 110 5.292M15 21H3v-1a6 6 0 0112 0v1zm0 0h6v-1a6 6 0 00-9-5.197M13 7a4 4 0 11-8 0 4 4 0 018 0z"/>
            </svg>
            <span>Manajemen User</span>
          </a>

          <a href="#section-profile" @click="activeTab = 'dashboard'; isMobileSidebarOpen = false" class="flex items-center space-x-3 px-3 py-3 rounded-lg text-slate-400 hover:bg-slate-800 hover:text-white font-medium text-sm transition">
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z"/>
            </svg>
            <span>Edit Profil</span>
          </a>

          <a href="#section-password" @click="activeTab = 'dashboard'; isMobileSidebarOpen = false" class="flex items-center space-x-3 px-3 py-3 rounded-lg text-slate-400 hover:bg-slate-800 hover:text-white font-medium text-sm transition">
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 15v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2zm10-10V7a4 4 0 00-8 0v4h8z"/>
            </svg>
            <span>Ganti Password</span>
          </a>
        </nav>

        <div class="p-4 border-t border-slate-800">
          <button @click="handleLogout" class="w-full flex items-center justify-center space-x-2 py-3 px-4 bg-rose-600/10 text-rose-400 hover:bg-rose-600 hover:text-white rounded-lg transition text-sm font-semibold">
            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 16l4-4m0 0l-4-4m4 4H7m6 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h4a3 3 0 013 3v1"/>
            </svg>
            <span>Keluar (Logout)</span>
          </button>
        </div>
      </div>
    </div>

    <!-- Main Content Wrapper -->
    <div class="flex-1 flex flex-col min-w-0 overflow-hidden">
      <!-- Top Header Responsif -->
      <header class="h-16 bg-white border-b border-slate-200 flex items-center justify-between px-4 sm:px-6 sticky top-0 z-30 shadow-sm">
        <div class="flex items-center space-x-3">
          <!-- Hamburger Button khusus Mobile/Tablet (<768px) -->
          <button @click="isMobileSidebarOpen = true" class="md:hidden p-2 rounded-lg text-slate-600 hover:bg-slate-100 focus:outline-none transition">
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h16"/>
            </svg>
          </button>

          <h1 class="text-base sm:text-xl font-bold text-slate-800 truncate">
            {{ activeTab === 'users' ? 'Manajemen Pengguna (RBAC)' : 'Overview Dashboard Vue 3' }}
          </h1>
        </div>

        <!-- User Info Header dengan Responsive Dropdown -->
        <div class="relative">
          <div @click.stop="isDropdownOpen = !isDropdownOpen" class="flex items-center space-x-2 sm:space-x-3 cursor-pointer select-none p-1.5 rounded-lg hover:bg-slate-100 transition">
            <div class="text-right hidden sm:block">
              <div class="flex items-center space-x-2 justify-end">
                <span class="text-sm font-semibold text-slate-800">{{ user.name || 'Loading...' }}</span>
                <!-- Role Badge Header -->
                <span :class="roleBadgeClass(user.role)">{{ (user.role || 'admin').toUpperCase() }}</span>
              </div>
              <div class="text-xs text-slate-500">{{ user.email || 'Loading...' }}</div>
            </div>
            <div class="w-9 h-9 sm:w-10 sm:h-10 rounded-full bg-indigo-600 text-white font-bold flex items-center justify-center shadow text-sm sm:text-base">
              {{ avatarInitial }}
            </div>
            <svg class="w-4 h-4 text-slate-400 hidden sm:block" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7"/>
            </svg>
          </div>

          <!-- Dropdown Menu Vue 3 -->
          <div v-if="isDropdownOpen" class="absolute right-0 mt-2 w-60 bg-white rounded-xl shadow-xl border border-slate-200 py-2 z-50 animate-in fade-in zoom-in duration-150">
            <div class="px-4 py-2 border-b border-slate-100">
              <p class="text-sm font-bold text-slate-800 truncate">{{ user.name }}</p>
              <p class="text-xs text-slate-500 truncate mb-1">{{ user.email }}</p>
              <span :class="roleBadgeClass(user.role)">Peran: {{ (user.role || 'admin').toUpperCase() }}</span>
            </div>

            <a href="#section-profile" @click="activeTab = 'dashboard'; isDropdownOpen = false" class="flex items-center px-4 py-2 text-sm text-slate-700 hover:bg-slate-50 hover:text-indigo-600 transition">
              <svg class="w-4 h-4 mr-2.5 text-slate-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z"/>
              </svg>
              Edit Profil
            </a>

            <a href="#section-password" @click="activeTab = 'dashboard'; isDropdownOpen = false" class="flex items-center px-4 py-2 text-sm text-slate-700 hover:bg-slate-50 hover:text-indigo-600 transition">
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

      <!-- Main Body Responsif (HP, Tablet, Laptop) -->
      <main class="flex-1 overflow-y-auto p-4 sm:p-6 lg:p-8 space-y-6">
        <!-- Reaktif Alert Box -->
        <div v-if="alert.message" :class="[
          'p-4 rounded-xl text-sm font-medium shadow-sm transition',
          alert.isSuccess ? 'bg-emerald-100 text-emerald-800 border border-emerald-300' : 'bg-rose-100 text-rose-800 border border-rose-300'
        ]">
          {{ alert.message }}
        </div>

        <!-- TAB 1: MAIN DASHBOARD -->
        <template v-if="activeTab === 'dashboard'">
          <!-- Stat Cards Grid Adaptif -->
          <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-4 sm:gap-6">
            <div class="bg-white p-5 sm:p-6 rounded-xl border border-slate-200 shadow-sm flex items-center space-x-4">
              <div class="p-3 bg-indigo-50 text-indigo-600 rounded-lg shrink-0">
                <svg class="w-7 h-7 sm:w-8 sm:h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4.354a4 4 0 110 5.292M15 21H3v-1a6 6 0 0112 0v1zm0 0h6v-1a6 6 0 00-9-5.197M13 7a4 4 0 11-8 0 4 4 0 018 0z"/>
                </svg>
              </div>
              <div class="min-w-0">
                <p class="text-xs font-medium text-slate-500 uppercase tracking-wider truncate">Peran Anda (RBAC)</p>
                <h3 class="text-lg sm:text-xl font-bold text-indigo-600 mt-0.5 uppercase truncate">{{ user.role || 'ADMIN' }}</h3>
              </div>
            </div>

            <div class="bg-white p-5 sm:p-6 rounded-xl border border-slate-200 shadow-sm flex items-center space-x-4">
              <div class="p-3 bg-emerald-50 text-emerald-600 rounded-lg shrink-0">
                <svg class="w-7 h-7 sm:w-8 sm:h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m5.618-4.016A11.955 11.955 0 0112 2.944a11.955 11.955 0 01-8.618 3.04A12.02 12.02 0 003 9c0 5.591 3.824 10.29 9 11.622 5.176-1.332 9-6.03 9-11.622 0-1.042-.133-2.052-.382-3.016z"/>
                </svg>
              </div>
              <div class="min-w-0">
                <p class="text-xs font-medium text-slate-500 uppercase tracking-wider truncate">Keamanan Redis & JWT</p>
                <h3 class="text-lg sm:text-xl font-bold text-slate-800 mt-0.5 truncate">Dual Token Auth</h3>
              </div>
            </div>

            <div class="bg-white p-5 sm:p-6 rounded-xl border border-slate-200 shadow-sm flex items-center space-x-4 sm:col-span-2 lg:col-span-1">
              <div class="p-3 bg-blue-50 text-blue-600 rounded-lg shrink-0">
                <svg class="w-7 h-7 sm:w-8 sm:h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 7v10c0 2.21 3.582 4 8 4s8-1.79 8-4V7M4 7c0 2.21 3.582 4 8 4s8-1.79 8-4M4 7c0-2.21 3.582-4 8-4s8 1.79 8 4"/>
                </svg>
              </div>
              <div class="min-w-0">
                <p class="text-xs font-medium text-slate-500 uppercase tracking-wider truncate">Database</p>
                <h3 class="text-lg sm:text-xl font-bold text-slate-800 mt-0.5 truncate">PostgreSQL</h3>
              </div>
            </div>
          </div>

          <!-- User Profile Information Card -->
          <div id="section-profile" class="bg-white rounded-xl border border-slate-200 shadow-sm overflow-hidden">
            <div class="px-5 sm:px-6 py-4 border-b border-slate-200 bg-slate-50 flex flex-col sm:flex-row items-start sm:items-center justify-between gap-2">
              <h2 class="text-base sm:text-lg font-bold text-slate-900">Informasi Pengguna (Profile Vue 3)</h2>
              <span :class="roleBadgeClass(user.role)">Peran: {{ (user.role || 'admin').toUpperCase() }}</span>
            </div>

            <div class="p-5 sm:p-6">
              <div class="grid grid-cols-1 sm:grid-cols-2 gap-4 sm:gap-6">
                <div>
                  <label class="block text-xs font-semibold text-slate-400 uppercase tracking-wider mb-1">User ID</label>
                  <p class="text-sm sm:text-base font-bold text-slate-800 bg-slate-100 px-4 py-2 rounded-lg border border-slate-200 truncate">{{ user.id || '-' }}</p>
                </div>

                <div>
                  <label class="block text-xs font-semibold text-slate-400 uppercase tracking-wider mb-1">Nama Lengkap</label>
                  <p class="text-sm sm:text-base font-bold text-slate-800 bg-slate-100 px-4 py-2 rounded-lg border border-slate-200 truncate">{{ user.name || '-' }}</p>
                </div>

                <div>
                  <label class="block text-xs font-semibold text-slate-400 uppercase tracking-wider mb-1">Email</label>
                  <p class="text-sm sm:text-base font-bold text-slate-800 bg-slate-100 px-4 py-2 rounded-lg border border-slate-200 truncate">{{ user.email || '-' }}</p>
                </div>

                <div>
                  <label class="block text-xs font-semibold text-slate-400 uppercase tracking-wider mb-1">Hak Akses / Role</label>
                  <p class="text-sm sm:text-base font-bold text-indigo-700 bg-indigo-50 px-4 py-2 rounded-lg border border-indigo-200 uppercase truncate">{{ user.role || 'ADMIN' }}</p>
                </div>
              </div>
            </div>
          </div>

          <!-- Action Forms Section: Edit Profil & Ganti Password -->
          <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
            <!-- Form Edit Profil Card -->
            <div class="bg-white rounded-xl border border-slate-200 shadow-sm overflow-hidden flex flex-col justify-between">
              <div>
                <div class="px-5 sm:px-6 py-4 border-b border-slate-200 bg-slate-50">
                  <h3 class="text-base sm:text-lg font-bold text-slate-900">Edit Profil</h3>
                  <p class="text-xs text-slate-500 mt-0.5">Perbarui nama dan alamat email reaktif</p>
                </div>

                <form @submit.prevent="handleUpdateProfile" class="p-5 sm:p-6 space-y-4">
                  <div>
                    <label class="block text-xs font-semibold text-slate-700 uppercase tracking-wider mb-1">Nama Lengkap</label>
                    <input
                      v-model="editForm.name"
                      type="text"
                      required
                      class="w-full px-4 py-2.5 border border-slate-300 rounded-lg focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 outline-none text-sm"
                    />
                  </div>

                  <div>
                    <label class="block text-xs font-semibold text-slate-700 uppercase tracking-wider mb-1">Alamat Email</label>
                    <input
                      v-model="editForm.email"
                      type="email"
                      required
                      class="w-full px-4 py-2.5 border border-slate-300 rounded-lg focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 outline-none text-sm"
                    />
                  </div>

                  <button
                    type="submit"
                    :disabled="isUpdating"
                    class="w-full py-3 bg-indigo-600 hover:bg-indigo-700 disabled:opacity-50 text-white font-semibold rounded-lg shadow-md transition duration-200 text-sm"
                  >
                    {{ isUpdating ? 'Menyimpan...' : 'Simpan Perubahan Profil' }}
                  </button>
                </form>
              </div>
            </div>

            <!-- Form Ganti Password Card -->
            <div id="section-password" class="bg-white rounded-xl border border-slate-200 shadow-sm overflow-hidden flex flex-col justify-between">
              <div>
                <div class="px-5 sm:px-6 py-4 border-b border-slate-200 bg-slate-50">
                  <h3 class="text-base sm:text-lg font-bold text-slate-900">Ganti Password</h3>
                  <p class="text-xs text-slate-500 mt-0.5">Perbarui kata sandi akun Anda</p>
                </div>

                <form @submit.prevent="handleChangePassword" class="p-5 sm:p-6 space-y-4">
                  <div>
                    <label class="block text-xs font-semibold text-slate-700 uppercase tracking-wider mb-1">Password Saat Ini</label>
                    <input
                      v-model="passwordForm.old_password"
                      type="password"
                      required
                      placeholder="••••••••"
                      class="w-full px-4 py-2.5 border border-slate-300 rounded-lg focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 outline-none text-sm"
                    />
                  </div>

                  <div>
                    <label class="block text-xs font-semibold text-slate-700 uppercase tracking-wider mb-1">Password Baru</label>
                    <input
                      v-model="passwordForm.new_password"
                      type="password"
                      required
                      placeholder="••••••••"
                      class="w-full px-4 py-2.5 border border-slate-300 rounded-lg focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 outline-none text-sm"
                    />
                  </div>

                  <button
                    type="submit"
                    :disabled="isChangingPassword"
                    class="w-full py-3 bg-slate-800 hover:bg-slate-900 disabled:opacity-50 text-white font-semibold rounded-lg shadow-md transition duration-200 text-sm"
                  >
                    {{ isChangingPassword ? 'Memproses...' : 'Perbarui Password' }}
                  </button>
                </form>
              </div>
            </div>
          </div>
        </template>

        <!-- TAB 2: MANAJEMEN USER (KHUSUS SUPERADMIN & OWNER) -->
        <template v-if="activeTab === 'users' && canManageUsers">
          <div class="bg-white rounded-xl border border-slate-200 shadow-sm overflow-hidden">
            <div class="px-5 sm:px-6 py-4 border-b border-slate-200 bg-slate-50 flex flex-col sm:flex-row sm:items-center justify-between gap-3">
              <div>
                <h2 class="text-base sm:text-lg font-bold text-slate-900">Manajemen Seluruh Pengguna Sistem</h2>
                <p class="text-xs text-slate-500">Khusus Peran Superadmin & Owner (`GET /api/admin/users`)</p>
              </div>
              <div class="flex items-center space-x-2.5 self-start sm:self-auto">
                <button v-if="isSuperadmin" @click="isAddUserModalOpen = true" class="px-3.5 py-2 bg-indigo-600 hover:bg-indigo-700 text-white text-xs font-semibold rounded-lg shadow transition flex items-center space-x-1.5">
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4"/>
                  </svg>
                  <span>Tambah User Baru</span>
                </button>
                <button @click="loadAllUsers" class="px-3 py-2 bg-indigo-50 text-indigo-600 hover:bg-indigo-100 text-xs font-semibold rounded-lg transition">
                  Refresh Data
                </button>
              </div>
            </div>

            <!-- Tabel dengan Touch Scroll di Mobile -->
            <div class="p-4 sm:p-6 overflow-x-auto">
              <table class="w-full text-left border-collapse min-w-[600px]">
                <thead>
                  <tr class="border-b border-slate-200 bg-slate-50 text-slate-500 uppercase text-xs tracking-wider">
                    <th class="py-3 px-4">ID</th>
                    <th class="py-3 px-4">Nama</th>
                    <th class="py-3 px-4">Email</th>
                    <th class="py-3 px-4">Peran (Role)</th>
                    <th class="py-3 px-4">Terdaftar</th>
                    <th v-if="isSuperadmin" class="py-3 px-4 text-center">Aksi (Ubah Role)</th>
                  </tr>
                </thead>
                <tbody class="divide-y divide-slate-100 text-sm">
                  <tr v-for="u in allUsers" :key="u.id" class="hover:bg-slate-50">
                    <td class="py-3 px-4 font-bold text-slate-700">#{{ u.id }}</td>
                    <td class="py-3 px-4 font-semibold text-slate-800">{{ u.name }}</td>
                    <td class="py-3 px-4 text-slate-600">{{ u.email }}</td>
                    <td class="py-3 px-4">
                      <span :class="roleBadgeClass(u.role)">{{ (u.role || 'admin').toUpperCase() }}</span>
                    </td>
                    <td class="py-3 px-4 text-xs text-slate-500">{{ new Date(u.created_at).toLocaleDateString('id-ID') }}</td>
                    
                    <!-- Tombol Ubah Role khusus Superadmin -->
                    <td v-if="isSuperadmin" class="py-3 px-4 text-center">
                      <select
                        :value="u.role || 'admin'"
                        @change="handleChangeRole(u.id, $event.target.value)"
                        class="px-2.5 py-1.5 text-xs border border-slate-300 rounded-md outline-none bg-white font-medium"
                      >
                        <option value="admin">ADMIN</option>
                        <option value="owner">OWNER</option>
                        <option value="superadmin">SUPERADMIN</option>
                      </select>
                    </td>
                  </tr>
                </tbody>
              </table>
            </div>
          </div>
        </template>
      </main>
    </div>

    <!-- Modal Form Tambah Pengguna Baru (Khusus Super Admin) -->
    <div v-if="isAddUserModalOpen" class="fixed inset-0 z-50 flex items-center justify-center p-4 bg-slate-900/60 backdrop-blur-sm">
      <div class="w-full max-w-md bg-white rounded-xl shadow-2xl border border-slate-200 overflow-hidden max-h-[90vh] flex flex-col animate-in fade-in zoom-in duration-200">
        <div class="px-6 py-4 border-b border-slate-200 bg-slate-50 flex items-center justify-between shrink-0">
          <div class="flex items-center space-x-2">
            <div class="p-1.5 bg-indigo-100 text-indigo-600 rounded-lg">
              <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18 9v3m0 0v3m0-3h3m-3 0h-3m-2-5a4 4 0 11-8 0 4 4 0 018 0zM3 20a6 6 0 0112 0v1H3v-1z"/>
              </svg>
            </div>
            <h3 class="text-base sm:text-lg font-bold text-slate-900">Tambah Pengguna Baru</h3>
          </div>
          <button @click="isAddUserModalOpen = false" class="text-slate-400 hover:text-slate-600 p-1 rounded-lg hover:bg-slate-100 transition">
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"/>
            </svg>
          </button>
        </div>

        <form @submit.prevent="handleCreateUser" class="p-6 space-y-4 overflow-y-auto flex-1">
          <div>
            <label class="block text-xs font-semibold text-slate-700 uppercase tracking-wider mb-1">Nama Lengkap</label>
            <input
              v-model="addUserForm.name"
              type="text"
              required
              placeholder="Nama Pengguna"
              class="w-full px-4 py-2.5 border border-slate-300 rounded-lg focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 outline-none text-sm"
            />
          </div>

          <div>
            <label class="block text-xs font-semibold text-slate-700 uppercase tracking-wider mb-1">Alamat Email</label>
            <input
              v-model="addUserForm.email"
              type="email"
              required
              placeholder="nama@email.com"
              class="w-full px-4 py-2.5 border border-slate-300 rounded-lg focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 outline-none text-sm"
            />
          </div>

          <div>
            <label class="block text-xs font-semibold text-slate-700 uppercase tracking-wider mb-1">Password</label>
            <input
              v-model="addUserForm.password"
              type="password"
              required
              placeholder="Minimal 6 karakter"
              class="w-full px-4 py-2.5 border border-slate-300 rounded-lg focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 outline-none text-sm"
            />
          </div>

          <div>
            <label class="block text-xs font-semibold text-slate-700 uppercase tracking-wider mb-1">Peran (Role)</label>
            <select
              v-model="addUserForm.role"
              class="w-full px-4 py-2.5 border border-slate-300 rounded-lg focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 outline-none text-sm bg-white font-medium"
            >
              <option value="admin">ADMIN (Standar)</option>
              <option value="owner">OWNER (Pemilik Laporan)</option>
              <option value="superadmin">SUPERADMIN (Hak Akses Penuh)</option>
            </select>
          </div>

          <div class="flex items-center justify-end space-x-3 pt-2">
            <button
              type="button"
              @click="isAddUserModalOpen = false"
              class="px-4 py-2.5 border border-slate-300 text-slate-700 font-semibold rounded-lg hover:bg-slate-50 transition text-sm"
            >
              Batal
            </button>
            <button
              type="submit"
              :disabled="isAddingUser"
              class="px-4 py-2.5 bg-indigo-600 hover:bg-indigo-700 disabled:opacity-50 text-white font-semibold rounded-lg shadow transition text-sm"
            >
              {{ isAddingUser ? 'Menyimpan...' : 'Simpan User Baru' }}
            </button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive, computed, onMounted, onUnmounted, watch } from 'vue'
import { useRouter } from 'vue-router'
import api from '../services/api'

const router = useRouter()
const activeTab = ref('dashboard')
const isDropdownOpen = ref(false)
const isMobileSidebarOpen = ref(false)
const isUpdating = ref(false)
const isChangingPassword = ref(false)
const isAddUserModalOpen = ref(false)
const isAddingUser = ref(false)
const allUsers = ref([])

const addUserForm = reactive({
  name: '',
  email: '',
  password: '',
  role: 'admin'
})

const user = reactive({
  id: '',
  name: '',
  email: '',
  role: 'admin',
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

const isSuperadmin = computed(() => {
  return user.role === 'superadmin'
})

const canManageUsers = computed(() => {
  return user.role === 'superadmin' || user.role === 'owner'
})

const roleBadgeClass = (role) => {
  switch (role) {
    case 'superadmin':
      return 'px-2 py-0.5 rounded text-[10px] font-extrabold bg-rose-100 text-rose-800 border border-rose-200'
    case 'owner':
      return 'px-2 py-0.5 rounded text-[10px] font-extrabold bg-purple-100 text-purple-800 border border-purple-200'
    default:
      return 'px-2 py-0.5 rounded text-[10px] font-extrabold bg-blue-100 text-blue-800 border border-blue-200'
  }
}

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
    user.role = data.role || 'admin'
    user.created_at = data.created_at

    editForm.name = data.name
    editForm.email = data.email
  } catch (err) {
    localStorage.removeItem('token')
    router.push('/login')
  }
}

const loadAllUsers = async () => {
  if (!canManageUsers.value) return
  try {
    const res = await api.get('/admin/users')
    allUsers.value = res.data.data
  } catch (err) {
    showAlert(err.response?.data?.message || 'Akses ditolak untuk mengambil data pengguna', false)
  }
}

watch(activeTab, (newTab) => {
  if (newTab === 'users') {
    loadAllUsers()
  }
})

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

const handleChangeRole = async (targetUserID, newRole) => {
  try {
    await api.put('/superadmin/change-role', {
      user_id: targetUserID,
      role: newRole
    })

    showAlert(`Berhasil mengubah peran pengguna #${targetUserID} menjadi ${newRole.toUpperCase()}`, true)
    loadAllUsers()
  } catch (err) {
    showAlert(err.response?.data?.message || 'Gagal mengubah peran pengguna', false)
  }
}

const handleCreateUser = async () => {
  isAddingUser.value = true
  try {
    const res = await api.post('/superadmin/users', {
      name: addUserForm.name,
      email: addUserForm.email,
      password: addUserForm.password,
      role: addUserForm.role
    })

    showAlert(`Pengguna baru '${res.data.data.name}' (${res.data.data.role.toUpperCase()}) berhasil dibuat!`, true)

    // Reset form & tutup modal
    addUserForm.name = ''
    addUserForm.email = ''
    addUserForm.password = ''
    addUserForm.role = 'admin'
    isAddUserModalOpen.value = false

    // Refresh daftar user
    loadAllUsers()
  } catch (err) {
    showAlert(err.response?.data?.message || 'Gagal menambahkan pengguna baru', false)
  } finally {
    isAddingUser.value = false
  }
}

const handleLogout = async () => {
  try {
    await api.post('/logout')
  } catch (err) {
    // Ignore error if server unreachable
  } finally {
    localStorage.removeItem('token')
    localStorage.removeItem('access_token')
    localStorage.removeItem('refresh_token')
    router.push('/login')
  }
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

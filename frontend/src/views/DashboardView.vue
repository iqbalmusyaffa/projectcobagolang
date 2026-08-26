<template>
  <div class="flex min-h-screen bg-slate-100 dark:bg-slate-950 font-sans text-slate-800 dark:text-slate-100 antialiased transition-colors duration-200">
    <!-- Desktop Sidebar TailAdmin -->
    <!-- Desktop Sidebar TailAdmin -->
    <aside class="w-64 bg-white dark:bg-slate-900 border-r border-slate-200 dark:border-slate-800 text-slate-700 dark:text-slate-300 flex flex-col justify-between hidden md:flex min-h-screen sticky top-0 h-screen transition-colors duration-200">
      <div>
        <div class="h-16 flex items-center px-6 border-b border-slate-200 dark:border-slate-800">
          <div class="inline-flex items-center justify-center w-8 h-8 bg-indigo-600 text-white rounded-lg font-bold mr-3 shadow">
            TA
          </div>
          <span class="text-lg font-bold text-slate-900 dark:text-white tracking-wide">TailAdmin Vue</span>
        </div>

        <nav class="p-4 space-y-1">
          <div class="px-3 text-xs font-semibold text-slate-400 dark:text-slate-500 uppercase tracking-wider mb-2">Menu Utama</div>
          
          <a href="#" @click.prevent="activeTab = 'dashboard'" :class="[
            'flex items-center space-x-3 px-3 py-2.5 rounded-lg font-medium text-sm transition',
            activeTab === 'dashboard' ? 'bg-indigo-50 dark:bg-slate-800 text-indigo-600 dark:text-white shadow-sm' : 'text-slate-600 dark:text-slate-400 hover:bg-slate-100 dark:hover:bg-slate-800 hover:text-slate-900 dark:hover:text-white'
          ]">
            <svg class="w-5 h-5 text-indigo-500 dark:text-indigo-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 12l2-2m0 0l7-7 7 7M5 10v10a1 1 0 001 1h3m10-11l2 2m-2-2v10a1 1 0 01-1 1h-3m-6 0a1 1 0 001-1v-4a1 1 0 011-1h2a1 1 0 011 1v4a1 1 0 001 1m-6 0h6"/>
            </svg>
            <span>Dashboard</span>
          </a>

          <!-- Menu Khusus Superadmin & Owner -->
          <a v-if="canManageUsers" href="#" @click.prevent="activeTab = 'users'" :class="[
            'flex items-center space-x-3 px-3 py-2.5 rounded-lg font-medium text-sm transition',
            activeTab === 'users' ? 'bg-purple-50 dark:bg-slate-800 text-purple-600 dark:text-white shadow-sm' : 'text-slate-600 dark:text-slate-400 hover:bg-slate-100 dark:hover:bg-slate-800 hover:text-slate-900 dark:hover:text-white'
          ]">
            <svg class="w-5 h-5 text-purple-500 dark:text-purple-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4.354a4 4 0 110 5.292M15 21H3v-1a6 6 0 0112 0v1zm0 0h6v-1a6 6 0 00-9-5.197M13 7a4 4 0 11-8 0 4 4 0 018 0z"/>
            </svg>
            <span>Manajemen User</span>
          </a>

          <!-- Halaman Profil Saya -->
          <a href="#" @click.prevent="activeTab = 'profile'" :class="[
            'flex items-center space-x-3 px-3 py-2.5 rounded-lg font-medium text-sm transition',
            activeTab === 'profile' ? 'bg-emerald-50 dark:bg-slate-800 text-emerald-600 dark:text-white shadow-sm' : 'text-slate-600 dark:text-slate-400 hover:bg-slate-100 dark:hover:bg-slate-800 hover:text-slate-900 dark:hover:text-white'
          ]">
            <svg class="w-5 h-5 text-emerald-500 dark:text-emerald-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z"/>
            </svg>
            <span>Profil Saya</span>
          </a>
        </nav>
      </div>

      <div class="p-4 border-t border-slate-200 dark:border-slate-800">
        <button @click="handleLogout" class="w-full flex items-center justify-center space-x-2 py-2 px-4 bg-rose-50 dark:bg-rose-600/10 text-rose-600 dark:text-rose-400 hover:bg-rose-600 hover:text-white rounded-lg transition text-sm font-semibold">
          <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 16l4-4m0 0l-4-4m4 4H7m6 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h4a3 3 0 013 3v1"/>
          </svg>
          <span>Keluar (Logout)</span>
        </button>
      </div>
    </aside>

    <!-- Mobile Drawer Off-Canvas Sidebar (<768px) -->
    <div v-if="isMobileSidebarOpen" class="fixed inset-0 z-50 flex md:hidden">
      <!-- Backdrop Overlay -->
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

          <a href="#" @click.prevent="activeTab = 'profile'; isMobileSidebarOpen = false" :class="[
            'flex items-center space-x-3 px-3 py-3 rounded-lg font-medium text-sm transition',
            activeTab === 'profile' ? 'bg-slate-800 text-white shadow-sm' : 'text-slate-400 hover:bg-slate-800 hover:text-white'
          ]">
            <svg class="w-5 h-5 text-emerald-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z"/>
            </svg>
            <span>Profil Saya</span>
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
      <header class="h-16 bg-white dark:bg-slate-900 border-b border-slate-200 dark:border-slate-800 flex items-center justify-between px-4 sm:px-6 sticky top-0 z-30 shadow-sm transition-colors duration-200">
        <div class="flex items-center space-x-3">
          <!-- Hamburger Button khusus Mobile/Tablet (<768px) -->
          <button @click="isMobileSidebarOpen = true" class="md:hidden p-2 rounded-lg text-slate-600 dark:text-slate-300 hover:bg-slate-100 dark:hover:bg-slate-800 focus:outline-none transition">
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h16"/>
            </svg>
          </button>

          <h1 class="text-base sm:text-xl font-bold text-slate-800 dark:text-white truncate">
            {{ activeTab === 'users' ? 'Manajemen Pengguna (RBAC)' : activeTab === 'profile' ? 'Pengaturan Profil Saya' : 'Overview Dashboard Vue 3' }}
          </h1>
        </div>

        <!-- Actions Header Right: Sinkron Data, Notifikasi Lonceng & Profil Sesi -->
        <div class="flex items-center space-x-2 sm:space-x-3">
          <!-- Tombol Switcher Theme (Sun/Moon) -->
          <button
            type="button"
            @click.stop="toggleTheme"
            class="p-2 text-slate-500 dark:text-slate-300 hover:text-amber-500 dark:hover:text-amber-400 hover:bg-slate-100 dark:hover:bg-slate-800 rounded-lg transition outline-none cursor-pointer"
            :title="isDarkMode ? 'Beralih ke Mode Terang' : 'Beralih ke Mode Gelap'"
          >
            <svg v-if="isDarkMode" class="w-5 h-5 text-amber-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 3v1m0 16v1m9-9h-1M4 12H3m15.364 6.364l-.707-.707M6.343 6.343l-.707-.707m12.728 0l-.707.707M6.343 17.657l-.707.707M16 12a4 4 0 11-8 0 4 4 0 018 0z" />
            </svg>
            <svg v-else class="w-5 h-5 text-slate-600 dark:text-slate-300" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20.354 15.354A9 9 0 018.646 3.646 9.003 9.003 0 0012 21a9.003 9.003 0 008.354-5.646z" />
            </svg>
          </button>

          <!-- Tombol Sinkron Data -->
          <button 
            @click="handleSyncData" 
            :disabled="isSyncing"
            class="px-3 py-1.5 bg-indigo-50 dark:bg-indigo-950/60 hover:bg-indigo-100 dark:hover:bg-indigo-900/60 text-indigo-700 dark:text-indigo-300 disabled:opacity-50 text-xs font-semibold rounded-lg transition flex items-center space-x-1.5"
            title="Sinkronkan data terbaru dari database PostgreSQL"
          >
            <svg :class="['w-4 h-4 text-indigo-600 dark:text-indigo-400', isSyncing ? 'animate-spin' : '']" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15"/>
            </svg>
            <span class="hidden sm:inline">{{ isSyncing ? 'Menyinkronkan...' : 'Sinkron Data' }}</span>
          </button>

          <!-- Notification Bell Menu Popover -->
          <div class="relative">
            <button 
              @click.stop="isNotificationOpen = !isNotificationOpen; isDropdownOpen = false" 
              class="relative p-2 text-slate-500 hover:text-indigo-600 hover:bg-slate-100 rounded-lg transition outline-none"
              title="Notifikasi Sistem"
            >
              <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 17h5l-1.405-1.405A2.032 2.032 0 0118 14.158V11a6.002 6.002 0 00-4-5.659V5a2 2 0 10-4 0v.341C7.67 6.165 6 8.388 6 11v3.159c0 .538-.214 1.055-.595 1.436L4 17h5m6 0v1a3 3 0 11-6 0v-1m6 0H9"/>
              </svg>
              <!-- Unread Badge Indicator -->
              <span v-if="unreadCount > 0" class="absolute top-1 right-1 flex h-4 w-4">
                <span class="animate-ping absolute inline-flex h-full w-full rounded-full bg-rose-400 opacity-75"></span>
                <span class="relative inline-flex rounded-full h-4 w-4 bg-rose-600 text-[10px] font-bold text-white items-center justify-center">{{ unreadCount }}</span>
              </span>
            </button>

            <!-- Notification Dropdown Popover Panel -->
            <div v-if="isNotificationOpen" class="absolute -right-16 sm:right-0 mt-2 w-80 sm:w-96 bg-white dark:bg-slate-900 rounded-2xl shadow-2xl border border-slate-200 dark:border-slate-800 py-3 z-50 animate-in fade-in zoom-in duration-150">
              <div class="px-4 py-2 border-b border-slate-100 dark:border-slate-800 flex items-center justify-between">
                <div class="flex items-center space-x-2">
                  <h3 class="text-sm font-bold text-slate-900 dark:text-white">Notifikasi Sistem</h3>
                  <span v-if="unreadCount > 0" class="px-2 py-0.5 text-[10px] font-extrabold bg-rose-100 dark:bg-rose-950/60 text-rose-700 dark:text-rose-300 rounded-full">{{ unreadCount }} Baru</span>
                </div>
                <div class="flex items-center space-x-2">
                  <button @click="markAllAsRead" class="text-[11px] text-indigo-600 dark:text-indigo-400 hover:text-indigo-800 font-semibold transition">Tandai Dibaca</button>
                  <span class="text-slate-300 dark:text-slate-700">|</span>
                  <button @click="clearNotifications" class="text-[11px] text-slate-400 hover:text-slate-600 dark:hover:text-slate-200 font-medium transition">Hapus</button>
                </div>
              </div>

              <!-- List Notifikasi -->
              <div class="max-h-72 overflow-y-auto divide-y divide-slate-100 dark:divide-slate-800">
                <div 
                  v-for="item in notifications" 
                  :key="item.id" 
                  @click="markAsRead(item.id)"
                  :class="[
                    'p-3.5 hover:bg-slate-50 dark:hover:bg-slate-800/60 transition cursor-pointer flex items-start space-x-3',
                    !item.read ? 'bg-indigo-50/40 dark:bg-indigo-950/40' : ''
                  ]"
                >
                  <div :class="[
                    'p-2 rounded-xl shrink-0 mt-0.5',
                    item.type === 'security' ? 'bg-emerald-100 dark:bg-emerald-950/60 text-emerald-600 dark:text-emerald-400' :
                    item.type === 'profile' ? 'bg-indigo-100 dark:bg-indigo-950/60 text-indigo-600 dark:text-indigo-400' :
                    'bg-purple-100 dark:bg-purple-950/60 text-purple-600 dark:text-purple-400'
                  ]">
                    <svg v-if="item.type === 'security'" class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m5.618-4.016A11.955 11.955 0 0112 2.944a11.955 11.955 0 01-8.618 3.04A12.02 12.02 0 003 9c0 5.591 3.824 10.29 9 11.622 5.176-1.332 9-6.03 9-11.622 0-1.042-.133-2.052-.382-3.016z"/>
                    </svg>
                    <svg v-else-if="item.type === 'profile'" class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z"/>
                    </svg>
                    <svg v-else class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"/>
                    </svg>
                  </div>
                  <div class="flex-1 min-w-0">
                    <div class="flex items-center justify-between">
                      <h4 :class="['text-xs font-bold truncate', !item.read ? 'text-indigo-950 dark:text-indigo-200' : 'text-slate-700 dark:text-slate-300']">{{ item.title }}</h4>
                      <span class="text-[10px] text-slate-400 shrink-0 ml-2">{{ formatRelativeTime(item.created_at) }}</span>
                    </div>
                    <p class="text-xs text-slate-500 dark:text-slate-400 mt-0.5 line-clamp-2 leading-relaxed">{{ item.message }}</p>
                  </div>
                </div>

                <div v-if="notifications.length === 0" class="p-6 text-center text-xs text-slate-400">
                  Tidak ada notifikasi saat ini
                </div>
              </div>
            </div>
          </div>

          <!-- User Info Header Dropdown -->
          <div class="relative">
            <div @click.stop="isDropdownOpen = !isDropdownOpen; isNotificationOpen = false" class="flex items-center space-x-2 sm:space-x-3 cursor-pointer select-none p-1.5 rounded-lg hover:bg-slate-100 dark:hover:bg-slate-800 transition">
              <div class="text-right hidden sm:block">
                <div class="flex items-center space-x-2 justify-end">
                  <span class="text-sm font-semibold text-slate-800 dark:text-white">{{ user.name || 'Loading...' }}</span>
                  <!-- Role Badge Header -->
                  <span :class="roleBadgeClass(user.role)">{{ (user.role || 'admin').toUpperCase() }}</span>
                </div>
                <div class="text-xs text-slate-500 dark:text-slate-400">{{ user.email || 'Loading...' }}</div>
              </div>
              
              <!-- Header Avatar Image / Initial -->
              <div class="w-9 h-9 sm:w-10 sm:h-10 rounded-full bg-indigo-600 text-white font-bold flex items-center justify-center shadow text-sm sm:text-base overflow-hidden border border-indigo-200 dark:border-indigo-900 shrink-0">
                <img v-if="avatarUrl" :src="avatarUrl" alt="Avatar" class="w-full h-full object-cover" />
                <span v-else>{{ avatarInitial }}</span>
              </div>

              <svg class="w-4 h-4 text-slate-400 hidden sm:block" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7"/>
              </svg>
            </div>

            <!-- Dropdown Menu Vue 3 -->
            <div v-if="isDropdownOpen" class="absolute right-0 mt-2 w-60 bg-white dark:bg-slate-900 rounded-xl shadow-xl border border-slate-200 dark:border-slate-800 py-2 z-50 animate-in fade-in zoom-in duration-150">
              <div class="px-4 py-2 border-b border-slate-100 dark:border-slate-800">
                <p class="text-sm font-bold text-slate-800 dark:text-white truncate">{{ user.name }}</p>
                <p class="text-xs text-slate-500 dark:text-slate-400 truncate mb-1">{{ user.email }}</p>
                <span :class="roleBadgeClass(user.role)">Peran: {{ (user.role || 'admin').toUpperCase() }}</span>
              </div>

              <a href="#" @click.prevent="activeTab = 'profile'; isDropdownOpen = false" class="flex items-center px-4 py-2 text-sm text-slate-700 dark:text-slate-300 hover:bg-slate-50 dark:hover:bg-slate-800 hover:text-indigo-600 dark:hover:text-indigo-400 transition">
                <svg class="w-4 h-4 mr-2.5 text-slate-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z"/>
                </svg>
                Profil Saya
              </a>

              <div class="border-t border-slate-100 dark:border-slate-800 my-1"></div>

              <button @click="handleLogout" class="w-full text-left flex items-center px-4 py-2 text-sm text-rose-600 dark:text-rose-400 hover:bg-rose-50 dark:hover:bg-rose-950/40 font-medium transition">
                <svg class="w-4 h-4 mr-2.5 text-rose-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 16l4-4m0 0l-4-4m4 4H7m6 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h4a3 3 0 013 3v1"/>
                </svg>
                Keluar (Logout)
              </button>
            </div>
          </div>
        </div>
      </header>

      <!-- Main Body Responsif -->
      <main class="flex-1 overflow-y-auto p-4 sm:p-6 lg:p-8 space-y-6 bg-slate-100 dark:bg-slate-950 transition-colors duration-200">
        <!-- Reaktif Alert Box -->
        <div v-if="alert.message" :class="[
          'p-4 rounded-xl text-sm font-medium shadow-sm transition',
          alert.isSuccess ? 'bg-emerald-100 dark:bg-emerald-950/50 text-emerald-800 dark:text-emerald-300 border border-emerald-300 dark:border-emerald-800' : 'bg-rose-100 dark:bg-rose-950/50 text-rose-800 dark:text-rose-300 border border-rose-300 dark:border-rose-800'
        ]">
          {{ alert.message }}
        </div>

        <!-- TAB 1: MAIN DASHBOARD -->
        <template v-if="activeTab === 'dashboard'">
          <!-- Stat Cards Grid 5 Kolom (Realtime dari PostgreSQL) -->
          <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-5 gap-4 sm:gap-6">
            <!-- Card 1: Total Users -->
            <div class="bg-white dark:bg-slate-900 p-4 sm:p-5 rounded-xl border border-slate-200 dark:border-slate-800 shadow-sm flex items-center space-x-3.5">
              <div class="p-3 bg-indigo-50 dark:bg-indigo-950/60 text-indigo-600 dark:text-indigo-400 rounded-xl shrink-0">
                <svg class="w-6 h-6 sm:w-7 sm:h-7" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z"/>
                </svg>
              </div>
              <div class="min-w-0">
                <p class="text-[11px] font-semibold text-slate-500 dark:text-slate-400 uppercase tracking-wider truncate">Total Akun</p>
                <h3 class="text-xl sm:text-2xl font-black text-slate-900 dark:text-white mt-0.5">{{ stats.total_users || 0 }}</h3>
              </div>
            </div>

            <!-- Card 2: Superadmin Count -->
            <div class="bg-white dark:bg-slate-900 p-4 sm:p-5 rounded-xl border border-slate-200 dark:border-slate-800 shadow-sm flex items-center space-x-3.5">
              <div class="p-3 bg-rose-50 dark:bg-rose-950/60 text-rose-600 dark:text-rose-400 rounded-xl shrink-0">
                <svg class="w-6 h-6 sm:w-7 sm:h-7" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m5.618-4.016A11.955 11.955 0 0112 2.944a11.955 11.955 0 01-8.618 3.04A12.02 12.02 0 003 9c0 5.591 3.824 10.29 9 11.622 5.176-1.332 9-6.03 9-11.622 0-1.042-.133-2.052-.382-3.016z"/>
                </svg>
              </div>
              <div class="min-w-0">
                <p class="text-[11px] font-semibold text-slate-500 dark:text-slate-400 uppercase tracking-wider truncate">Super Admin</p>
                <h3 class="text-xl sm:text-2xl font-black text-rose-600 dark:text-rose-400 mt-0.5">{{ stats.total_superadmin || 0 }}</h3>
              </div>
            </div>

            <!-- Card 3: Owner Count -->
            <div class="bg-white dark:bg-slate-900 p-4 sm:p-5 rounded-xl border border-slate-200 dark:border-slate-800 shadow-sm flex items-center space-x-3.5">
              <div class="p-3 bg-purple-50 dark:bg-purple-950/60 text-purple-600 dark:text-purple-400 rounded-xl shrink-0">
                <svg class="w-6 h-6 sm:w-7 sm:h-7" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 21V5a2 2 0 00-2-2H7a2 2 0 00-2 2v16m14 0h2m-2 0h-5m-9 0H3m2 0h5m0 0h4m-4 0a2 2 0 104 0m-4 0a2 2 0 114 0"/>
                </svg>
              </div>
              <div class="min-w-0">
                <p class="text-[11px] font-semibold text-slate-500 dark:text-slate-400 uppercase tracking-wider truncate">Owner Laporan</p>
                <h3 class="text-xl sm:text-2xl font-black text-purple-600 dark:text-purple-400 mt-0.5">{{ stats.total_owner || 0 }}</h3>
              </div>
            </div>

            <!-- Card 4: Admin Count -->
            <div class="bg-white dark:bg-slate-900 p-4 sm:p-5 rounded-xl border border-slate-200 dark:border-slate-800 shadow-sm flex items-center space-x-3.5">
              <div class="p-3 bg-blue-50 dark:bg-blue-950/60 text-blue-600 dark:text-blue-400 rounded-xl shrink-0">
                <svg class="w-6 h-6 sm:w-7 sm:h-7" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z"/>
                </svg>
              </div>
              <div class="min-w-0">
                <p class="text-[11px] font-semibold text-slate-500 dark:text-slate-400 uppercase tracking-wider truncate">Admin Standar</p>
                <h3 class="text-xl sm:text-2xl font-black text-blue-600 dark:text-blue-400 mt-0.5">{{ stats.total_admin || 0 }}</h3>
              </div>
            </div>

            <!-- Card 5: User Regular Count -->
            <div class="bg-white dark:bg-slate-900 p-4 sm:p-5 rounded-xl border border-slate-200 dark:border-slate-800 shadow-sm flex items-center space-x-3.5">
              <div class="p-3 bg-emerald-50 dark:bg-emerald-950/60 text-emerald-600 dark:text-emerald-400 rounded-xl shrink-0">
                <svg class="w-6 h-6 sm:w-7 sm:h-7" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4.354a4 4 0 110 5.292M15 21H3v-1a6 6 0 0112 0v1zm0 0h6v-1a6 6 0 00-9-5.197M13 7a4 4 0 11-8 0 4 4 0 018 0z"/>
                </svg>
              </div>
              <div class="min-w-0">
                <p class="text-[11px] font-semibold text-slate-500 dark:text-slate-400 uppercase tracking-wider truncate">User Biasa</p>
                <h3 class="text-xl sm:text-2xl font-black text-emerald-600 dark:text-emerald-400 mt-0.5">{{ stats.total_user_role || 0 }}</h3>
              </div>
            </div>
          </div>

          <!-- Section Grafik Visualisasi Chart.js Analytics -->
          <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
            <!-- Chart 1: Doughnut Chart (Distribusi Hak Akses / Role) -->
            <div class="bg-white dark:bg-slate-900 rounded-xl border border-slate-200 dark:border-slate-800 shadow-sm p-5 sm:p-6 flex flex-col justify-between">
              <div class="flex items-center justify-between border-b border-slate-100 dark:border-slate-800 pb-3 mb-4">
                <div>
                  <h3 class="text-base font-bold text-slate-900 dark:text-white">Distribusi Peran Pengguna (RBAC)</h3>
                  <p class="text-xs text-slate-500 dark:text-slate-400 mt-0.5">Persentase hak akses pengguna di sistem</p>
                </div>
                <span class="px-2.5 py-1 bg-indigo-50 dark:bg-indigo-950/60 text-indigo-700 dark:text-indigo-300 text-[11px] font-bold rounded-lg">Chart.js</span>
              </div>

              <div class="relative h-64 sm:h-72 flex items-center justify-center">
                <canvas ref="roleDoughnutChartCanvas"></canvas>
              </div>
            </div>

            <!-- Chart 2: Bar Chart (Komparasi Jumlah User Berdasarkan Role) -->
            <div class="bg-white dark:bg-slate-900 rounded-xl border border-slate-200 dark:border-slate-800 shadow-sm p-5 sm:p-6 flex flex-col justify-between">
              <div class="flex items-center justify-between border-b border-slate-100 dark:border-slate-800 pb-3 mb-4">
                <div>
                  <h3 class="text-base font-bold text-slate-900 dark:text-white">Statistik Pengguna Berdasarkan Role</h3>
                  <p class="text-xs text-slate-500 dark:text-slate-400 mt-0.5">Jumlah akun aktif per kategori peran</p>
                </div>
                <span class="px-2.5 py-1 bg-emerald-50 dark:bg-emerald-950/60 text-emerald-700 dark:text-emerald-300 text-[11px] font-bold rounded-lg">Realtime</span>
              </div>

              <div class="relative h-64 sm:h-72 flex items-center justify-center">
                <canvas ref="roleBarChartCanvas"></canvas>
              </div>
            </div>
          </div>

          <!-- User Quick Overview Card -->
          <div class="bg-white dark:bg-slate-900 rounded-xl border border-slate-200 dark:border-slate-800 shadow-sm overflow-hidden">
            <div class="px-5 sm:px-6 py-4 border-b border-slate-200 dark:border-slate-800 bg-slate-50 dark:bg-slate-800/60 flex items-center justify-between">
              <h2 class="text-base sm:text-lg font-bold text-slate-900 dark:text-white">Ringkasan Sesi Pengguna</h2>
              <button @click="activeTab = 'profile'" class="text-xs text-indigo-600 dark:text-indigo-400 hover:text-indigo-800 font-semibold">Buka Profil Saya &rarr;</button>
            </div>
            <div class="p-5 sm:p-6">
              <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-4 sm:gap-6">
                <div>
                  <label class="block text-xs font-semibold text-slate-400 uppercase tracking-wider mb-1">User ID</label>
                  <p class="text-sm font-bold text-slate-800 dark:text-slate-200 bg-slate-100 dark:bg-slate-800 px-4 py-2 rounded-lg border border-slate-200 dark:border-slate-700 truncate">{{ user.id || '-' }}</p>
                </div>
                <div>
                  <label class="block text-xs font-semibold text-slate-400 uppercase tracking-wider mb-1">Nama Lengkap</label>
                  <p class="text-sm font-bold text-slate-800 dark:text-slate-200 bg-slate-100 dark:bg-slate-800 px-4 py-2 rounded-lg border border-slate-200 dark:border-slate-700 truncate">{{ user.name || '-' }}</p>
                </div>
                <div>
                  <label class="block text-xs font-semibold text-slate-400 uppercase tracking-wider mb-1">Email</label>
                  <p class="text-sm font-bold text-slate-800 dark:text-slate-200 bg-slate-100 dark:bg-slate-800 px-4 py-2 rounded-lg border border-slate-200 dark:border-slate-700 truncate">{{ user.email || '-' }}</p>
                </div>
                <div>
                  <label class="block text-xs font-semibold text-slate-400 uppercase tracking-wider mb-1">Nomor Telepon / HP</label>
                  <p class="text-sm font-bold text-slate-800 dark:text-slate-200 bg-slate-100 dark:bg-slate-800 px-4 py-2 rounded-lg border border-slate-200 dark:border-slate-700 truncate">{{ user.phone || '-' }}</p>
                </div>
                <div>
                  <label class="block text-xs font-semibold text-slate-400 uppercase tracking-wider mb-1">Jenis Kelamin</label>
                  <p class="text-sm font-bold text-slate-800 dark:text-slate-200 bg-slate-100 dark:bg-slate-800 px-4 py-2 rounded-lg border border-slate-200 dark:border-slate-700 truncate">{{ user.gender || '-' }}</p>
                </div>
                <div>
                  <label class="block text-xs font-semibold text-slate-400 uppercase tracking-wider mb-1">Tanggal Lahir</label>
                  <p class="text-sm font-bold text-slate-800 dark:text-slate-200 bg-slate-100 dark:bg-slate-800 px-4 py-2 rounded-lg border border-slate-200 dark:border-slate-700 truncate">{{ user.birth_date || '-' }}</p>
                </div>
                <div>
                  <label class="block text-xs font-semibold text-slate-400 uppercase tracking-wider mb-1">Hak Akses / Role</label>
                  <p class="text-sm font-bold text-indigo-700 dark:text-indigo-300 bg-indigo-50 dark:bg-indigo-950/60 px-4 py-2 rounded-lg border border-indigo-200 dark:border-indigo-800 uppercase truncate">{{ user.role || 'ADMIN' }}</p>
                </div>
                <div>
                  <label class="block text-xs font-semibold text-slate-400 uppercase tracking-wider mb-1">Terakhir Diperbarui</label>
                  <p class="text-sm font-bold text-slate-800 dark:text-slate-200 bg-slate-100 dark:bg-slate-800 px-4 py-2 rounded-lg border border-slate-200 dark:border-slate-700 truncate">{{ formattedUpdatedAt }}</p>
                </div>
              </div>
            </div>
          </div>
        </template>

        <!-- TAB 2: HALAMAN PROFIL SAYA (DEDICATED PAGE) -->
        <template v-if="activeTab === 'profile'">
          <div class="space-y-6">
            <!-- Header Profile Card + Upload Avatar -->
            <div class="bg-white dark:bg-slate-900 rounded-xl border border-slate-200 dark:border-slate-800 shadow-sm p-6 flex flex-col md:flex-row items-center justify-between gap-6">
              <div class="flex flex-col sm:flex-row items-center space-y-4 sm:space-y-0 sm:space-x-6 text-center sm:text-left">
                <!-- Avatar Preview (Click to Zoom & Download) -->
                <div @click="openAvatarModal(user.name, avatarUrl)" class="relative group cursor-pointer shrink-0" title="Klik untuk perbesar & unduh foto">
                  <img v-if="avatarUrl" :src="avatarUrl" alt="Avatar" class="w-24 h-24 rounded-full object-cover border-4 border-indigo-100 dark:border-indigo-900/60 shadow-md group-hover:scale-105 transition duration-200" />
                  <div v-else class="w-24 h-24 rounded-full bg-indigo-600 text-white font-extrabold text-3xl flex items-center justify-center border-4 border-indigo-100 dark:border-indigo-900/60 shadow-md group-hover:scale-105 transition duration-200">
                    {{ avatarInitial }}
                  </div>
                  <div class="absolute inset-0 bg-slate-900/40 rounded-full opacity-0 group-hover:opacity-100 flex flex-col items-center justify-center transition duration-200 text-white text-[10px] font-bold">
                    <svg class="w-5 h-5 mb-0.5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0zM10 7v3m0 0v3m0-3h3m-3 0H7"/>
                    </svg>
                    Lihat Foto
                  </div>
                </div>

                <div>
                  <h2 class="text-xl font-bold text-slate-900 dark:text-white">{{ user.name || '-' }}</h2>
                  <p class="text-sm text-slate-500 dark:text-slate-400">{{ user.email || '-' }}</p>
                  <div class="mt-2 flex items-center space-x-2">
                    <span :class="roleBadgeClass(user.role)">Peran: {{ (user.role || 'admin').toUpperCase() }}</span>
                    <span v-if="user.updated_at" class="text-[11px] text-slate-400">Diuji: {{ new Date(user.updated_at).toLocaleDateString('id-ID') }}</span>
                  </div>
                </div>
              </div>

              <!-- Form Upload Avatar File -->
              <div class="w-full md:w-auto bg-slate-50 dark:bg-slate-800/60 p-4 rounded-xl border border-slate-200 dark:border-slate-700 text-center sm:text-left">
                <label class="block text-xs font-bold text-slate-700 dark:text-slate-300 uppercase tracking-wider mb-2">Unggah Foto Profil (Avatar)</label>
                <div class="flex flex-col sm:flex-row items-center gap-3">
                  <input 
                    type="file" 
                    @change="handleFileSelected" 
                    accept="image/png, image/jpeg, image/jpg, image/webp" 
                    class="text-xs text-slate-500 dark:text-slate-400 file:mr-3 file:py-2 file:px-3 file:rounded-lg file:border-0 file:text-xs file:font-semibold file:bg-indigo-50 dark:file:bg-indigo-950/60 file:text-indigo-700 dark:file:text-indigo-300 hover:file:bg-indigo-100 cursor-pointer"
                  />
                  <button 
                    @click="handleUploadAvatar" 
                    :disabled="!selectedFile || isUploadingAvatar"
                    class="w-full sm:w-auto px-4 py-2 bg-indigo-600 hover:bg-indigo-700 disabled:opacity-40 text-white text-xs font-semibold rounded-lg shadow transition shrink-0"
                  >
                    {{ isUploadingAvatar ? 'Mengunggah...' : 'Unggah Foto' }}
                  </button>
                </div>
                <p class="text-[11px] text-slate-400 mt-2">Format: JPG, PNG, WEBP (Maksimal 2MB)</p>
              </div>
            </div>

            <!-- Form Action: Edit Profil & Ganti Password -->
            <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
              <!-- Form Edit Profil Card -->
              <div class="bg-white dark:bg-slate-900 rounded-xl border border-slate-200 dark:border-slate-800 shadow-sm overflow-hidden flex flex-col justify-between">
                <div>
                  <div class="px-5 sm:px-6 py-4 border-b border-slate-200 dark:border-slate-800 bg-slate-50 dark:bg-slate-800/60">
                    <h3 class="text-base sm:text-lg font-bold text-slate-900 dark:text-white">Edit Profil</h3>
                    <p class="text-xs text-slate-500 dark:text-slate-400 mt-0.5">Perbarui biodata dan informasi kontak akun Anda</p>
                  </div>

                  <form @submit.prevent="handleUpdateProfile" class="p-5 sm:p-6 space-y-4">
                    <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
                      <div>
                        <label class="block text-xs font-semibold text-slate-700 dark:text-slate-300 uppercase tracking-wider mb-1">Nama Lengkap</label>
                        <input
                          v-model="editForm.name"
                          type="text"
                          required
                          class="w-full px-4 py-2.5 border border-slate-300 dark:border-slate-700 bg-white dark:bg-slate-800 text-slate-900 dark:text-white rounded-lg focus:ring-2 focus:ring-indigo-500 outline-none text-sm transition"
                        />
                      </div>

                      <div>
                        <label class="block text-xs font-semibold text-slate-700 dark:text-slate-300 uppercase tracking-wider mb-1">Alamat Email</label>
                        <input
                          v-model="editForm.email"
                          type="email"
                          required
                          class="w-full px-4 py-2.5 border border-slate-300 dark:border-slate-700 bg-white dark:bg-slate-800 text-slate-900 dark:text-white rounded-lg focus:ring-2 focus:ring-indigo-500 outline-none text-sm transition"
                        />
                      </div>
                    </div>

                    <div class="grid grid-cols-1 sm:grid-cols-3 gap-4">
                      <div>
                        <label class="block text-xs font-semibold text-slate-700 dark:text-slate-300 uppercase tracking-wider mb-1">Nomor Telepon / HP</label>
                        <input
                          v-model="editForm.phone"
                          type="text"
                          placeholder="081234567890"
                          class="w-full px-4 py-2.5 border border-slate-300 dark:border-slate-700 bg-white dark:bg-slate-800 text-slate-900 dark:text-white dark:placeholder-slate-500 rounded-lg focus:ring-2 focus:ring-indigo-500 outline-none text-sm transition"
                        />
                      </div>

                      <div>
                        <label class="block text-xs font-semibold text-slate-700 dark:text-slate-300 uppercase tracking-wider mb-1">Jenis Kelamin</label>
                        <select
                          v-model="editForm.gender"
                          class="w-full px-4 py-2.5 border border-slate-300 dark:border-slate-700 bg-white dark:bg-slate-800 text-slate-900 dark:text-white rounded-lg focus:ring-2 focus:ring-indigo-500 outline-none text-sm font-medium transition"
                        >
                          <option value="">-- Pilih --</option>
                          <option value="Laki-laki">Laki-laki</option>
                          <option value="Perempuan">Perempuan</option>
                        </select>
                      </div>

                      <div>
                        <label class="block text-xs font-semibold text-slate-700 dark:text-slate-300 uppercase tracking-wider mb-1">Tanggal Lahir</label>
                        <input
                          v-model="editForm.birth_date"
                          type="date"
                          class="w-full px-4 py-2.5 border border-slate-300 dark:border-slate-700 bg-white dark:bg-slate-800 text-slate-900 dark:text-white rounded-lg focus:ring-2 focus:ring-indigo-500 outline-none text-sm transition"
                        />
                      </div>
                    </div>

                    <div>
                      <label class="block text-xs font-semibold text-slate-700 dark:text-slate-300 uppercase tracking-wider mb-1">Alamat Lengkap</label>
                      <textarea
                        v-model="editForm.address"
                        rows="2"
                        placeholder="Alamat domisili Anda"
                        class="w-full px-4 py-2.5 border border-slate-300 dark:border-slate-700 bg-white dark:bg-slate-800 text-slate-900 dark:text-white dark:placeholder-slate-500 rounded-lg focus:ring-2 focus:ring-indigo-500 outline-none text-sm transition"
                      ></textarea>
                    </div>

                    <div>
                      <label class="block text-xs font-semibold text-slate-700 dark:text-slate-300 uppercase tracking-wider mb-1">Bio / Deskripsi Singkat</label>
                      <textarea
                        v-model="editForm.bio"
                        rows="2"
                        placeholder="Tuliskan bio atau informasi singkat mengenai Anda"
                        class="w-full px-4 py-2.5 border border-slate-300 dark:border-slate-700 bg-white dark:bg-slate-800 text-slate-900 dark:text-white dark:placeholder-slate-500 rounded-lg focus:ring-2 focus:ring-indigo-500 outline-none text-sm transition"
                      ></textarea>
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
              <div class="bg-white dark:bg-slate-900 rounded-xl border border-slate-200 dark:border-slate-800 shadow-sm overflow-hidden flex flex-col justify-between">
                <div>
                  <div class="px-5 sm:px-6 py-4 border-b border-slate-200 dark:border-slate-800 bg-slate-50 dark:bg-slate-800/60">
                    <h3 class="text-base sm:text-lg font-bold text-slate-900 dark:text-white">Ganti Password</h3>
                    <p class="text-xs text-slate-500 dark:text-slate-400 mt-0.5">Perbarui kata sandi akun Anda</p>
                  </div>

                  <form @submit.prevent="handleChangePassword" class="p-5 sm:p-6 space-y-4">
                    <div>
                      <label class="block text-xs font-semibold text-slate-700 dark:text-slate-300 uppercase tracking-wider mb-1">Password Saat Ini</label>
                      <div class="relative">
                        <input
                          v-model="passwordForm.old_password"
                          :type="showOldPassword ? 'text' : 'password'"
                          required
                          placeholder="••••••••"
                          class="w-full px-4 py-2.5 pr-10 border border-slate-300 dark:border-slate-700 bg-white dark:bg-slate-800 text-slate-900 dark:text-white dark:placeholder-slate-500 rounded-lg focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 outline-none text-sm transition"
                        />
                        <button
                          type="button"
                          @click="showOldPassword = !showOldPassword"
                          class="absolute inset-y-0 right-0 flex items-center pr-3 text-slate-400 hover:text-slate-600 dark:hover:text-slate-200 focus:outline-none"
                          :title="showOldPassword ? 'Sembunyikan password' : 'Lihat password'"
                        >
                          <svg v-if="showOldPassword" class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13.875 18.825A10.05 10.05 0 0112 19c-4.478 0-8.268-2.943-9.543-7a9.97 9.97 0 011.563-3.029m5.858-5.908a8.956 8.956 0 013.122-.763c4.478 0 8.268 2.943 9.543 7a10.025 10.025 0 01-4.132 5.411m0 0L21 21M3 3l18 18" />
                          </svg>
                          <svg v-else class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
                            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z" />
                          </svg>
                        </button>
                      </div>
                    </div>

                    <div>
                      <label class="block text-xs font-semibold text-slate-700 dark:text-slate-300 uppercase tracking-wider mb-1">Password Baru</label>
                      <div class="relative">
                        <input
                          v-model="passwordForm.new_password"
                          :type="showNewPassword ? 'text' : 'password'"
                          required
                          placeholder="••••••••"
                          class="w-full px-4 py-2.5 pr-10 border border-slate-300 dark:border-slate-700 bg-white dark:bg-slate-800 text-slate-900 dark:text-white dark:placeholder-slate-500 rounded-lg focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 outline-none text-sm transition"
                        />
                        <button
                          type="button"
                          @click="showNewPassword = !showNewPassword"
                          class="absolute inset-y-0 right-0 flex items-center pr-3 text-slate-400 hover:text-slate-600 dark:hover:text-slate-200 focus:outline-none"
                          :title="showNewPassword ? 'Sembunyikan password' : 'Lihat password'"
                        >
                          <svg v-if="showNewPassword" class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
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
                      :disabled="isChangingPassword"
                      class="w-full py-3 bg-slate-800 hover:bg-slate-900 dark:bg-slate-700 dark:hover:bg-slate-600 disabled:opacity-50 text-white font-semibold rounded-lg shadow-md transition duration-200 text-sm"
                    >
                      {{ isChangingPassword ? 'Memproses...' : 'Perbarui Password' }}
                    </button>
                  </form>
                </div>
              </div>
            </div>

            <!-- Zona Bahaya / Danger Zone: Soft Delete Akun -->
            <div class="bg-rose-50/60 dark:bg-rose-950/40 rounded-xl border border-rose-200 dark:border-rose-900/60 p-6 flex flex-col sm:flex-row items-start sm:items-center justify-between gap-4">
              <div>
                <h3 class="text-base font-bold text-rose-900 dark:text-rose-300">Zona Bahaya (Danger Zone)</h3>
                <p class="text-xs text-rose-700 dark:text-rose-400 mt-1">Hapus akun Anda secara Soft Delete. Data akun Anda akan dinonaktifkan dari sistem.</p>
              </div>
              <button 
                @click="isDeleteModalOpen = true"
                class="px-4 py-2.5 bg-rose-600 hover:bg-rose-700 text-white text-xs font-semibold rounded-lg shadow transition shrink-0"
              >
                Hapus Akun Saya
              </button>
            </div>
          </div>
        </template>

        <!-- TAB 3: MANAJEMEN USER (KHUSUS SUPERADMIN & OWNER) -->
        <template v-if="activeTab === 'users' && canManageUsers">
          <div class="bg-white dark:bg-slate-900 rounded-xl border border-slate-200 dark:border-slate-800 shadow-sm overflow-hidden">
            <div class="px-5 sm:px-6 py-4 border-b border-slate-200 dark:border-slate-800 bg-slate-50 dark:bg-slate-800/60 flex flex-col sm:flex-row sm:items-center justify-between gap-3">
              <div>
                <h2 class="text-base sm:text-lg font-bold text-slate-900 dark:text-white">Manajemen Seluruh Pengguna Sistem</h2>
                <p class="text-xs text-slate-500 dark:text-slate-400">Khusus Peran Superadmin & Owner (`GET /api/admin/users`)</p>
              </div>
              <div class="flex items-center space-x-2.5 self-start sm:self-auto">
                <button v-if="isSuperadmin" @click="isAddUserModalOpen = true" class="px-3.5 py-2 bg-indigo-600 hover:bg-indigo-700 text-white text-xs font-semibold rounded-lg shadow transition flex items-center space-x-1.5">
                  <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4"/>
                  </svg>
                  <span>Tambah User Baru</span>
                </button>
                <button 
                  type="button"
                  @click.stop="loadAllUsers(true)" 
                  :disabled="isRefreshingUsers"
                  class="px-3.5 py-2 bg-indigo-50 dark:bg-indigo-950/60 text-indigo-600 dark:text-indigo-400 hover:bg-indigo-100 dark:hover:bg-indigo-900/60 disabled:opacity-50 text-xs font-semibold rounded-lg transition flex items-center space-x-1.5 cursor-pointer"
                  title="Klik untuk memuat ulang data pengguna dari database"
                >
                  <svg :class="['w-4 h-4', isRefreshingUsers ? 'animate-spin' : '']" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15" />
                  </svg>
                  <span>{{ isRefreshingUsers ? 'Memuat...' : 'Refresh Data' }}</span>
                </button>
              </div>
            </div>

            <!-- Tabel dengan Touch Scroll di Mobile -->
            <div class="p-4 sm:p-6 overflow-x-auto">
              <table class="w-full text-left border-collapse min-w-[600px]">
                <thead>
                  <tr class="border-b border-slate-200 dark:border-slate-800 bg-slate-50 dark:bg-slate-800/80 text-slate-500 dark:text-slate-400 uppercase text-xs tracking-wider">
                    <th class="py-3 px-4">ID</th>
                    <th class="py-3 px-4">Nama</th>
                    <th class="py-3 px-4">Email</th>
                    <th class="py-3 px-4">Peran (Role)</th>
                    <th class="py-3 px-4">Terdaftar</th>
                    <th v-if="isSuperadmin" class="py-3 px-4 text-center">Aksi Pengelola</th>
                  </tr>
                </thead>
                <tbody class="divide-y divide-slate-100 dark:divide-slate-800 text-sm">
                  <tr v-for="u in allUsers" :key="u.id" class="hover:bg-slate-50 dark:hover:bg-slate-800/50 transition">
                    <td class="py-3 px-4 font-bold text-slate-700 dark:text-slate-300">#{{ u.id }}</td>
                    <td class="py-3 px-4 font-semibold text-slate-800 dark:text-white">{{ u.name }}</td>
                    <td class="py-3 px-4 text-slate-600 dark:text-slate-300">{{ u.email }}</td>
                    <td class="py-3 px-4">
                      <span :class="roleBadgeClass(u.role)">{{ (u.role || 'admin').toUpperCase() }}</span>
                    </td>
                    <td class="py-3 px-4 text-xs text-slate-500 dark:text-slate-400">{{ new Date(u.created_at).toLocaleDateString('id-ID') }}</td>
                    
                    <!-- Tombol Aksi khusus Superadmin (Ubah Role, Edit Data, Hapus User) -->
                    <td v-if="isSuperadmin" class="py-3 px-4 text-center">
                      <div class="flex items-center justify-center space-x-2">
                        <!-- Quick Role Selector -->
                        <select
                          :value="u.role || 'user'"
                          @change="handleChangeRole(u.id, $event.target.value)"
                          class="px-2 py-1 text-xs border border-slate-300 dark:border-slate-700 rounded-md outline-none bg-white dark:bg-slate-800 text-slate-800 dark:text-slate-200 font-medium transition"
                        >
                          <option value="user">USER</option>
                          <option value="admin">ADMIN</option>
                          <option value="owner">OWNER</option>
                          <option value="superadmin">SUPERADMIN</option>
                        </select>

                        <!-- Tombol Edit User -->
                        <button 
                          @click="openAdminEditModal(u)" 
                          class="p-1.5 bg-indigo-50 text-indigo-600 hover:bg-indigo-600 hover:text-white rounded-lg transition"
                          title="Edit Data Pengguna"
                        >
                          <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z"/>
                          </svg>
                        </button>

                        <!-- Tombol Kirim OTP Reset Password -->
                        <button 
                          @click="handleAdminTriggerForgotPassword(u)" 
                          class="p-1.5 bg-amber-50 text-amber-600 hover:bg-amber-500 hover:text-white rounded-lg transition"
                          title="Kirim Kode OTP Reset Password ke Email User"
                          :disabled="isSendingResetOTP"
                        >
                          <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 8l7.89 5.26a2 2 0 002.22 0L21 8M5 19h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z"/>
                          </svg>
                        </button>

                        <!-- Tombol Reset Password Langsung oleh Superadmin (Bypass OTP/Ngebug) -->
                        <button 
                          @click="openAdminDirectResetPasswordModal(u)" 
                          class="p-1.5 bg-emerald-50 text-emerald-600 hover:bg-emerald-600 hover:text-white rounded-lg transition"
                          title="Reset Password Langsung oleh Superadmin (Bantu User)"
                        >
                          <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 15v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2zm10-10V7a4 4 0 00-8 0v4h8z"/>
                          </svg>
                        </button>

                        <!-- Tombol Hapus User -->
                        <button 
                          @click="openAdminDeleteModal(u)" 
                          class="p-1.5 bg-rose-50 text-rose-600 hover:bg-rose-600 hover:text-white rounded-lg transition"
                          title="Hapus Pengguna (Soft Delete)"
                        >
                          <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"/>
                          </svg>
                        </button>
                      </div>
                    </td>
                  </tr>
                </tbody>
              </table>
            </div>
          </div>
        </template>
      </main>
    </div>

    <!-- Modal Tambah User Baru (Khusus Superadmin) -->
    <div v-if="isAddUserModalOpen" class="fixed inset-0 z-50 flex items-center justify-center p-4 bg-slate-900/60 backdrop-blur-sm">
      <div class="w-full max-w-md bg-white dark:bg-slate-900 rounded-xl shadow-2xl border border-slate-200 dark:border-slate-800 overflow-hidden flex flex-col max-h-[90vh] animate-in fade-in zoom-in duration-200">
        <div class="px-6 py-4 border-b border-slate-200 dark:border-slate-800 bg-slate-50 dark:bg-slate-800/60 flex items-center justify-between">
          <div class="flex items-center space-x-2">
            <div class="p-1.5 bg-indigo-100 dark:bg-indigo-950/60 text-indigo-600 dark:text-indigo-400 rounded-lg">
              <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18 9v3m0 0v3m0-3h3m-3 0h-3m-2-5a4 4 0 11-8 0 4 4 0 018 0zM3 20a6 6 0 0112 0v1H3v-1z"/>
              </svg>
            </div>
            <h3 class="text-base sm:text-lg font-bold text-slate-900 dark:text-white">Tambah Pengguna Baru</h3>
          </div>
          <button @click="isAddUserModalOpen = false" class="text-slate-400 hover:text-slate-600 dark:hover:text-slate-200 p-1 rounded-lg hover:bg-slate-100 dark:hover:bg-slate-800 transition">
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"/>
            </svg>
          </button>
        </div>

        <form @submit.prevent="handleCreateUser" class="p-6 space-y-4 overflow-y-auto flex-1">
          <div>
            <label class="block text-xs font-semibold text-slate-700 dark:text-slate-300 uppercase tracking-wider mb-1">Nama Lengkap</label>
            <input
              v-model="addUserForm.name"
              type="text"
              required
              placeholder="Nama Pengguna"
              class="w-full px-4 py-2.5 border border-slate-300 dark:border-slate-700 bg-white dark:bg-slate-800 text-slate-900 dark:text-white dark:placeholder-slate-500 rounded-lg focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 outline-none text-sm transition"
            />
          </div>

          <div>
            <label class="block text-xs font-semibold text-slate-700 dark:text-slate-300 uppercase tracking-wider mb-1">Alamat Email</label>
            <input
              v-model="addUserForm.email"
              type="email"
              required
              placeholder="nama@email.com"
              class="w-full px-4 py-2.5 border border-slate-300 dark:border-slate-700 bg-white dark:bg-slate-800 text-slate-900 dark:text-white dark:placeholder-slate-500 rounded-lg focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 outline-none text-sm transition"
            />
          </div>

          <div>
            <label class="block text-xs font-semibold text-slate-700 dark:text-slate-300 uppercase tracking-wider mb-1">Password</label>
            <div class="relative">
              <input
                v-model="addUserForm.password"
                :type="showAddUserPassword ? 'text' : 'password'"
                required
                placeholder="Minimal 6 karakter"
                class="w-full px-4 py-2.5 pr-10 border border-slate-300 dark:border-slate-700 bg-white dark:bg-slate-800 text-slate-900 dark:text-white dark:placeholder-slate-500 rounded-lg focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 outline-none text-sm transition"
              />
              <button
                type="button"
                @click="showAddUserPassword = !showAddUserPassword"
                class="absolute inset-y-0 right-0 flex items-center pr-3 text-slate-400 hover:text-slate-600 dark:hover:text-slate-200 focus:outline-none"
                :title="showAddUserPassword ? 'Sembunyikan password' : 'Lihat password'"
              >
                <svg v-if="showAddUserPassword" class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13.875 18.825A10.05 10.05 0 0112 19c-4.478 0-8.268-2.943-9.543-7a9.97 9.97 0 011.563-3.029m5.858-5.908a8.956 8.956 0 013.122-.763c4.478 0 8.268 2.943 9.543 7a10.025 10.025 0 01-4.132 5.411m0 0L21 21M3 3l18 18" />
                </svg>
                <svg v-else class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z" />
                </svg>
              </button>
            </div>
          </div>

          <div>
            <label class="block text-xs font-semibold text-slate-700 dark:text-slate-300 uppercase tracking-wider mb-1">Peran (Role)</label>
            <select
              v-model="addUserForm.role"
              class="w-full px-4 py-2.5 border border-slate-300 dark:border-slate-700 bg-white dark:bg-slate-800 text-slate-900 dark:text-white rounded-lg focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 outline-none text-sm font-medium transition"
            >
              <option value="user">USER (Pengguna Biasa)</option>
              <option value="admin">ADMIN (Admin Standar)</option>
              <option value="owner">OWNER (Pemilik Laporan)</option>
              <option value="superadmin">SUPERADMIN (Hak Akses Penuh)</option>
            </select>
          </div>

          <div class="flex items-center justify-end space-x-3 pt-2">
            <button
              type="button"
              @click="isAddUserModalOpen = false"
              class="px-4 py-2.5 border border-slate-300 dark:border-slate-700 text-slate-700 dark:text-slate-300 font-semibold rounded-lg hover:bg-slate-50 dark:hover:bg-slate-800 transition text-sm"
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

    <!-- Modal Konfirmasi Hapus Akun Mandiri (Soft Delete) -->
    <div v-if="isDeleteModalOpen" class="fixed inset-0 z-50 flex items-center justify-center p-4 bg-slate-900/60 backdrop-blur-sm">
      <div class="w-full max-w-md bg-white dark:bg-slate-900 rounded-xl shadow-2xl border border-slate-200 dark:border-slate-800 overflow-hidden p-6 space-y-4 animate-in fade-in zoom-in duration-200">
        <div class="flex items-center space-x-3 text-rose-600 dark:text-rose-400">
          <div class="p-2 bg-rose-100 dark:bg-rose-950/60 rounded-full">
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z"/>
            </svg>
          </div>
          <h3 class="text-lg font-bold text-slate-900 dark:text-white">Konfirmasi Hapus Akun</h3>
        </div>
        <p class="text-sm text-slate-600 dark:text-slate-300">Apakah Anda yakin ingin menghapus akun Anda? Setelah dihapus, Anda akan otomatis ter-logout dan data akun akan dinonaktifkan (Soft Delete).</p>
        <div class="flex items-center justify-end space-x-3 pt-2">
          <button @click="isDeleteModalOpen = false" class="px-4 py-2 border border-slate-300 dark:border-slate-700 text-slate-700 dark:text-slate-300 font-semibold rounded-lg hover:bg-slate-50 dark:hover:bg-slate-800 transition text-sm">
            Batal
          </button>
          <button @click="handleDeleteAccount" :disabled="isDeletingAccount" class="px-4 py-2 bg-rose-600 hover:bg-rose-700 disabled:opacity-50 text-white font-semibold rounded-lg shadow transition text-sm">
            {{ isDeletingAccount ? 'Menghapus...' : 'Ya, Hapus Akun' }}
          </button>
        </div>
      </div>
    </div>

    <!-- Modal Edit Data Pengguna (Khusus Super Admin) -->
    <div v-if="isAdminEditModalOpen" class="fixed inset-0 z-50 flex items-center justify-center p-4 bg-slate-900/60 backdrop-blur-sm">
      <div class="w-full max-w-md bg-white dark:bg-slate-900 rounded-xl shadow-2xl border border-slate-200 dark:border-slate-800 overflow-hidden animate-in fade-in zoom-in duration-200">
        <div class="px-6 py-4 border-b border-slate-200 dark:border-slate-800 bg-slate-50 dark:bg-slate-800/60 flex items-center justify-between">
          <h3 class="text-base sm:text-lg font-bold text-slate-900 dark:text-white">Edit Data Pengguna #{{ adminEditForm.user_id }}</h3>
          <button @click="isAdminEditModalOpen = false" class="text-slate-400 hover:text-slate-600 dark:hover:text-slate-200 p-1 rounded-lg">
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"/>
            </svg>
          </button>
        </div>

        <form @submit.prevent="handleAdminUpdateUser" class="p-6 space-y-4 max-h-[80vh] overflow-y-auto">
          <div>
            <label class="block text-xs font-semibold text-slate-700 dark:text-slate-300 uppercase tracking-wider mb-1">Nama Lengkap</label>
            <input
              v-model="adminEditForm.name"
              type="text"
              required
              class="w-full px-4 py-2.5 border border-slate-300 dark:border-slate-700 bg-white dark:bg-slate-800 text-slate-900 dark:text-white rounded-lg focus:ring-2 focus:ring-indigo-500 outline-none text-sm transition"
            />
          </div>

          <div>
            <label class="block text-xs font-semibold text-slate-700 dark:text-slate-300 uppercase tracking-wider mb-1">Alamat Email</label>
            <input
              v-model="adminEditForm.email"
              type="email"
              required
              class="w-full px-4 py-2.5 border border-slate-300 dark:border-slate-700 bg-white dark:bg-slate-800 text-slate-900 dark:text-white rounded-lg focus:ring-2 focus:ring-indigo-500 outline-none text-sm transition"
            />
          </div>

          <div class="grid grid-cols-2 gap-3">
            <div>
              <label class="block text-xs font-semibold text-slate-700 dark:text-slate-300 uppercase tracking-wider mb-1">Peran (Role)</label>
              <select
                v-model="adminEditForm.role"
                class="w-full px-4 py-2.5 border border-slate-300 dark:border-slate-700 bg-white dark:bg-slate-800 text-slate-900 dark:text-white rounded-lg focus:ring-2 focus:ring-indigo-500 outline-none text-sm font-medium transition"
              >
                <option value="user">USER (Pengguna Biasa)</option>
                <option value="admin">ADMIN (Admin Standar)</option>
                <option value="owner">OWNER (Pemilik Laporan)</option>
                <option value="superadmin">SUPERADMIN (Hak Akses Penuh)</option>
              </select>
            </div>

            <div>
              <label class="block text-xs font-semibold text-slate-700 dark:text-slate-300 uppercase tracking-wider mb-1">Nomor Telepon / HP</label>
              <input
                v-model="adminEditForm.phone"
                type="text"
                placeholder="081234567890"
                class="w-full px-4 py-2.5 border border-slate-300 dark:border-slate-700 bg-white dark:bg-slate-800 text-slate-900 dark:text-white dark:placeholder-slate-500 rounded-lg focus:ring-2 focus:ring-indigo-500 outline-none text-sm transition"
              />
            </div>
          </div>

          <div class="grid grid-cols-2 gap-3">
            <div>
              <label class="block text-xs font-semibold text-slate-700 dark:text-slate-300 uppercase tracking-wider mb-1">Jenis Kelamin</label>
              <select
                v-model="adminEditForm.gender"
                class="w-full px-4 py-2.5 border border-slate-300 dark:border-slate-700 bg-white dark:bg-slate-800 text-slate-900 dark:text-white rounded-lg focus:ring-2 focus:ring-indigo-500 outline-none text-sm font-medium transition"
              >
                <option value="">-- Pilih --</option>
                <option value="Laki-laki">Laki-laki</option>
                <option value="Perempuan">Perempuan</option>
              </select>
            </div>

            <div>
              <label class="block text-xs font-semibold text-slate-700 dark:text-slate-300 uppercase tracking-wider mb-1">Tanggal Lahir</label>
              <input
                v-model="adminEditForm.birth_date"
                type="date"
                class="w-full px-4 py-2.5 border border-slate-300 dark:border-slate-700 bg-white dark:bg-slate-800 text-slate-900 dark:text-white rounded-lg focus:ring-2 focus:ring-indigo-500 outline-none text-sm transition"
              />
            </div>
          </div>

          <div>
            <label class="block text-xs font-semibold text-slate-700 dark:text-slate-300 uppercase tracking-wider mb-1">Alamat Lengkap</label>
            <textarea
              v-model="adminEditForm.address"
              rows="2"
              placeholder="Alamat domisili"
              class="w-full px-4 py-2.5 border border-slate-300 dark:border-slate-700 bg-white dark:bg-slate-800 text-slate-900 dark:text-white dark:placeholder-slate-500 rounded-lg focus:ring-2 focus:ring-indigo-500 outline-none text-sm transition"
            ></textarea>
          </div>

          <div>
            <label class="block text-xs font-semibold text-slate-700 dark:text-slate-300 uppercase tracking-wider mb-1">Bio / Deskripsi</label>
            <textarea
              v-model="adminEditForm.bio"
              rows="2"
              placeholder="Catatan / Bio pengguna"
              class="w-full px-4 py-2.5 border border-slate-300 dark:border-slate-700 bg-white dark:bg-slate-800 text-slate-900 dark:text-white dark:placeholder-slate-500 rounded-lg focus:ring-2 focus:ring-indigo-500 outline-none text-sm transition"
            ></textarea>
          </div>

          <div class="flex items-center justify-end space-x-3 pt-2">
            <button
              type="button"
              @click="isAdminEditModalOpen = false"
              class="px-4 py-2 border border-slate-300 dark:border-slate-700 text-slate-700 dark:text-slate-300 font-semibold rounded-lg hover:bg-slate-50 dark:hover:bg-slate-800 text-sm transition"
            >
              Batal
            </button>
            <button
              type="submit"
              :disabled="isAdminUpdating"
              class="px-4 py-2 bg-indigo-600 hover:bg-indigo-700 disabled:opacity-50 text-white font-semibold rounded-lg shadow text-sm transition"
            >
              {{ isAdminUpdating ? 'Menyimpan...' : 'Simpan Perubahan' }}
            </button>
          </div>
        </form>
      </div>
    </div>

    <!-- Modal Konfirmasi Hapus Pengguna (Khusus Super Admin) -->
    <div v-if="isAdminDeleteModalOpen" class="fixed inset-0 z-50 flex items-center justify-center p-4 bg-slate-900/60 backdrop-blur-sm">
      <div class="w-full max-w-md bg-white dark:bg-slate-900 rounded-xl shadow-2xl border border-slate-200 dark:border-slate-800 overflow-hidden p-6 space-y-4 animate-in fade-in zoom-in duration-200">
        <div class="flex items-center space-x-3 text-rose-600 dark:text-rose-400">
          <div class="p-2 bg-rose-100 dark:bg-rose-950/60 rounded-full">
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z"/>
            </svg>
          </div>
          <h3 class="text-lg font-bold text-slate-900 dark:text-white">Hapus Pengguna</h3>
        </div>
        <p class="text-sm text-slate-600 dark:text-slate-300">Apakah Anda yakin ingin menghapus pengguna <strong class="text-slate-900 dark:text-white">{{ adminDeleteTarget.name }}</strong> (ID: #{{ adminDeleteTarget.id }})? Pengguna ini akan di-Soft Delete dari sistem.</p>
        <div class="flex items-center justify-end space-x-3 pt-2">
          <button @click="isAdminDeleteModalOpen = false" class="px-4 py-2 border border-slate-300 dark:border-slate-700 text-slate-700 dark:text-slate-300 font-semibold rounded-lg hover:bg-slate-50 dark:hover:bg-slate-800 text-sm transition">
            Batal
          </button>
          <button @click="handleAdminDeleteUser" :disabled="isAdminDeleting" class="px-4 py-2 bg-rose-600 hover:bg-rose-700 disabled:opacity-50 text-white font-semibold rounded-lg shadow text-sm">
            {{ isAdminDeleting ? 'Menghapus...' : 'Ya, Hapus Pengguna' }}
          </button>
        </div>
      </div>
    </div>

    <!-- Modal Direct Reset Password oleh Superadmin (Untuk Membantu User jika OTP Bug/Kendala) -->
    <div v-if="isAdminDirectResetModalOpen" class="fixed inset-0 z-50 flex items-center justify-center p-4 bg-slate-900/60 backdrop-blur-sm">
      <div class="w-full max-w-md bg-white dark:bg-slate-900 rounded-xl shadow-2xl border border-slate-200 dark:border-slate-800 overflow-hidden p-6 space-y-4 animate-in fade-in zoom-in duration-200">
        <div class="flex items-center space-x-3 text-emerald-600 dark:text-emerald-400">
          <div class="p-2 bg-emerald-100 dark:bg-emerald-950/60 rounded-full">
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 15v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2zm10-10V7a4 4 0 00-8 0v4h8z"/>
            </svg>
          </div>
          <h3 class="text-lg font-bold text-slate-900 dark:text-white">Reset Password Langsung</h3>
        </div>
        <p class="text-xs sm:text-sm text-slate-600 dark:text-slate-300">
          Bantu reset password pengguna <strong class="text-slate-900 dark:text-white">{{ adminResetTarget.name }}</strong> ({{ adminResetTarget.email }}) secara langsung tanpa menggunakan kode OTP.
        </p>

        <form @submit.prevent="handleAdminDirectResetPassword" class="space-y-4 pt-1">
          <div>
            <label class="block text-xs font-semibold text-slate-700 dark:text-slate-300 uppercase tracking-wider mb-1">Password Baru Pengguna</label>
            <div class="relative">
              <input
                v-model="adminResetNewPassword"
                :type="showAdminResetPassword ? 'text' : 'password'"
                required
                placeholder="Minimal 6 karakter"
                class="w-full px-4 py-2.5 pr-10 border border-slate-300 dark:border-slate-700 bg-white dark:bg-slate-800 text-slate-900 dark:text-white dark:placeholder-slate-500 rounded-lg focus:ring-2 focus:ring-emerald-500 focus:border-emerald-500 outline-none text-sm transition"
              />
              <button
                type="button"
                @click="showAdminResetPassword = !showAdminResetPassword"
                class="absolute inset-y-0 right-0 flex items-center pr-3 text-slate-400 hover:text-slate-600 dark:hover:text-slate-200 focus:outline-none"
                :title="showAdminResetPassword ? 'Sembunyikan password' : 'Lihat password'"
              >
                <svg v-if="showAdminResetPassword" class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13.875 18.825A10.05 10.05 0 0112 19c-4.478 0-8.268-2.943-9.543-7a9.97 9.97 0 011.563-3.029m5.858-5.908a8.956 8.956 0 013.122-.763c4.478 0 8.268 2.943 9.543 7a10.025 10.025 0 01-4.132 5.411m0 0L21 21M3 3l18 18" />
                </svg>
                <svg v-else class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z" />
                </svg>
              </button>
            </div>
          </div>

          <div class="flex items-center justify-end space-x-3 pt-2">
            <button
              type="button"
              @click="isAdminDirectResetModalOpen = false"
              class="px-4 py-2 border border-slate-300 dark:border-slate-700 text-slate-700 dark:text-slate-300 font-semibold rounded-lg hover:bg-slate-50 dark:hover:bg-slate-800 text-sm transition"
            >
              Batal
            </button>
            <button
              type="submit"
              :disabled="isAdminResettingPassword"
              class="px-4 py-2 bg-emerald-600 hover:bg-emerald-700 disabled:opacity-50 text-white font-semibold rounded-lg shadow text-sm transition"
            >
              {{ isAdminResettingPassword ? 'Mereset...' : 'Reset Password Sekarang' }}
            </button>
          </div>
        </form>
      </div>
    </div>

    <!-- Modal Pop-up Foto Profil & Download Avatar -->
    <div v-if="isAvatarModalOpen" class="fixed inset-0 z-50 flex items-center justify-center p-4 bg-slate-900/80 backdrop-blur-md animate-in fade-in duration-200">
      <div class="w-full max-w-lg bg-white dark:bg-slate-900 rounded-2xl shadow-2xl overflow-hidden border border-slate-200 dark:border-slate-800 p-6 flex flex-col items-center space-y-5 animate-in zoom-in duration-200 relative">
        <button @click="isAvatarModalOpen = false" class="absolute top-4 right-4 text-slate-400 hover:text-slate-600 dark:hover:text-slate-200 p-1.5 rounded-full hover:bg-slate-100 dark:hover:bg-slate-800 transition">
          <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"/>
          </svg>
        </button>

        <div class="text-center">
          <h3 class="text-lg font-bold text-slate-900 dark:text-white">{{ avatarModalName }}</h3>
          <p class="text-xs text-slate-500 dark:text-slate-400">Pratinjau Foto Profil (Avatar)</p>
        </div>

        <!-- High-res Image Preview -->
        <div class="w-64 h-64 sm:w-72 sm:h-72 rounded-full overflow-hidden border-4 border-indigo-100 dark:border-indigo-900/60 shadow-xl bg-indigo-50 dark:bg-indigo-950 flex items-center justify-center">
          <img v-if="avatarModalUrl" :src="avatarModalUrl" alt="Avatar Pop-up" class="w-full h-full object-cover" />
          <div v-else class="w-full h-full bg-indigo-600 text-white font-extrabold text-6xl flex items-center justify-center">
            {{ avatarInitial }}
          </div>
        </div>

        <!-- Download & Action Buttons -->
        <div class="flex items-center space-x-3 w-full pt-2">
          <button 
            @click="handleDownloadAvatar" 
            :disabled="!avatarModalUrl"
            class="flex-1 py-2.5 bg-indigo-600 hover:bg-indigo-700 disabled:opacity-50 text-white text-xs sm:text-sm font-semibold rounded-xl shadow transition flex items-center justify-center space-x-2"
          >
            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16v1a3 3 0 003 3h10a3 3 0 003-3v-1m-4-4l-4 4m0 0l-4-4m4 4V4"/>
            </svg>
            <span>Unduh Foto Profil</span>
          </button>
          <button 
            @click="isAvatarModalOpen = false" 
            class="px-5 py-2.5 border border-slate-300 dark:border-slate-700 text-slate-700 dark:text-slate-300 text-xs sm:text-sm font-semibold rounded-xl hover:bg-slate-50 dark:hover:bg-slate-800 transition"
          >
            Tutup
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive, computed, onMounted, onUnmounted, watch, nextTick } from 'vue'
import { useRouter } from 'vue-router'
import api from '../services/api'
import Chart from 'chart.js/auto'
import { useTheme } from '../utils/theme'

const { isDarkMode, toggleTheme } = useTheme()
const router = useRouter()
const activeTab = ref('dashboard')
const isDropdownOpen = ref(false)
const isMobileSidebarOpen = ref(false)
const isUpdating = ref(false)
const isChangingPassword = ref(false)
const isAddUserModalOpen = ref(false)
const isAddingUser = ref(false)
const isUploadingAvatar = ref(false)
const isDeleteModalOpen = ref(false)
const isDeletingAccount = ref(false)
const isAdminEditModalOpen = ref(false)
const isAdminUpdating = ref(false)
const isAdminDeleteModalOpen = ref(false)
const isAdminDeleting = ref(false)
const isSyncing = ref(false)
const isAvatarModalOpen = ref(false)
const avatarModalName = ref('')
const avatarModalUrl = ref('')
const selectedFile = ref(null)
const allUsers = ref([])

const stats = reactive({
  total_users: 0,
  total_superadmin: 0,
  total_owner: 0,
  total_admin: 0,
  total_user_role: 0
})

const roleDoughnutChartCanvas = ref(null)
const roleBarChartCanvas = ref(null)
let doughnutChartInstance = null
let barChartInstance = null

const isNotificationOpen = ref(false)
const notifications = ref([
  {
    id: 1,
    title: 'Sesi Autentikasi Berhasil',
    message: 'Login sebagai SUPERADMIN dengan Dual Token JWT & Redis Blacklist.',
    type: 'security',
    read: false,
    created_at: new Date().toISOString()
  },
  {
    id: 2,
    title: 'Profil Terintegrasi',
    message: 'Biodata profil lengkap (No. HP, Gender, Tgl Lahir, Alamat, Bio) telah disinkronkan.',
    type: 'profile',
    read: false,
    created_at: new Date(Date.now() - 5 * 60 * 1000).toISOString()
  },
  {
    id: 3,
    title: 'Database AutoMigrate Active',
    message: 'PostgreSQL & Seeder 3 Akun Default (Super Admin, Owner, Admin) berjalan normal.',
    type: 'system',
    read: false,
    created_at: new Date(Date.now() - 15 * 60 * 1000).toISOString()
  }
])

const unreadCount = computed(() => {
  return notifications.value.filter(n => !n.read).length
})

const markAllAsRead = () => {
  notifications.value.forEach(n => n.read = true)
}

const clearNotifications = () => {
  notifications.value = []
}

const markAsRead = (id) => {
  const item = notifications.value.find(n => n.id === id)
  if (item) item.read = true
}

const formatRelativeTime = (dateStr) => {
  if (!dateStr) return ''
  const diffSec = Math.floor((new Date() - new Date(dateStr)) / 1000)
  if (diffSec < 60) return 'Baru saja'
  if (diffSec < 3600) return `${Math.floor(diffSec / 60)} mnt lalu`
  if (diffSec < 86400) return `${Math.floor(diffSec / 3600)} jam lalu`
  return `${Math.floor(diffSec / 86400)} hr lalu`
}

const addUserForm = reactive({
  name: '',
  email: '',
  password: '',
  role: 'admin'
})

const adminEditForm = reactive({
  user_id: 0,
  name: '',
  email: '',
  role: 'admin',
  phone: '',
  gender: '',
  birth_date: '',
  address: '',
  bio: ''
})

const adminDeleteTarget = reactive({
  id: 0,
  name: ''
})

const user = reactive({
  id: '',
  name: '',
  email: '',
  role: 'admin',
  avatar: '',
  phone: '',
  gender: '',
  birth_date: '',
  address: '',
  bio: '',
  created_at: '',
  updated_at: ''
})

const editForm = reactive({
  name: '',
  email: '',
  phone: '',
  gender: '',
  birth_date: '',
  address: '',
  bio: ''
})

const passwordForm = reactive({
  old_password: '',
  new_password: ''
})

const showOldPassword = ref(false)
const showNewPassword = ref(false)
const showAddUserPassword = ref(false)
const isSendingResetOTP = ref(false)
const isRefreshingUsers = ref(false)

const isAdminDirectResetModalOpen = ref(false)
const isAdminResettingPassword = ref(false)
const showAdminResetPassword = ref(false)
const adminResetNewPassword = ref('')
const adminResetTarget = reactive({
  id: 0,
  name: '',
  email: ''
})

const alert = reactive({
  message: '',
  isSuccess: true
})

const avatarInitial = computed(() => {
  return user.name ? user.name.charAt(0).toUpperCase() : 'U'
})

const avatarUrl = computed(() => {
  if (user.avatar) {
    if (user.avatar.startsWith('http://') || user.avatar.startsWith('https://')) {
      return user.avatar
    }
    return `http://localhost:8080/${user.avatar}`
  }
  return null
})

const formattedUpdatedAt = computed(() => {
  if (!user.updated_at || user.updated_at.startsWith('0001-01-01')) {
    return 'Belum pernah diperbarui'
  }
  try {
    const d = new Date(user.updated_at)
    if (isNaN(d.getTime())) return 'Belum pernah diperbarui'
    return d.toLocaleString('id-ID', {
      day: 'numeric',
      month: 'long',
      year: 'numeric',
      hour: '2-digit',
      minute: '2-digit'
    })
  } catch (e) {
    return 'Belum pernah diperbarui'
  }
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
    case 'admin':
      return 'px-2 py-0.5 rounded text-[10px] font-extrabold bg-blue-100 text-blue-800 border border-blue-200'
    case 'user':
    default:
      return 'px-2 py-0.5 rounded text-[10px] font-extrabold bg-emerald-100 text-emerald-800 border border-emerald-200'
  }
}

const showAlert = (message, isSuccess = true) => {
  alert.message = message
  alert.isSuccess = isSuccess

  notifications.value.unshift({
    id: Date.now(),
    title: isSuccess ? 'Aktivitas Berhasil' : 'Peringatan Sistem',
    message: message,
    type: isSuccess ? 'profile' : 'security',
    read: false,
    created_at: new Date().toISOString()
  })

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
    user.avatar = data.avatar || ''
    user.phone = data.phone || ''
    user.gender = data.gender || ''
    user.birth_date = data.birth_date || ''
    user.address = data.address || ''
    user.bio = data.bio || ''
    user.created_at = data.created_at
    user.updated_at = data.updated_at

    editForm.name = data.name
    editForm.email = data.email
    editForm.phone = data.phone || ''
    editForm.gender = data.gender || ''
    editForm.birth_date = data.birth_date || ''
    editForm.address = data.address || ''
    editForm.bio = data.bio || ''

    notifications.value = [
      {
        id: 1,
        title: 'Sesi Autentikasi Berhasil',
        message: `Login sebagai ${user.name} (${(user.role || 'admin').toUpperCase()}) dengan Dual Token JWT & Redis Blacklist.`,
        type: 'security',
        read: false,
        created_at: new Date().toISOString()
      },
      {
        id: 2,
        title: 'Profil Terintegrasi',
        message: 'Biodata profil lengkap (No. HP, Gender, Tgl Lahir, Alamat, Bio) telah disinkronkan.',
        type: 'profile',
        read: false,
        created_at: new Date(Date.now() - 5 * 60 * 1000).toISOString()
      },
      {
        id: 3,
        title: 'Database AutoMigrate Active',
        message: 'PostgreSQL & Seeder 3 Akun Default (Super Admin, Owner, Admin) berjalan normal.',
        type: 'system',
        read: false,
        created_at: new Date(Date.now() - 15 * 60 * 1000).toISOString()
      }
    ]
  } catch (err) {
    localStorage.removeItem('token')
    router.push('/login')
  }
}

const loadAllUsers = async (showNotification = false) => {
  if (!canManageUsers.value) return
  isRefreshingUsers.value = true
  try {
    const res = await api.get('/admin/users')
    allUsers.value = res.data.data
    if (showNotification) {
      showAlert('Data pengguna berhasil diperbarui! 🔄', true)
    }
  } catch (err) {
    showAlert(err.response?.data?.message || 'Akses ditolak untuk mengambil data pengguna', false)
  } finally {
    isRefreshingUsers.value = false
  }
}

watch(activeTab, (newTab) => {
  if (newTab === 'users') {
    loadAllUsers()
  }
})

const handleFileSelected = (event) => {
  const file = event.target.files[0]
  if (file) {
    selectedFile.value = file
  }
}

const handleUploadAvatar = async () => {
  if (!selectedFile.value) return
  isUploadingAvatar.value = true
  const formData = new FormData()
  formData.append('avatar', selectedFile.value)

  try {
    const res = await api.post('/profile/avatar', formData, {
      headers: {
        'Content-Type': 'multipart/form-data'
      }
    })

    user.avatar = res.data.data.avatar
    showAlert('Foto profil berhasil diperbarui!', true)
    selectedFile.value = null
  } catch (err) {
    showAlert(err.response?.data?.message || 'Gagal mengunggah foto profil', false)
  } finally {
    isUploadingAvatar.value = false
  }
}

const handleUpdateProfile = async () => {
  isUpdating.value = true
  try {
    const res = await api.put('/profile', {
      name: editForm.name,
      email: editForm.email,
      phone: editForm.phone,
      gender: editForm.gender,
      birth_date: editForm.birth_date,
      address: editForm.address,
      bio: editForm.bio
    })

    const d = res.data.data
    user.name = d.name
    user.email = d.email
    user.phone = d.phone || ''
    user.gender = d.gender || ''
    user.birth_date = d.birth_date || ''
    user.address = d.address || ''
    user.bio = d.bio || ''
    user.updated_at = d.updated_at
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

const handleDeleteAccount = async () => {
  isDeletingAccount.value = true
  try {
    await api.delete('/profile')
    showAlert('Akun Anda berhasil dihapus (Soft Delete).', true)
    setTimeout(() => {
      localStorage.removeItem('token')
      localStorage.removeItem('access_token')
      localStorage.removeItem('refresh_token')
      router.push('/login')
    }, 1500)
  } catch (err) {
    showAlert(err.response?.data?.message || 'Gagal menghapus akun', false)
  } finally {
    isDeletingAccount.value = false
    isDeleteModalOpen.value = false
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

    addUserForm.name = ''
    addUserForm.email = ''
    addUserForm.password = ''
    addUserForm.role = 'admin'
    isAddUserModalOpen.value = false

    loadAllUsers()
  } catch (err) {
    showAlert(err.response?.data?.message || 'Gagal menambahkan pengguna baru', false)
  } finally {
    isAddingUser.value = false
  }
}

const openAdminEditModal = (targetUser) => {
  adminEditForm.user_id = targetUser.id
  adminEditForm.name = targetUser.name
  adminEditForm.email = targetUser.email
  adminEditForm.role = targetUser.role || 'admin'
  adminEditForm.phone = targetUser.phone || ''
  adminEditForm.gender = targetUser.gender || ''
  adminEditForm.birth_date = targetUser.birth_date || ''
  adminEditForm.address = targetUser.address || ''
  adminEditForm.bio = targetUser.bio || ''
  isAdminEditModalOpen.value = true
}

const handleAdminUpdateUser = async () => {
  isAdminUpdating.value = true
  try {
    const res = await api.put('/superadmin/users', {
      user_id: adminEditForm.user_id,
      name: adminEditForm.name,
      email: adminEditForm.email,
      role: adminEditForm.role,
      phone: adminEditForm.phone,
      gender: adminEditForm.gender,
      birth_date: adminEditForm.birth_date,
      address: adminEditForm.address,
      bio: adminEditForm.bio
    })
    showAlert(`Data pengguna '${res.data.data.name}' berhasil diperbarui!`, true)
    isAdminEditModalOpen.value = false
    loadAllUsers()
  } catch (err) {
    showAlert(err.response?.data?.message || 'Gagal mengedit data pengguna', false)
  } finally {
    isAdminUpdating.value = false
  }
}

const openAdminDeleteModal = (targetUser) => {
  adminDeleteTarget.id = targetUser.id
  adminDeleteTarget.name = targetUser.name
  isAdminDeleteModalOpen.value = true
}

const handleAdminDeleteUser = async () => {
  isAdminDeleting.value = true
  try {
    await api.delete(`/superadmin/users/${adminDeleteTarget.id}`)
    showAlert(`Pengguna '${adminDeleteTarget.name}' berhasil dihapus (Soft Delete)!`, true)
    isAdminDeleteModalOpen.value = false
    loadAllUsers()
  } catch (err) {
    showAlert(err.response?.data?.message || 'Gagal menghapus pengguna', false)
  } finally {
    isAdminDeleting.value = false
  }
}

const handleAdminTriggerForgotPassword = async (targetUser) => {
  if (!confirm(`Kirim kode OTP Reset Password ke email pengguna '${targetUser.name}' (${targetUser.email})?`)) {
    return
  }
  isSendingResetOTP.value = true
  try {
    const res = await api.post('/forgot-password', { email: targetUser.email })
    showAlert(res.data?.message || `Kode OTP Reset Password berhasil dikirim ke email '${targetUser.email}'!`, true)
  } catch (err) {
    showAlert(err.response?.data?.message || 'Gagal mengirim kode OTP reset password', false)
  } finally {
    isSendingResetOTP.value = false
  }
}

const openAdminDirectResetPasswordModal = (targetUser) => {
  adminResetTarget.id = targetUser.id
  adminResetTarget.name = targetUser.name
  adminResetTarget.email = targetUser.email
  adminResetNewPassword.value = ''
  isAdminDirectResetModalOpen.value = true
}

const handleAdminDirectResetPassword = async () => {
  if (!adminResetNewPassword.value) return
  isAdminResettingPassword.value = true
  try {
    const res = await api.put('/superadmin/users/reset-password', {
      user_id: adminResetTarget.id,
      new_password: adminResetNewPassword.value
    })
    showAlert(res.data?.message || `Password pengguna '${adminResetTarget.name}' berhasil diperbarui oleh Super Admin!`, true)
    isAdminDirectResetModalOpen.value = false
  } catch (err) {
    showAlert(err.response?.data?.message || 'Gagal mereset password pengguna', false)
  } finally {
    isAdminResettingPassword.value = false
  }
}

const loadDashboardStats = async () => {
  try {
    const res = await api.get('/dashboard/stats')
    const d = res.data.data
    stats.total_users = d.total_users || 0
    stats.total_superadmin = d.total_superadmin || 0
    stats.total_owner = d.total_owner || 0
    stats.total_admin = d.total_admin || 0
    stats.total_user_role = d.total_user_role || 0

    await nextTick()
    renderCharts()
  } catch (err) {
    // Ignore error
  }
}

const renderCharts = () => {
  if (doughnutChartInstance) doughnutChartInstance.destroy()
  if (barChartInstance) barChartInstance.destroy()

  const labels = ['SUPERADMIN', 'OWNER', 'ADMIN', 'USER']
  const dataValues = [stats.total_superadmin, stats.total_owner, stats.total_admin, stats.total_user_role]
  const colors = ['#e11d48', '#9333ea', '#2563eb', '#10b981']
  const textColor = isDarkMode.value ? '#cbd5e1' : '#475569'
  const borderColor = isDarkMode.value ? '#0f172a' : '#ffffff'
  const gridColor = isDarkMode.value ? 'rgba(255, 255, 255, 0.1)' : 'rgba(0, 0, 0, 0.05)'

  if (roleDoughnutChartCanvas.value) {
    doughnutChartInstance = new Chart(roleDoughnutChartCanvas.value, {
      type: 'doughnut',
      data: {
        labels: labels,
        datasets: [{
          data: dataValues,
          backgroundColor: colors,
          borderWidth: 2,
          borderColor: borderColor
        }]
      },
      options: {
        responsive: true,
        maintainAspectRatio: false,
        plugins: {
          legend: {
            position: 'bottom',
            labels: { color: textColor, font: { size: 12, weight: 'bold' }, padding: 15 }
          }
        }
      }
    })
  }

  if (roleBarChartCanvas.value) {
    barChartInstance = new Chart(roleBarChartCanvas.value, {
      type: 'bar',
      data: {
        labels: labels,
        datasets: [{
          label: 'Jumlah Pengguna',
          data: dataValues,
          backgroundColor: ['rgba(225, 29, 72, 0.85)', 'rgba(147, 51, 234, 0.85)', 'rgba(37, 99, 235, 0.85)', 'rgba(16, 185, 129, 0.85)'],
          borderRadius: 8,
          borderWidth: 0
        }]
      },
      options: {
        responsive: true,
        maintainAspectRatio: false,
        scales: {
          x: { ticks: { color: textColor }, grid: { color: gridColor } },
          y: { beginAtZero: true, ticks: { precision: 0, color: textColor }, grid: { color: gridColor } }
        },
        plugins: {
          legend: { display: false }
        }
      }
    })
  }
}

watch(isDarkMode, () => {
  renderCharts()
})

watch(activeTab, (newTab) => {
  if (newTab === 'dashboard') {
    loadDashboardStats()
  } else if (newTab === 'users') {
    loadAllUsers()
  }
})

const handleSyncData = async () => {
  isSyncing.value = true
  try {
    await loadUserProfile()
    await loadDashboardStats()
    if (canManageUsers.value) {
      await loadAllUsers()
    }
    showAlert('Data & statistik berhasil disinkronkan dari database PostgreSQL!', true)
  } catch (err) {
    showAlert('Gagal menyinkronkan data dari server', false)
  } finally {
    isSyncing.value = false
  }
}

const openAvatarModal = (name, url) => {
  avatarModalName.value = name || user.name || 'Pengguna'
  avatarModalUrl.value = url || avatarUrl.value || ''
  isAvatarModalOpen.value = true
}

const handleDownloadAvatar = () => {
  if (!avatarModalUrl.value) {
    showAlert('Foto profil tidak tersedia untuk diunduh', false)
    return
  }

  const link = document.createElement('a')
  link.href = avatarModalUrl.value
  link.download = `${(avatarModalName.value || 'profile').replace(/\s+/g, '_')}_avatar.png`
  link.target = '_blank'
  document.body.appendChild(link)
  link.click()
  document.body.removeChild(link)
  showAlert(`Foto profil '${avatarModalName.value}' berhasil diunduh!`, true)
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
  isNotificationOpen.value = false
}

onMounted(() => {
  loadUserProfile()
  loadDashboardStats()
  window.addEventListener('click', closeDropdownOnOutsideClick)
})

onUnmounted(() => {
  if (doughnutChartInstance) doughnutChartInstance.destroy()
  if (barChartInstance) barChartInstance.destroy()
  window.removeEventListener('click', closeDropdownOnOutsideClick)
})
</script>

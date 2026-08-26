import { ref, onMounted } from 'vue'

function getInitialTheme() {
  const savedTheme = localStorage.getItem('theme')
  if (savedTheme === 'dark') return true
  if (savedTheme === 'light') return false
  if (typeof window !== 'undefined') {
    return document.documentElement.classList.contains('dark') || 
           window.matchMedia('(prefers-color-scheme: dark)').matches
  }
  return false
}

export const isDarkMode = ref(getInitialTheme())

export function setTheme(dark) {
  isDarkMode.value = dark
  if (typeof document !== 'undefined') {
    if (dark) {
      document.documentElement.classList.add('dark')
      localStorage.setItem('theme', 'dark')
    } else {
      document.documentElement.classList.remove('dark')
      localStorage.setItem('theme', 'light')
    }
  }
}

export function toggleTheme() {
  setTheme(!isDarkMode.value)
}

export function initTheme() {
  setTheme(getInitialTheme())
}

export function useTheme() {
  onMounted(() => {
    initTheme()
  })

  return {
    isDarkMode,
    toggleTheme,
    setTheme
  }
}

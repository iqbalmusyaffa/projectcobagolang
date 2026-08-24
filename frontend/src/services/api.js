import axios from 'axios'

// Instance Axios terintegrasi dengan REST API Golang
const api = axios.create({
  baseURL: 'http://localhost:8080/api',
  headers: {
    'Content-Type': 'application/json'
  }
})

// Request Interceptor: Otomatis menyelipkan Authorization Bearer Token dari localStorage
api.interceptors.request.use((config) => {
  const token = localStorage.getItem('token') || localStorage.getItem('access_token')
  if (token) {
    config.headers.Authorization = `Bearer ${token}`
  }
  return config
}, (error) => {
  return Promise.reject(error)
})

// Flag untuk melacak status refreshing agar tidak terjadi infinite loop
let isRefreshing = false
let failedQueue = []

const processQueue = (error, token = null) => {
  failedQueue.forEach(prom => {
    if (error) {
      prom.reject(error)
    } else {
      prom.resolve(token)
    }
  })
  failedQueue = []
}

// Response Interceptor: Menangani Token Expired (401 Unauthorized) dengan Auto-Refresh Token Mulus
api.interceptors.response.use((response) => {
  return response
}, async (error) => {
  const originalRequest = error.config

  // Jika error 401 dan request bukan berasal dari /login atau /refresh-token
  if (error.response && error.response.status === 401 && !originalRequest._retry) {
    if (originalRequest.url.includes('/login') || originalRequest.url.includes('/refresh-token')) {
      return Promise.reject(error)
    }

    if (isRefreshing) {
      return new Promise((resolve, reject) => {
        failedQueue.push({ resolve, reject })
      }).then(token => {
        originalRequest.headers.Authorization = `Bearer ${token}`
        return api(originalRequest)
      }).catch(err => {
        return Promise.reject(err)
      })
    }

    originalRequest._retry = true
    isRefreshing = true

    const refreshToken = localStorage.getItem('refresh_token')
    if (!refreshToken) {
      isRefreshing = false
      localStorage.removeItem('token')
      localStorage.removeItem('access_token')
      localStorage.removeItem('refresh_token')
      if (window.location.pathname !== '/login' && window.location.pathname !== '/register') {
        window.location.href = '/login'
      }
      return Promise.reject(error)
    }

    try {
      // Panggil endpoint POST /api/refresh-token secara transparan
      const res = await axios.post('http://localhost:8080/api/refresh-token', {
        refresh_token: refreshToken
      })

      const newAccessToken = res.data.data.access_token
      localStorage.setItem('token', newAccessToken)
      localStorage.setItem('access_token', newAccessToken)

      api.defaults.headers.common['Authorization'] = `Bearer ${newAccessToken}`
      originalRequest.headers['Authorization'] = `Bearer ${newAccessToken}`

      processQueue(null, newAccessToken)
      return api(originalRequest)
    } catch (refreshErr) {
      processQueue(refreshErr, null)
      localStorage.removeItem('token')
      localStorage.removeItem('access_token')
      localStorage.removeItem('refresh_token')
      if (window.location.pathname !== '/login' && window.location.pathname !== '/register') {
        window.location.href = '/login'
      }
      return Promise.reject(refreshErr)
    } finally {
      isRefreshing = false
    }
  }

  return Promise.reject(error)
})

export default api

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
  const token = localStorage.getItem('token')
  if (token) {
    config.headers.Authorization = `Bearer ${token}`
  }
  return config
}, (error) => {
  return Promise.reject(error)
})

// Response Interceptor: Menangani Token Expired (401 Unauthorized)
api.interceptors.response.use((response) => {
  return response
}, (error) => {
  if (error.response && error.response.status === 401) {
    localStorage.removeItem('token')
    if (window.location.pathname !== '/login' && window.location.pathname !== '/register') {
      window.location.href = '/login'
    }
  }
  return Promise.reject(error)
})

export default api

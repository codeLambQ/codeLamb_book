import axios from 'axios'

const http = axios.create({
  baseURL: '',
  withCredentials: true,
  timeout: 10000,
})

// 统一解包 data，并把后端 message 转成可读的错误
http.interceptors.response.use(
  (res) => res.data,
  (err) => {
    const message =
      err.response?.data?.message ||
      (err.code === 'ECONNABORTED' ? '请求超时，请稍后重试' : '网络错误，请确认后端已启动')
    return Promise.reject(new Error(message))
  },
)

export function login(email, password) {
  return http.post('/login', { email, password })
}

export function register(payload) {
  return http.post('/users', payload)
}

export function logout() {
  return http.post('/logout')
}

export function fetchMe() {
  return http.get('/me')
}

export default http

import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

// 开发环境：前端 5173 端口，通过代理转发到 Go 后端 8080，
// 这样请求同源，Cookie 自动携带，也无需处理跨域。
export default defineConfig({
  plugins: [vue()],
  server: {
    host: '127.0.0.1',
    port: 5173,
    proxy: {
      '/login': { target: 'http://127.0.0.1:8080', changeOrigin: true },
      '/logout': { target: 'http://127.0.0.1:8080', changeOrigin: true },
      '/users': { target: 'http://127.0.0.1:8080', changeOrigin: true },
      '/me': { target: 'http://127.0.0.1:8080', changeOrigin: true },
    },
  },
})

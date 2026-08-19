import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import { login, register, logout, fetchMe } from '../api'

const STORAGE_KEY = 'red_user'

function readUser() {
  try {
    return JSON.parse(localStorage.getItem(STORAGE_KEY) || 'null')
  } catch {
    return null
  }
}

export const useUserStore = defineStore('user', () => {
  const user = ref(readUser())
  const isAuthenticated = computed(() => !!user.value)

  function setUser(u) {
    user.value = u
    if (u) {
      localStorage.setItem(STORAGE_KEY, JSON.stringify(u))
    } else {
      localStorage.removeItem(STORAGE_KEY)
    }
  }

  async function loginAction(email, password) {
    const data = await login(email, password)
    setUser({ id: data.id, email: data.email })
    return data
  }

  function registerAction(payload) {
    return register(payload)
  }

  async function logoutAction() {
    try {
      await logout()
    } catch {
      // 登出接口异常时本地照常清除
    }
    setUser(null)
  }

  async function loadProfile() {
    const data = await fetchMe()
    setUser({ id: data.id, email: data.email })
    return data
  }

  return { user, isAuthenticated, setUser, loginAction, registerAction, logoutAction, loadProfile }
})

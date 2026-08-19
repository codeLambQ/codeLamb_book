<template>
  <div class="auth">
    <div class="auth-card">
      <div class="auth-aside">
        <div class="brand">
          <span class="brand-mark">书</span>
          <span>小红书</span>
        </div>
        <p class="slogan">标记我的生活</p>
        <p class="slogan-sub">分享读书与生活中的美好瞬间</p>
      </div>

      <div class="auth-main">
        <h2 class="title">欢迎登录</h2>
        <p class="subtitle">登录后开始记录你的生活</p>

        <div v-if="registered" class="success">注册成功，请使用新账号登录</div>

        <form @submit.prevent="onSubmit">
          <div class="field">
            <input
              v-model.trim="form.email"
              class="input"
              type="email"
              placeholder="邮箱"
              autocomplete="email"
            />
          </div>
          <div class="field">
            <input
              v-model="form.password"
              class="input"
              type="password"
              placeholder="密码"
              autocomplete="current-password"
            />
          </div>
          <p v-if="error" class="error">{{ error }}</p>
          <button class="btn btn-primary btn-block" type="submit" :disabled="loading">
            {{ loading ? '登录中...' : '登录' }}
          </button>
        </form>

        <p class="switch">
          还没有账号？
          <router-link to="/register" class="link">立即注册</router-link>
        </p>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useUserStore } from '../store/user'

const route = useRoute()
const router = useRouter()
const userStore = useUserStore()

const form = ref({ email: '', password: '' })
const error = ref('')
const loading = ref(false)
const registered = computed(() => route.query.registered === '1')

async function onSubmit() {
  error.value = ''
  if (!form.value.email || !form.value.password) {
    error.value = '请输入邮箱和密码'
    return
  }
  loading.value = true
  try {
    await userStore.loginAction(form.value.email, form.value.password)
    router.replace('/profile')
  } catch (e) {
    error.value = e.message || '登录失败'
  } finally {
    loading.value = false
  }
}
</script>

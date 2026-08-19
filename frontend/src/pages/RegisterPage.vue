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
        <h2 class="title">创建账号</h2>
        <p class="subtitle">加入我们，记录每一个瞬间</p>

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
              autocomplete="new-password"
            />
            <p class="hint">至少 8 位，需包含大写字母、小写字母和特殊字符</p>
          </div>
          <div class="field">
            <input
              v-model="form.confirm"
              class="input"
              type="password"
              placeholder="确认密码"
              autocomplete="new-password"
            />
          </div>
          <p v-if="error" class="error">{{ error }}</p>
          <button class="btn btn-primary btn-block" type="submit" :disabled="loading">
            {{ loading ? '注册中...' : '注册' }}
          </button>
        </form>

        <p class="switch">
          已有账号？
          <router-link to="/login" class="link">去登录</router-link>
        </p>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useUserStore } from '../store/user'

const router = useRouter()
const userStore = useUserStore()

const form = ref({ email: '', password: '', confirm: '' })
const error = ref('')
const loading = ref(false)

const EMAIL_RE = /^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/
const PASSWORD_RE = /^(?=.*[A-Z])(?=.*[a-z])(?=.*[!@#$%^&*()_+\-=[\]{};':"\\|,.<>/?]).{8,}$/

function validate() {
  if (!EMAIL_RE.test(form.value.email)) return '邮箱格式有误'
  if (form.value.password !== form.value.confirm) return '两次密码不一致'
  if (!PASSWORD_RE.test(form.value.password)) return '密码至少 8 位，需包含大写、小写字母和特殊字符'
  return ''
}

async function onSubmit() {
  error.value = ''
  const msg = validate()
  if (msg) {
    error.value = msg
    return
  }
  loading.value = true
  try {
    await userStore.registerAction({
      email: form.value.email,
      password: form.value.password,
      confirm_password: form.value.confirm,
    })
    router.push('/login?registered=1')
  } catch (e) {
    error.value = e.message || '注册失败'
  } finally {
    loading.value = false
  }
}
</script>

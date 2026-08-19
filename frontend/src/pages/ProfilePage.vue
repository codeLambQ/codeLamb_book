<template>
  <div class="profile">
    <header class="topbar">
      <router-link to="/profile" class="brand">
        <span class="brand-mark">书</span>
        <span>小红书</span>
      </router-link>
      <button class="btn btn-ghost logout" @click="onLogout">退出登录</button>
    </header>

    <div v-if="loading" class="loading">加载中...</div>

    <main v-else class="content">
      <section class="hero">
        <div class="hero-cover"></div>
        <div class="hero-body">
          <div class="avatar">{{ initial }}</div>
          <div class="info">
            <h1 class="name">{{ user.email }}</h1>
            <p class="red-id">小红书号：{{ user.id }}</p>
          </div>
        </div>
        <div class="stats">
          <div class="stat">
            <b>0</b>
            <span>关注</span>
          </div>
          <div class="stat">
            <b>0</b>
            <span>粉丝</span>
          </div>
          <div class="stat">
            <b>0</b>
            <span>获赞与收藏</span>
          </div>
        </div>
      </section>

      <section class="tabs">
        <button
          v-for="t in tabs"
          :key="t.key"
          class="tab"
          :class="{ active: activeTab === t.key }"
          @click="activeTab = t.key"
        >
          {{ t.label }}
        </button>
      </section>

      <section class="empty">
        <div class="empty-icon">📖</div>
        <p class="empty-title">{{ emptyTitle }}</p>
        <p class="empty-sub">功能开发中，敬请期待</p>
        <button class="btn btn-primary" @click="onPublish">发布笔记</button>
      </section>
    </main>

    <transition name="fade">
      <div v-if="toast" class="toast">{{ toast }}</div>
    </transition>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useUserStore } from '../store/user'

const router = useRouter()
const userStore = useUserStore()

const loading = ref(true)
const activeTab = ref('notes')
const toast = ref('')

const tabs = [
  { key: 'notes', label: '笔记' },
  { key: 'collect', label: '收藏' },
  { key: 'like', label: '赞过' },
]

const user = computed(() => userStore.user || {})
const initial = computed(() => (user.value.email || '书').charAt(0).toUpperCase())
const emptyTitle = computed(() => {
  const map = {
    notes: '还没有发布笔记',
    collect: '还没有收藏内容',
    like: '还没有点赞内容',
  }
  return map[activeTab.value] || '暂无内容'
})

onMounted(async () => {
  try {
    await userStore.loadProfile()
  } catch {
    userStore.setUser(null)
    router.replace('/login')
    return
  } finally {
    loading.value = false
  }
})

async function onLogout() {
  await userStore.logoutAction()
  router.replace('/login')
}

function onPublish() {
  toast.value = '笔记发布功能开发中'
  setTimeout(() => {
    toast.value = ''
  }, 2000)
}
</script>

<style scoped>
.profile {
  min-height: 100vh;
  background: var(--bg);
}

.topbar {
  position: sticky;
  top: 0;
  z-index: 10;
  height: 60px;
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 0 24px;
  background: #fff;
  border-bottom: 1px solid var(--border);
}

.logout {
  height: 36px;
  font-size: 14px;
  padding: 0 18px;
}

.loading {
  text-align: center;
  padding: 80px 0;
  color: var(--text-3);
}

.content {
  max-width: 760px;
  margin: 0 auto;
  padding: 24px 20px 60px;
}

.hero {
  background: #fff;
  border-radius: var(--radius);
  overflow: hidden;
  box-shadow: var(--shadow);
}

.hero-cover {
  height: 120px;
  background: linear-gradient(135deg, #ff2442, #ff6a80);
}

.hero-body {
  display: flex;
  align-items: flex-end;
  gap: 16px;
  padding: 0 24px;
  margin-top: -34px;
}

.avatar {
  width: 84px;
  height: 84px;
  border-radius: 50%;
  border: 3px solid #fff;
  display: flex;
  align-items: center;
  justify-content: center;
  background: linear-gradient(135deg, #ff2442, #ff7a8e);
  color: #fff;
  font-size: 36px;
  font-weight: 800;
}

.info {
  padding-bottom: 8px;
}

.name {
  font-size: 22px;
  font-weight: 800;
  word-break: break-all;
}

.red-id {
  margin-top: 4px;
  font-size: 13px;
  color: var(--text-3);
}

.stats {
  display: flex;
  padding: 20px 24px;
}

.stat {
  flex: 1;
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 2px;
}

.stat b {
  font-size: 20px;
  font-weight: 800;
}

.stat span {
  font-size: 13px;
  color: var(--text-3);
}

.tabs {
  display: flex;
  gap: 8px;
  margin-top: 16px;
  background: #fff;
  border-radius: var(--radius);
  padding: 0 16px;
  box-shadow: var(--shadow);
}

.tab {
  flex: 1;
  height: 48px;
  background: none;
  border: none;
  cursor: pointer;
  font-size: 15px;
  color: var(--text-2);
  font-weight: 600;
  border-bottom: 2px solid transparent;
}

.tab.active {
  color: var(--red);
  border-bottom-color: var(--red);
}

.empty {
  margin-top: 16px;
  background: #fff;
  border-radius: var(--radius);
  padding: 60px 24px;
  display: flex;
  flex-direction: column;
  align-items: center;
  box-shadow: var(--shadow);
}

.empty-icon {
  font-size: 48px;
}

.empty-title {
  margin-top: 16px;
  font-size: 16px;
  font-weight: 600;
}

.empty-sub {
  margin: 6px 0 24px;
  font-size: 13px;
  color: var(--text-3);
}

.toast {
  position: fixed;
  left: 50%;
  bottom: 40px;
  transform: translateX(-50%);
  background: rgba(0, 0, 0, 0.8);
  color: #fff;
  padding: 10px 22px;
  border-radius: 22px;
  font-size: 14px;
  z-index: 999;
}

.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.25s;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}
</style>

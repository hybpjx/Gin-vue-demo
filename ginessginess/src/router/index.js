import store from '@/store';
import { createRouter, createWebHistory, useRouter } from 'vue-router';
import HomeView from '../views/HomeView.vue';
import userRoutes from './module/user';



const routes = [
  {
    path: '/',
    name: 'home',
    component: HomeView
  },
  {
    path: '/about',
    name: 'about',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/AboutView.vue')
  },

  ...userRoutes

]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

router.beforeEach((to, from, next) => {
  // 判断是否需要登录
  if (to.meta.auth) {
    // 判断用户是否登录
    if (store.state.userModule.token) {
      // 这里还需要判断token的有效期 比如有没有过期，需要后台发放token，带上token的有效期
      // 如果token无效 需要请求token


      next()
    } else {
      // 跳转登录
      router.push({ name: 'login' })
    }
  } else {
    next()
  }
})

export default router

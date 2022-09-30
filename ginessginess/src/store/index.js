import { createStore } from 'vuex'
import userModule from './module/user'

// module的值
export default createStore({
  strict: process.env.NODE_ENV != 'production',
  state: {
  },
  getters: {
  },
  mutations: {
  },
  actions: {
  },
  modules: {
    // 通过模块去取vuex里的值
    userModule
  }
})

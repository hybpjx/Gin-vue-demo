import { createApp } from 'vue'
import Vuelidate from 'vuelidate'

import App from './App.vue'
import router from './router'
import store from './store'
import bulma from "bulma/css/bulma.css"
import axios from 'axios'
import VueAxios from 'vue-axios'
// 导入 element-ui
import elementPlus from 'element-plus';
import 'element-plus/dist/index.css'
const app = createApp(App)

app.use(store)
app.use(Vuelidate)
app.use(router)
// axios
app.use(VueAxios, axios)
app.config.globalProperties.$axios = axios
app.use(elementPlus)
app.mount('#app')


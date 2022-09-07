import { createApp, Vue } from 'vue';
import { BootstrapVue, IconsPlugin } from 'bootstrap-vue';
import App from './App.vue';
import router from './router';
import store from './store';

import '@/assets/scss/index.scss';

// Install BootstrapVue
Vue.use(BootstrapVue);
// Optionally install the BootstrapVue icon components plugin
Vue.use(IconsPlugin);

// eslint 1. 代码规范 2. 代码报错
createApp(App).use(store).use(router).mount('#app');

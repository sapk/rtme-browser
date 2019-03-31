import Vue from 'vue'
import AsyncComputed from 'vue-async-computed'

import App from './App.vue'

Vue.config.productionTip = false
Vue.use(AsyncComputed);

import axios from "axios";
axios.defaults.baseURL = process.env.VUE_APP_API_URL;

new Vue({
  render: h => h(App),
}).$mount('#app')

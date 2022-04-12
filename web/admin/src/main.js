import Vue from 'vue'
import App from './App.vue'
import router from './router'
import axios from 'axios'
import './plugin/antui'

import './assets/css/style.css'
import mavonEditor from 'mavon-editor'
import 'mavon-editor/dist/css/index.css'
Vue.use(mavonEditor)

axios.defaults.baseURL = "http://localhost:3000/api/v1";

axios.defaults.headers['Content-Type'] = 'application/json;charset=UTF-8'; //配置请求头的内容协商
// axios.defaults.headers['Authorization'] = 'Bearer ' + window.sessionStorage.getItem('token');    //配置请求头的token


// 添加拦截器
axios.interceptors.request.use(config => {
	config.headers.Authorization = `Bearer ${window.sessionStorage.getItem('token')}`;
	return config;
});

Vue.prototype.$axios = axios; //axios挂接到vue

Vue.config.productionTip = false;

new Vue({
	router,
	render: h => h(App)
}).$mount('#app')
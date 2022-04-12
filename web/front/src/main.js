import Vue from 'vue'
import App from './App.vue'
import router from './router'
import vuetify from './plugins/vuetify';

import axios from 'axios'
import moment from 'moment'


import './assets/css/prism-customer.css'
// import '../node_modules/prismjs/themes/prism.css'
// import '../node_modules/prismjs/themes/prism-okaidia.css'


Vue.filter('dateformat', function (inDate, outDate) {
	return moment(inDate).format(outDate);
})


// 设置axios基础url地址
axios.defaults.baseURL = "http://localhost:3000/api/v1";
// axios挂载到vue原型
Vue.prototype.$axios = axios



Vue.config.productionTip = false

new Vue({
	router,
	vuetify,
	render: h => h(App)
}).$mount('#app')

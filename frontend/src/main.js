import { createApp } from 'vue'
import App from './App.vue'


// 라우터 사용
// 1. router import
// 2. use(router) 등록
import router from './router'
createApp(App).use(router).mount('#app')

import { createApp } from 'vue'
import App from './App.vue'

// fontAwesome import
import { library } from '@fortawesome/fontawesome-svg-core'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'
import { faUserSecret } from '@fortawesome/free-solid-svg-icons'
library.add(faUserSecret)

// router import
import router from './router'



createApp(App)
  .use(router)  // router 등록
  .component('font-awesome-icon', FontAwesomeIcon)  // fontAwesome
  .mount('#app')

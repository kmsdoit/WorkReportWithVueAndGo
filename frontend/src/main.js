import { createApp } from 'vue'
import App from './App.vue'

// fontAwesome import
import { library } from '@fortawesome/fontawesome-svg-core'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'

// 이렇게하면 모든 아이콘들을 불러올 수 있다.
import { fas } from '@fortawesome/free-solid-svg-icons'
import { far } from '@fortawesome/free-regular-svg-icons'
import { fab } from '@fortawesome/free-brands-svg-icons'

// 불러온 아이콘을 라이브러리에 담는다.
library.add(
  fas,
  far,
  fab
)

// router import
import router from './router'


createApp(App)
  .use(router)  // router 등록
  .component('font-awesome-icon', FontAwesomeIcon)  // fontAwesome
  .mount('#app')

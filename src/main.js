import {createApp} from 'vue'
import './style.css'
import router from './router'
import App from './App.vue'

import naive from 'naive-ui'

createApp(App).use(naive).use(router).mount('#app')

import { registerPlugins } from '@/plugins'
import App from './App.vue'
import router from './router.js'
import { createApp } from 'vue'

import 'unfonts.css'

const app = createApp(App)

app.use(router)
registerPlugins(app)

app.mount('#app')

import { createApp } from 'vue'
import App from './App.vue'

const app = createApp(App)

app.config.productionTip = false

app.mount('#app')

console.log('Vue 3 app mounted successfully')

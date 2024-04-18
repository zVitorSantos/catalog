import { createApp } from 'vue';
import App from './App.vue';
import router from './router';

const app = createApp(App);

import './assets/css/main.css';

app.use(router);

app.mount('#app');

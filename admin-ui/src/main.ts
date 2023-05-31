import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import { i18n } from "./i18n";
import pinia from "./store";

// 引入jQuery、bootstrap
import $ from 'jquery'
import 'bootstrap'

// 引入bootstrap样式
import 'bootstrap/dist/css/bootstrap.min.css'
import 'bootstrap/dist/js/bootstrap.min.js'

// 引入图标库
import "@/assets/css/nucleo-icons.css"
import "@/assets/css/nucleo-svg.css"

import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'

import formatter from '@/utils/formatter';
const app = createApp(App)
// 全局注册 $
app.config.globalProperties.$ = $
app.config.globalProperties.$formatter = formatter;

app.use(i18n);
app.use(pinia);
app.use(router);
app.use(ElementPlus);
app.mount('#app')

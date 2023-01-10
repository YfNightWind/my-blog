import { createApp } from "vue";
import { createPinia } from "pinia";

import App from "./App.vue";
import router from "./router";
import axios from "axios";
import VueAxios from "vue-axios"

import "./assets/main.css";
import "./plugins/antui.ts";

axios.defaults.baseURL = "http://localhost:3000/api/v1";
axios.interceptors.request.use((config: any) => {
    config.headers.Authorization = `Bearer ${window.localStorage.getItem("token")}`
    return config;
})
VueAxios.prototype.http = axios;

const app = createApp(App);

app.use(createPinia());
app.use(router);
app.use(VueAxios, axios);
// 全局注入axios
app.provide('axios', app.config.globalProperties.axios)

app.mount("#app");

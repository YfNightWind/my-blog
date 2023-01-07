import { createApp } from "vue"
import { Button, Input, Form, message, Layout, Menu,  } from "ant-design-vue"
import "ant-design-vue/es/message/style/css";
import AppVue from "@/App.vue"

message.config({
    top: `60px`,
    duration: 2,
    maxCount: 3,
});

createApp(AppVue)
    .use(Button)
    .use(Form)
    .use(Input)
    .use(Layout)
    .use(Menu)
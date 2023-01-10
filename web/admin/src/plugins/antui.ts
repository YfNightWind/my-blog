import { createApp } from "vue"
import { Button, Input, Form, message, Layout, Menu, Card, Table, Row, Col, Modal, Select, Switch } from "ant-design-vue"
import "ant-design-vue/es/message/style/css";
import "ant-design-vue/es/modal/style/css";
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
    .use(Card)
    .use(Table)
    .use(Row)
    .use(Col)
    .use(Modal)
    .use(Select)
    .use(Switch)
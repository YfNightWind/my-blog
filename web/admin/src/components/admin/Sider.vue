<template>
  <a-layout-sider v-model:collapsed="collapsed" collapsible breakpoint="lg">
    <div class="logo">
      <span>{{ collapsed ? "Blog" : "My-Blog" }}</span>
    </div>
    <div>
      <a-menu
        v-model:openKeys="openKeys"
        v-model:selectedKeys="selectedKeys"
        mode="inline"
        theme="dark"
        :inline-collapsed="collapsed"
        @click="GoIndex"
      >
        <a-menu-item key="index">
          <template #icon>
            <DashboardOutlined />
          </template>
          <span>DashBoard</span>
        </a-menu-item>
        <a-sub-menu key="article">
          <template #icon>
            <FileOutlined />
          </template>
          <template #title>文章管理</template>
          <a-menu-item key="addarticle">
            <template #icon><EditOutlined /> </template><span>写文章</span>
          </a-menu-item>
          <a-menu-item key="articlelist">
            <template #icon><HddOutlined /> </template><span>文章列表</span>
          </a-menu-item>
        </a-sub-menu>
        <a-menu-item key="categorylist">
          <template #icon>
            <BookOutlined />
          </template>
          <span>分类列表</span>
        </a-menu-item>
        <a-menu-item key="userlist">
          <template #icon>
            <UserOutlined />
          </template>
          <span>用户列表</span>
        </a-menu-item>
      </a-menu>
    </div>
  </a-layout-sider>
</template>

<script lang="ts">
import { reactive, toRefs, watch } from "vue";
import {
  EditOutlined,
  FileOutlined,
  DashboardOutlined,
  BookOutlined,
  HddOutlined,
  UserOutlined,
} from "@ant-design/icons-vue";
import router from "@/router";
export default {
  components: {
    EditOutlined,
    HddOutlined,
    FileOutlined,
    DashboardOutlined,
    BookOutlined,
    UserOutlined,
  },
  setup() {
    const state = reactive({
      collapsed: false,
      selectedKeys: ["index"],
      openKeys: ["article"],
      preOpenKeys: ["article"],
    });

    watch(
      () => state.openKeys,
      (_val, oldVal) => {
        state.preOpenKeys = oldVal;
      }
    );
    const toggleCollapsed: any = () => {
      state.collapsed = !state.collapsed;
      state.openKeys = state.collapsed ? [] : state.preOpenKeys;
    };
    const GoIndex: any = (item: any) => {
      router.push(item.key);
    };

    return {
      ...toRefs(state),
      toggleCollapsed,
      GoIndex,
    };
  },
};
</script>

<style scoped>
.logo {
  height: 32px;
  margin: 16px;
  background-color: #ffffff;
  display: flex;
  justify-content: center;
  align-items: center;
  font-size: 18px;
}
</style>

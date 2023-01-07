<template>
  <a-layout class="container">
    <Sider></Sider>
    <a-layout>
      <a-layout-header class="header-button">
        <Header></Header>
      </a-layout-header>
      <a-layout-content>
        <router-view></router-view>
      </a-layout-content>
      <a-layout-footer>
        <Footer></Footer>
      </a-layout-footer>
    </a-layout>
  </a-layout>
</template>

<script lang="ts">
import Footer from "@/components/admin/Footer.vue";
import Sider from "@/components/admin/Sider.vue";
import Header from "@/components/admin/Header.vue";
export default {
  components: { Header, Sider, Footer },
  setup() {},
};
</script>

<style scoped>
.container {
  height: 100%;
  min-height: 100vh;
}

.header-button {
  display: flex;
  justify-content: flex-end;
  align-items: center;
}
</style>

<template>
  <a-button type="danger" @click="logout">退出</a-button>
</template>

<script lang="ts">
import router from "@/router";

export default {
  setup() {
    const logout = () => {
      window.localStorage.removeItem("token");
      router.push("/login");
    };
    return { logout };
  },
};
</script>

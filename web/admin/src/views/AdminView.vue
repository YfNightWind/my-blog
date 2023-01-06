<template>
  <div>管理页面</div>
  <a-button type="primary" @click="test">button here</a-button>
</template>

<script lang="ts">
import { message } from "ant-design-vue";
import { defineComponent } from "vue";

export default defineComponent({
  setup() {
    const test = () => {
      message.success("dada");
    };
    return { test };
  },
});
</script>

<template>
  <div class="container">
    <a-form
      ref="formRef"
      :model="formState"
      :rules="rules"
      name="normal_login"
      class="login-form"
      @finish="onFinish"
      @finishFailed="onFinishFailed"
    >
      <a-form-item label="用户名" name="username">
        <a-input v-model:value="formState.username">
          <template #prefix>
            <UserOutlined class="site-form-item-icon" />
          </template>
        </a-input>
      </a-form-item>

      <a-form-item label="密&nbsp;&nbsp;&nbsp;&nbsp;码" name="password">
        <a-input-password v-model:value="formState.password">
          <template #prefix>
            <LockOutlined class="site-form-item-icon" />
          </template>
        </a-input-password>
      </a-form-item>
      <a-form-item :wrapper-col="{ span: 14, offset: 4 }">
        <a-button
          :disabled="disabled"
          type="primary"
          html-type="submit"
          class="login-form-button"
        >
          登录
        </a-button>
        <a-button style="margin-left: 10px" @click="resetForm">Reset</a-button>
      </a-form-item>
    </a-form>
  </div>
</template>
<script lang="ts">
import { defineComponent, reactive, computed, ref } from "vue";
import { UserOutlined, LockOutlined } from "@ant-design/icons-vue";
import type { Rule } from "ant-design-vue/lib/form/interface";
import { message, type FormInstance } from "ant-design-vue";
import axios from "axios";
import router from "@/router";
interface FormState {
  username: string;
  password: string;
}
export default defineComponent({
  components: {
    UserOutlined,
    LockOutlined,
  },
  setup() {
    const formRef = ref<FormInstance>();
    const formState = reactive<FormState>({
      username: "",
      password: "",
    });
    const rules: Record<string, Rule[]> = {
      username: [
        { required: true, message: "请输入用户名!", trigger: "blur" },
        { min: 4, max: 15, message: "用户名必须在4-15位之间!" },
      ],
      password: [
        { required: true, message: "请输入密码!", trigger: "blur" },
        { min: 6, max: 20, message: "密码必须在6-20位之间!" },
      ],
    };
    const onFinish = async (values: any) => {
      var result: any = await axios.post("/login", formState);
      console.log(result);

      var statusCode: number = result.data.status;
      var msg: string = result.data.msg;
      var token: string = result.data.token;

      if (statusCode != 200) {
        return message.error(msg);
      } else {
        window.localStorage.setItem("token", token);
        router.push("/");
      }
    };

    const onFinishFailed = (errorInfo: any) => {
      message.error("输入非法数据，请检查后重试!");
    };
    const disabled = computed(() => {
      return !(formState.username && formState.password);
    });
    const resetForm = () => {
      formRef.value?.resetFields();
    };
    return {
      formRef,
      formState,
      rules,
      onFinish,
      onFinishFailed,
      disabled,
      resetForm,
    };
  },
});
</script>

<style scoped>
.container {
  height: 100%;
  background-color: #282c34;
}
.login-form {
  background-color: aliceblue;
  width: 450px;
  height: 300px;
  padding: 1%;
  position: absolute;
  top: 50%;
  left: 70%;
  transform: translate(-50%, -50%);
  border-radius: 15px;
  box-sizing: border-box;
}
/* .login-form-forgot {
  float: right;
} */
</style>

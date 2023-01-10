<template>
  <a-card>
    <a-row :gutter="16">
      <a-col :span="8">
        <a-input-search
          placeholder="输入要查找的用户"
          enter-button
          :v-model="queryData"
        />
      </a-col>
      <a-col>
        <a-button type="primary" @click="showModal">新增用户</a-button>
      </a-col>
    </a-row>
    <a-table
      bordered
      rowKey="id"
      :columns="columns"
      :data-source="dataSource"
      :pagination="pagination"
      :loading="loading"
      @change="handleTableChange"
    >
      <template #bodyCell="{ column, text, record }">
        <template v-if="column.key === 'role'">
          <span>{{ text == 1 ? "管理员" : "订阅者" }}</span>
        </template>
        <template v-if="column.key === 'action'">
          <a-button type="primary" style="margin-right: 15px">编辑</a-button>
          <a-button type="danger" @click="deleteUser(record.ID)">删除</a-button>
        </template>
      </template>
    </a-table>
    <!-- 新增用户 -->
    <a-modal
      v-model:visible="visible"
      title="新增用户"
      @ok="handleOk"
      :destrovOnClose="true"
    >
      <template #footer>
        <a-button key="back" @click="handleCancel">取消</a-button>
        <a-button
          key="submit"
          type="primary"
          :loading="userLoading"
          @click="handleOk"
          >提交</a-button
        >
      </template>
      <a-form
        :model="userInfo"
        :rules="rules"
        ref="addUserRef"
        v-model:value="userInfo"
      >
        <a-form-item label="用户名&nbsp;" name="username" has-feedback>
          <a-input v-model:value="userInfo.username"></a-input>
        </a-form-item>
        <a-form-item
          label="密&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;码"
          name="password"
          has-feedback
        >
          <a-input-password
            v-model:value="userInfo.password"
          ></a-input-password>
        </a-form-item>
        <a-form-item label="确认密码" name="confirmPassword" has-feedback>
          <a-input-password
            v-model:value="userInfo.confirmPassword"
          ></a-input-password>
        </a-form-item>
        <a-form-item label="是否为管理员" name="role">
          <a-select
            defaultValue="2"
            @change="adminChange"
            v-model:value="userInfo.role"
          >
            <a-select-option key="1" value="1">是</a-select-option>
            <a-select-option key="2" value="2">否</a-select-option>
          </a-select>
        </a-form-item>
      </a-form>
    </a-modal>
  </a-card>
</template>

<script lang="ts">
import axios from "axios";
import { usePagination } from "vue-request";
import { computed, createVNode, defineComponent, reactive, ref } from "vue";
import { message, Modal, type TableProps } from "ant-design-vue";
import { ExclamationCircleOutlined } from "@ant-design/icons-vue";
import ARow from "ant-design-vue/lib/grid/Row";
import pagination from "ant-design-vue/lib/pagination";
import type { Rule } from "ant-design-vue/lib/form/interface";

const columns = [
  {
    title: "ID",
    dataIndex: "ID",
    width: "20%",
    key: "id",
    align: "center",
  },
  {
    title: "用户名",
    dataIndex: "username",
    width: "20%",
    key: "username",
  },
  {
    title: "密码",
    dataIndex: "password",
    width: "20%",
    key: "password",
  },
  {
    title: "身份",
    dataIndex: "role",
    width: "20%",
    key: "role",
    align: "center",
  },
  {
    title: "操作",
    width: "20%",
    key: "action",
    align: "center",
  },
];

type APIParams = {
  username: string;
  pagenum: number;
  pagesize: number;
};

type APIResult = {
  data: {
    id: number;
    username: string;
    password: string;
  }[];
  total: number;
};

export default defineComponent({
  setup() {
    const queryData = async (params: APIParams) => {
      return await axios.get<APIResult>("users", { params });
    };

    const {
      data: dataSource,
      run,
      loading,
      current,
      pageSize,
    } = usePagination(queryData, {
      formatResult: (res) => res.data.data,
      pagination: {
        currentKey: "pagenum",
        pageSizeKey: "pagesize",
      },
    });

    // 获取用户总数
    const total = ref(10);
    async function updateTotal() {
      const temp = await axios.get("users");
      total.value = temp.data.total;
    }
    updateTotal();

    const pagination = computed(() => ({
      total: total.value,
      current: current.value,
      pageSize: pageSize.value,
    }));

    // @ts-ignore
    // TODO 显示总数量
    const handleTableChange: TableProps["onChange"] = (
      pag: { pageSize: number; current: number },
      filters: any,
      sorter: any
    ) => {
      run({
        pagenum: pag?.current!,
        pagesize: pag.pageSize!,
        sortField: sorter.field,
        sortOrder: sorter.order,
        ...filters,
      });
    };

    // 删除用户
    const deleteUser = (id: Number) => {
      Modal.confirm({
        title: "提示",
        icon: createVNode(ExclamationCircleOutlined),
        content: "确定要删除该用户吗?",
        okText: "Yes",
        okType: "danger",
        cancelText: "No",
        async onOk() {
          console.log("OK删除ID为 " + id);
          const res = await axios.delete(`user/${id}`);
          let statusCode: number = res.data.status;
          let msg: string = res.data.msg;
          if (statusCode != 200) {
            return message.error(msg);
          } else {
            location.reload();
            message.success(msg);
          }
        },
        onCancel() {
          message.info("已取消删除");
        },
      });
    };

    // 新增用户
    const addUserRef = ref();
    const userInfo = reactive({
      username: "",
      password: "",
      confirmPassword: "",
      role: 2,
    });
    let validatePass = async (_rule: Rule, value: string) => {
      if (value === "") {
        return Promise.reject("密码不能为空");
      } else {
        if (userInfo.password !== "") {
          addUserRef.value.validateFields("confirmPassowrd");
        }
        if (userInfo.password.length < 6 || userInfo.password.length > 20) {
          return Promise.reject("密码必须在6-20位之间!");
        }
        return Promise.resolve();
      }
    };
    let validatePass2 = async (_rule: Rule, value: string) => {
      if (value === "") {
        return Promise.reject("请再次输入密码!");
      } else if (value !== userInfo.password) {
        return Promise.reject("两次密码不一致!");
      } else {
        return Promise.resolve();
      }
    };
    const rules: Record<string, Rule[]> = {
      username: [
        { required: true, message: "请输入用户名!", trigger: "blur" },
        { min: 4, max: 15, message: "用户名必须在4-15位之间!" },
      ],
      password: [{ required: true, validator: validatePass, trigger: "blur" }],
      confirmPassword: [{ validator: validatePass2, trigger: "blur" }],
    };

    const userLoading = ref<boolean>(false);
    const visible = ref<boolean>(false);

    const showModal = () => {
      visible.value = true;
    };
    // 提交
    const handleOk = async (values: any) => {
      values = await axios.post("/user/add", {
        username: userInfo.username,
        password: userInfo.password,
        role: userInfo.role,
      });
      console.log(values);

      let statusCode: number = values.data.status;
      let msg: string = values.data.msg;

      if (statusCode != 200) {
        return message.error(msg);
      } else {
        visible.value = false;
        return message.success(msg);
      }
    };
    // 取消
    const handleCancel = () => {
      addUserRef.value.resetFields();
      visible.value = false;
    };

    const adminChange = (value: number) => {
      userInfo.role = Number(value);
    };

    return {
      queryData,
      dataSource,
      pagination,
      loading,
      total,
      updateTotal,
      columns,
      handleTableChange,
      deleteUser,
      userLoading,
      visible,
      showModal,
      handleOk,
      handleCancel,
      userInfo,
      rules,
      adminChange,
      addUserRef,
    };
  },
});
</script>

<style>
.ant-card {
  margin: 10px;
}

.ant-table {
  margin: 10px;
  height: 100%;
}
</style>

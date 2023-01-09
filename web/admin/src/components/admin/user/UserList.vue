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
        <a-button type="primary">新增用户</a-button>
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
  </a-card>
</template>

<script lang="ts">
import axios from "axios";
import { usePagination } from "vue-request";
import { computed, createVNode, defineComponent, ref } from "vue";
import { message, Modal, type TableProps } from "ant-design-vue";
import { ExclamationCircleOutlined } from "@ant-design/icons-vue";
import ARow from "ant-design-vue/lib/grid/Row";
import pagination from "ant-design-vue/lib/pagination";

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

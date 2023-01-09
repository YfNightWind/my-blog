<template>
  <a-card>
    <a-row :gutter="16">
      <a-col :span="8">
        <a-input-search placeholder="输入要查找的用户" enter-button />
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
      <template slot="action">
        <div class="actionSlot">
          <a-button type="primary">编辑</a-button>
          <a-button type="danger">编辑</a-button>
        </div>
      </template>
    </a-table>
  </a-card>
</template>

<script lang="ts">
import axios from "axios";
import { usePagination } from "vue-request";
import { computed, defineComponent, reactive, ref } from "vue";
import type { TableProps } from "ant-design-vue";
import { useCounterStore } from "@/stores/counter";
const columns = [
  {
    title: "ID",
    dataIndex: "ID",
    width: "20%",
    key: "id",
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
    title: "身份码",
    dataIndex: "role",
    width: "20%",
    key: "role",
  },
  {
    title: "操作",
    width: "20%",
    key: "action",
    scopedslots: { customender: "action" },
  },
];

type APIParams = {
  pagenum: number;
  pagesize: number;
};

type APIResult = {
  data: {
    id: number;
    username: string;
    password: string;
    total: number;
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

    return {
      dataSource,
      pagination,
      loading,
      total,
      updateTotal,
      columns,
      handleTableChange,
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

<template>
  <a-card>
    <a-row :gutter="16">
      <a-col :span="8">
        <a-input-search
          placeholder="输入要查找的文章"
          enter-button
          :v-model="queryData"
        />
      </a-col>
      <a-col :span="6" offset="8">
        <!-- <a-select>
          <a-select-option
            v-for="item in categoryList"
            v-model:key="item.id"
            v-model:value="item.name"
            >{{ item.name }}</a-select-option
          >
        </a-select> -->
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
      <template #bodyCell="{ column, record, text }">
        <template v-if="column.key === 'action'">
          <a-button
            type="primary"
            style="margin-right: 15px"
            @click=""
            size="small"
            >编辑</a-button
          >
          <a-button type="danger" @click="deleteArticle(record.ID)" size="small"
            >删除</a-button
          >
        </template>
        <template v-if="column.key === 'img'">
          <a-image :width="100" :src="record.img" />
        </template>
      </template>
    </a-table>
  </a-card>
</template>

<script lang="ts">
import axios from "axios";
import { usePagination } from "vue-request";
import {
  computed,
  createVNode,
  defineComponent,
  onMounted,
  reactive,
  ref,
  toRefs,
} from "vue";
import { message, Modal, type TableProps } from "ant-design-vue";
import { ExclamationCircleOutlined } from "@ant-design/icons-vue";
import ARow from "ant-design-vue/lib/grid/Row";
import pagination from "ant-design-vue/lib/pagination";

const columns = [
  {
    title: "ID",
    dataIndex: "ID",
    width: "10%",
    key: "id",
    align: "center",
  },
  {
    title: "文章标题",
    dataIndex: "title",
    width: "20%",
    key: "title",
  },
  {
    title: "分类名",
    dataIndex: ["Category", "name"],
    width: "10%",
    key: "name",
  },
  {
    title: "文章描述",
    dataIndex: "description",
    width: "20%",
    key: "description",
    align: "center",
  },
  {
    title: "缩略图",
    dataIndex: "img",
    width: "10%",
    key: "img",
    align: "center",
  },
  {
    title: "操作",
    width: "15%",
    key: "action",
    align: "center",
  },
];

type APIParams = {
  title: string;
  pagenum: number;
  pagesize: number;
};

type APIResult = {
  data: {
    id: number;
    title: string;
    name: string;
    description: string;
    img: string;
  }[];
  total: number;
};

export default defineComponent({
  setup() {
    const queryData = async (params: APIParams) => {
      return await axios.get<APIResult>("article", { params });
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

    // 获取文章总数
    const total = ref(10);
    async function updateTotal() {
      const temp = await axios.get("article");
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

    // 删除文章
    const deleteArticle = (id: Number) => {
      Modal.confirm({
        title: "提示",
        icon: createVNode(ExclamationCircleOutlined),
        content: "确定要删除该文章吗?",
        okText: "Yes",
        okType: "danger",
        cancelText: "No",
        async onOk() {
          console.log("OK删除ID为 " + id);
          const res = await axios.delete(`article/${id}`);
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

    // 获取分类
    const Data = reactive({
      categoryList: [],
    });
    // let categoryList: any = reactive([]);
    type CategoryParams = {
      pagenum: number;
      pagesize: number;
    };

    type CategoryResult = {
      data: {
        id: number;
        name: string;
      }[];
      total: number;
    };

    const getCategory = async (params: CategoryParams) => {
      const res = await axios.get<CategoryResult>("category", { params });
      //   categoryList.push(...res.data.data);
      console.log(res.data.data);
      
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
      deleteArticle,
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

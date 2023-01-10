<template>
  <a-card>
    <a-row>
      <a-col>
        <a-button type="primary" @click="showModal">新增分类</a-button>
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
      <template #bodyCell="{ column, record }">
        <template v-if="column.key === 'action'">
          <a-button
            type="primary"
            style="margin-right: 15px"
            @click="editCategory(record.id)"
            >编辑</a-button
          >
          <a-button type="danger" @click="deleteCategory(record.id)"
            >删除</a-button
          >
        </template>
      </template>
    </a-table>
    <!-- 新增分类 -->
    <a-modal
      v-model:visible="addCategoryVisible"
      title="新增分类"
      @ok="handleAddOk"
      :destrovOnClose="true"
    >
      <template #footer>
        <a-button key="back" @click="handleAddCancel">取消</a-button>
        <a-button
          key="submit"
          type="primary"
          :loading="categoryLoading"
          @click="handleAddOk"
          >提交</a-button
        >
      </template>
      <a-form
        :model="categoryInfo"
        :rules="rules"
        ref="addCategoryRef"
        v-model:value="categoryInfo"
      >
        <a-form-item label="分类名&nbsp;" name="name" has-feedback>
          <a-input v-model:value="categoryInfo.name"></a-input>
        </a-form-item>
      </a-form>
    </a-modal>
    <!-- 编辑分类 -->
    <a-modal
      v-model:visible="editCategoryVisible"
      title="编辑分类"
      @ok="handleEditOk"
    >
      <template #footer>
        <a-button key="back" @click="handleEditCancel">取消</a-button>
        <a-button
          key="submit"
          type="primary"
          :loading="categoryLoading"
          @click="handleEditOk"
          >提交</a-button
        >
      </template>
      <a-form
        :model="categoryInfo"
        :rules="rules"
        ref="editCategoryRef"
        v-model:value="categoryInfo"
      >
        <a-form-item label="分类名&nbsp;" name="name" has-feedback>
          <a-input v-model:value="categoryInfo.name"></a-input>
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
    dataIndex: "id",
    width: "20%",
    key: "id",
    align: "center",
  },
  {
    title: "分类名",
    dataIndex: "name",
    width: "20%",
    key: "name",
  },
  {
    title: "操作",
    width: "20%",
    key: "action",
    align: "center",
  },
];

type APIParams = {
  pagenum: number;
  pagesize: number;
};

type APIResult = {
  data: {
    id: number;
    name: string;
  }[];
  total: number;
};

export default defineComponent({
  setup() {
    const queryData = async (params: APIParams) => {
      return await axios.get<APIResult>("category", { params });
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
      const temp = await axios.get("category");
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

    // 删除分类
    const deleteCategory = (id: Number) => {
      Modal.confirm({
        title: "提示",
        icon: createVNode(ExclamationCircleOutlined),
        content: "确定要删除该分类吗?",
        okText: "Yes",
        okType: "danger",
        cancelText: "No",
        async onOk() {
          console.log("OK删除ID为 " + id);
          const res = await axios.delete(`/category/${id}`);
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

    //
    // 新增分类
    //
    const addCategoryRef = ref();
    const checked = ref<boolean>(false);
    const categoryInfo = reactive({
      name: "",
      id: 1,
    });

    const rules: Record<string, Rule[]> = {
      name: [{ required: true, message: "请输入分类名!", trigger: "blur" }],
    };

    const categoryLoading = ref<boolean>(false);
    const addCategoryVisible = ref<boolean>(false);

    const showModal = () => {
      addCategoryVisible.value = true;
    };
    // 提交
    const handleAddOk = async (values: any) => {
      values = await axios.post("/category/add", {
        name: categoryInfo.name,
      });
      console.log(values);

      let statusCode: number = values.data.status;
      let msg: string = values.data.msg;

      if (statusCode != 200) {
        return message.error(msg);
      } else {
        addCategoryVisible.value = false;
        return message.success(msg);
      }
    };
    // 取消
    const handleAddCancel = () => {
      addCategoryRef.value.resetFields();
      addCategoryVisible.value = false;
    };

    //
    // 编辑分类
    //
    const editCategoryVisible = ref<boolean>(false);

    // 点击编辑按钮
    const editCategory = async (id: number) => {
      editCategoryVisible.value = true;
      const res = await axios.get(`category/${id}`);
      console.log(res.data);
      categoryInfo.id = id;
      categoryInfo.name = res.data.data["name"];
    };
    // 确定
    const handleEditOk = async () => {
      let values = await axios.put(`/category/${categoryInfo.id}`, {
        name: categoryInfo.name,
      });
      console.log(values);

      let statusCode: number = values.data.status;
      let msg: string = values.data.msg;

      if (statusCode != 200) {
        return message.error(msg);
      } else {
        editCategoryVisible.value = false;
        return message.success(msg);
      }
    };
    // 取消
    const handleEditCancel = () => {
      editCategoryVisible.value = false;
      message.info("编辑该分类已取消");
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
      deleteCategory,
      checked,
      categoryLoading,
      addCategoryVisible,
      showModal,
      handleAddOk,
      handleAddCancel,
      categoryInfo,
      rules,
      addCategoryRef,
      editCategoryVisible,
      editCategory,
      handleEditOk,
      handleEditCancel,
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

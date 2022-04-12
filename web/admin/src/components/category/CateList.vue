<!-- 分类列表页面 -->
<template>
  <div>
    <!-- <h4 style="margin-left: 10px;color: #d6a20f;">分类列表页面</h4> -->
    <a-card>
      <a-row :gutter="20">
        <!-- 搜索框 -->
        <!-- <a-col :span="6">
                    <a-input-search v-model="queryParam.name" placeholder="请输入分类名查找" enter-button allowClear
                        @search="getCateList" />
                </a-col> -->

        <!-- 添加按钮 -->
        <a-col :span="4">
          <a-button type="primary" @click="addCate()">添加</a-button>
        </a-col>
      </a-row>

      <!-- 分类信息显示表格 -->
      <a-table :columns="columns" rowKey="id" :pagination="pagination" :data-source="cateList" @change="handleTableChange" bordered>

        <!-- 操作按钮 -->
        <template slot="action" slot-scope="data">
          <div class="actionSlot">
            <a-button type="primary" icon="edit" style="margin: 0 5px;" @click="editCate(data.id)">编辑
            </a-button>
            <a-button type="danger" icon="delete" style="margin: 0 5px;" @click="deleteCate(data.id)">删除
            </a-button>

          </div>
        </template>
      </a-table>

      <!-- 新增分类区域 -->
      <a-modal closable title="新增分类" :visible="addCateVisible" @ok="addCateOk" @cancel="addCateCancel">
        <a-form-model :model="newCateInfo" :rules="addCateRules" ref="addCateRef">
          <a-form-model-item label="分类名" prop="name">
            <a-input v-model="newCateInfo.name"></a-input>
          </a-form-model-item>

        </a-form-model>
      </a-modal>

      <!-- 编辑分类区域 -->
      <a-modal closable title="编辑分类" :visible="editCateVisible" @ok="editCateOk" @cancel="editCateCancel">
        <a-form-model :model="cateInfo" :rules="cateRules" ref="editCateRef">
          <a-form-model-item label="分类名" prop="name">
            <a-input v-model="cateInfo.name"></a-input>
          </a-form-model-item>

        </a-form-model>
      </a-modal>
    </a-card>
  </div>
</template>

<script>
// 表格列配置
const columns = [
  {
    title: 'ID',
    dataIndex: 'id',
    width: '10%',
    key: 'id',
    align: 'center',
  },
  {
    title: '分类名',
    dataIndex: 'name',
    width: '20%',
    key: 'name',
    align: 'center',
  },
  {
    title: '操作',
    width: '30%',
    key: 'action',
    scopedSlots: { customRender: 'action' },
    align: 'center',
  },
]

export default {
  data () {
    return {
      // 分页配置
      pagination: {
        pageSizeOptions: ['5', '10', '20'],
        pageSize: 5,
        total: 0,
        showSizeChanger: true,
        showTotal: (total) => `共${total}条`,
      },
      // 分类列表数据
      cateList: [],
      // 表格列信息（不可少）
      columns,
      // 分页查询参数
      queryParam: {
        name: '',  // 分类搜索文本
        pagesize: 5,
        pagenum: 1,
      },
      visible: false,

      addCateVisible: false,  //新增分类对话框是否显示
      editCateVisible: false, //编辑分类对话框是否显示
      //分类信息(新增分类使用)
      newCateInfo: {
        id: 0,
        name: '',
      },
      //分类信息（更新分类使用）
      cateInfo: {
        id: 0,
        name: '',
      },

      // 新增分类的表单验证
      addCateRules: {
        // 分类名校验规则
        name: [
          { required: true, message: '请输入分类名', trigger: 'blur' },
          { min: 2, max: 50, message: '分类名长度在2到50个字符之间', trigger: 'blur' },
        ],
      },

      // 更新分类的表单验证
      cateRules: {
        // 分类名校验规则
        name: [
          { required: true, message: '请输入分类名', trigger: 'blur' },
          { min: 2, max: 50, message: '分类名长度在2到50个字符之间', trigger: 'blur' },
        ],
      },
    }
  },
  //生命周期 - 创建完成（访问当前this实例）
  created () {
    this.getCateList();
  },
  //生命周期 - 挂载完成（访问DOM元素）
  mounted () {

  },
  //方法集
  methods: {
    // 获取分类列表
    async getCateList () {
      const { data: res } = await this.$axios.get('categories', {
        params: {
          name: this.queryParam.name,
          pagesize: this.queryParam.pagesize,
          pagenum: this.queryParam.pagenum,
        },
      });
      if (res.status != 200) {
        return this.$message.error(res.message);
      }
      this.cateList = res.data;
      this.pagination.total = res.total;
    },

    // 处理表格分页改变事件
    // eslint-disable-next-line no-unused-vars
    handleTableChange (pagination, filters, sorter) {
      var pager = { ...this.pagination }
      pager.current = pagination.current;
      pager.pageSize = pagination.pageSize;
      this.queryParam.pagesize = pagination.pageSize;
      this.queryParam.pagenum = pagination.current;

      if (pagination.pageSize !== this.pagination.pageSize) {
        this.queryParam.pagenum = 1;
        pager.current = 1;
      }
      this.pagination = pager;
      this.getCateList();
    },

    // 删除分类
    deleteCate (id) {
      this.$confirm({
        title: '提示：请再次确认',
        content: '确定删除此分类吗?一旦删除，不可恢复。',
        onOk: async () => {
          const res = await this.$axios.delete(`category/${id}`);
          console.log(res);
          if (res.status != 200) {
            return this.$message.error(res.message);
          }
          this.$message.success('删除成功');
          this.getCateList();
        },
        onCancel: () => {
          this.$message.info('操作已取消');
        },
      });

    },

    // 显示新增分类对话框
    addCate () {
      this.addCateVisible = true;
    },

    // 新增分类
    addCateOk () {
      this.$refs.addCateRef.validate(async (valid) => {
        if (!valid) return this.$message.error("参数不符合要求请重新输入");
        const { data: res } = await this.$axios.post("category", {
          name: this.newCateInfo.name,
        });
        if (res.status != 200) {
          return this.$message.error(res.message);
        }
        this.$refs.addCateRef.resetFields();
        this.addCateVisible = false;
        this.$message.success('新增分类成功');
        this.getCateList();
      });
    },
    // 取消新增分类
    addCateCancel () {
      this.$refs.addCateRef.resetFields();
      this.addCateVisible = false;
      this.$message.info('新增分类取消');
    },

    // 显示编辑分类对话框
    async editCate (id) {
      console.log(id);
      this.editCateVisible = true;
      const { data: res } = await this.$axios.get(`category/${id}`);
      this.cateInfo = res.data;
      this.cateInfo.id = id;
    },
    // 编辑分类确认
    editCateOk () {
      this.$refs.editCateRef.validate(async (valid) => {
        if (!valid) return this.$message.error("参数不符合要求请重新输入");

        const { data: res } = await this.$axios.put(`category/${this.cateInfo.id}`, {
          name: this.cateInfo.name,
        });
        if (res.status != 200) {
          return this.$message.error(res.message);
        }
        this.$refs.editCateRef.resetFields();
        this.editCateVisible = false;
        this.$message.success('编辑分类信息成功');
        this.getCateList();
      });
    },
    // 编辑分类取消
    editCateCancel () {
      this.$refs.editCateRef.resetFields();
      this.editCateVisible = false;
      this.$message.info('编辑取消');
    },
  }
}
</script>
<style scoped>
.actionSlot {
  display: flex;
  justify-content: center;
}
</style>
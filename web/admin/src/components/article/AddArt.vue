<!-- 添加文章页面 -->
<template>
  <div>
    <a-card>
      <h4 style="margin-left: 10px;color: #d6a20f;text-align: center;">{{id? '编辑文章':'写文章'}}</h4>
      <a-form-model :model="artInfo" ref="artInfoRef" :rules="artInfoRules" :hideRequiredMark="true">
        <!-- 标题 -->
        <a-form-model-item label="文章标题" prop="title">
          <a-input v-model="artInfo.title"></a-input>
        </a-form-model-item>
        <!-- 分类 -->
        <a-form-model-item label="文章分类" prop="cid">
          <a-select v-model="artInfo.cid" @change="cateChange" placeholder="请选择分类">
            <a-select-option v-for="cate in cateList" :key="cate.id" :value="cate.id">
              {{cate.name}}
            </a-select-option>
          </a-select>
        </a-form-model-item>
        <!-- 原创标记 -->
        <a-form-model-item label="原创标记" prop="flag">
          <a-input v-model="artInfo.flag"></a-input>
        </a-form-model-item>
        <!-- 描述 -->
        <a-form-model-item label="文章描述" prop="desc">
          <a-textarea v-model="artInfo.desc"></a-textarea>
        </a-form-model-item>
        <!-- 文章内容 -->
        <a-form-model-item label="文章内容" prop="content">
          <div style="height: 600px;">
            <!-- Markdown编辑器 -->
            <mavon-editor v-model="artInfo.content" :toolbars="toolbars" :codeStyle="codeStyle" ref=artContentMd @imgAdd="handleImgAdd" @imgDel="handleImgDel" placeholder="编辑文章内容..." style="height:100%;z-index: 999;font-size: 16px;" />
          </div>

        </a-form-model-item>
        <!-- 提交 -->
        <a-form-model-item>
          <a-button type="primary" style="margin: 0 10px 0 0 ;" @click="artSubmit(artInfo.id)">
            提交
          </a-button>
          <a-button type="danger" style="margin: 0 10px;" @click="addCancel()">清空</a-button>
        </a-form-model-item>
      </a-form-model>
    </a-card>
  </div>
</template>

<script>
export default {
  props: ['id'],
  data () {
    return {
      artInfo: {
        id: 0,
        title: '',
        cid: undefined,
        flag: '',
        desc: '',
        content: '',
      },
      cateList: [],

      // 表单验证
      artInfoRules: {
        title: [
          { required: true, message: "请输入文章标题", trigger: "blur" },
        ],
        cid: [
          { required: true, message: "请选择文章分类", trigger: "change" },
        ],
        flag: [
          { required: true, message: "请输入文章标记", trigger: "blur" },
        ],
        desc: [
          { required: true, message: "请选择文章描述", trigger: "blur" },
          { max: 120, message: "文章描述最多20个字符", trigger: "change" },
        ],
        content: [
          { required: true, message: "请输入文章内容", trigger: "blur" },
        ],
      },

      // Markdown工具栏配置
      toolbars: {
        bold: true, // 粗体
        italic: true, // 斜体
        header: true, // 标题
        underline: true, // 下划线
        mark: true, // 标记
        superscript: true, // 上角标
        quote: true, // 引用
        ol: true, // 有序列表
        link: true, // 链接
        imagelink: true, // 图片链接
        help: true, // 帮助
        code: true, // code
        subfield: true, // 是否需要分栏
        fullscreen: true, // 全屏编辑
        readmodel: true, // 沉浸式阅读
        undo: true, // 上一步
        trash: true, // 清空
        save: true, // 保存（触发events中的save事件）
        navigation: true, // 导航目录
        alignleft: false, // 左对齐
        aligncenter: false, // 居中
        alignright: false, // 右对齐

        preview: false, // 预览
      },
      // markdown编辑器代码样式 monokai-sublime
      codeStyle: "atom-one-light",//主题

    }
  },
  //生命周期 - 创建完成（访问当前this实例）
  created () {
    this.getCateList();
    if (this.id) {
      this.getArtInfo(this.id);
    }
  },
  //生命周期 - 挂载完成（访问DOM元素）
  mounted () {

  },
  //方法集
  methods: {
    //查询文章信息
    async getArtInfo (id) {
      const { data: res } = await this.$axios.get(`/article/${id}`);
      if (res.status != 200) {
        return this.$message.error(res.message);
      }
      this.artInfo = res.data;
      this.artInfo.id = res.data.ID;
    },

    // 获取分类列表
    async getCateList () {
      const { data: res } = await this.$axios.get('categories');
      if (res.status != 200) {
        return this.$message.error(res.message);
      }
      this.cateList = res.data;
    },

    //分类选择
    cateChange (val) {
      this.artInfo.cid = val;
    },

    //提交(更新)文章
    artSubmit (id) {
      // eslint-disable-next-line no-unused-vars
      this.$refs.artInfoRef.validate(async (valid) => {
        if (id == 0) {  //添加文章
          const { data: res } = await this.$axios.post("/article", this.artInfo);
          if (res.status != 200) {
            return this.$message.error(res.message);
          }
          this.$router.push("/artlist");
          this.$message.success("添加文章成功");
        } else {  //更新文章
          const { data: res } = await this.$axios.put(`/article/${id}`, this.artInfo);
          if (res.status != 200) {
            return this.$message.error(res.message);
          }
          this.$router.push("/artlist");
          this.$message.success("更新文章成功");
        }
      });
    },

    // 取消
    addCancel () {
      this.$refs.artInfoRef.resetFields();
    },

    // Markdown编辑器内容改变
    // change(value, render) {
    //     this.html = render;
    //     this.artInfo.content = value;
    //     this.artInfo.content = render;
    // },

    // 处理Markdown编辑器图片上传
    handleImgAdd (pos, $file) {
      // 第一步.将图片上传到服务器
      var formdata = new FormData();
      formdata.append('file', $file);
      this.$axios.post("/upload", formdata, {
        headers: {
          'Content-Type': 'multipart/form-data',
        }
      }).then(res => {
        //console.log("url:" , res.data.fileUrl);
        var url = res.data.fileUrl;
        // 第二步.将返回的url替换到文本原位置![...](0) -> ![...](url)
        this.$refs.artContentMd.$img2Url(pos, url);
      });
    },

    // 处理Markdown编辑器图片删除
    handleImgDel () {
      console.log("处理图片删除");
    },
  }
}
</script>
<style>
</style>
<template>
  <a-layout-sider breakpoint="lg" v-model="collapsed">
    <div class="logo">
      <div v-if="collapsed">
        <img src="https://avatars.githubusercontent.com/u/73086282?s=400&u=872c71b735992fc404eb2efeefe27c26b223fbdf&v=4" alt="" style="width: 48px;height: 48px;border-radius: 50%;">
      </div>
      <div v-else>
        <span>{{collapsed ? "听雨" : "个人博客后台管理系统"}}</span>
      </div>
    </div>

    <a-menu theme="dark" mode="inline" @click="goToPage">
      <a-menu-item key="index">
        <a-icon type="dashboard" /><span>仪表盘</span>
      </a-menu-item>
      <!-- 文章管理 -->
      <a-sub-menu>
        <span slot="title">
          <a-icon type="file" /><span>文章管理</span>
        </span>
        <a-menu-item key="addart">
          <a-icon type="form" /><span>写文章</span>
        </a-menu-item>
        <a-menu-item key="artlist">
          <a-icon type="snippets" /><span>文章列表</span>
        </a-menu-item>
      </a-sub-menu>
      <!-- 分类管理 -->
      <a-menu-item key="catelist">
        <a-icon type="unordered-list" /><span>分类列表</span>
      </a-menu-item>
      <!-- <a-menu-item key="commentlist">
        <a-icon type="message" />
        <span>评论管理</span>
      </a-menu-item> -->
      <!-- 用户列表 -->
      <a-menu-item key="userlist">
        <a-icon type="contacts" /><span>用户列表</span>
      </a-menu-item>
      <!-- 用户信息 -->
      <a-menu-item key="profile">
        <a-icon type="setting" /><span>个人设置</span>
      </a-menu-item>

    </a-menu>

  </a-layout-sider>
</template>

<script>
export default {
  data () {
    return {
      collapsed: false,
    }
  },
  created () {

  },
  mounted () {

  },
  //方法集
  methods: {
    // 跳转到对应菜单
    goToPage (item) {
      this.$router.push('/' + item.key).catch((err) => err);
    }
  }
}
</script>

<style scoped>
.logo {
  height: 48px;
  margin: 16px;
  /* background-color: white; */
  color: #ffd04b;
  display: flex;
  justify-content: center;
  align-items: center;
  font-size: 16px;
}
</style>
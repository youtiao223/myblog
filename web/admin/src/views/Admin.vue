<!-- 后台管理页面 -->
<template>
    <a-layout class="container">
        <!-- 导航 -->
        <Nav></Nav>
        
        <a-layout>
            <!-- 头部 -->
          <a-layout-header class="header-btn">
              <Header></Header>
          </a-layout-header>
          <!-- 中间内容 -->
          <a-layout-content>
              <router-view></router-view>
          </a-layout-content>
          <!-- 底部 -->
          <a-layout-footer>
              <Footer></Footer>
          </a-layout-footer>
        </a-layout>
      </a-layout>
</template>

<script>
    import Nav from '../components/admin/Nav.vue'
    import Footer from '../components/admin/Footer.vue'
    import Header from '../components/admin/Header.vue'

    export default {
        components:{ Nav,Footer,Header },
        data() {
            return {

            }
        },
        //生命周期 - 创建完成（访问当前this实例）
        created() {

        },
        //生命周期 - 挂载完成（访问DOM元素）
        mounted() {

        },
        //方法集
        methods: {
            
        }
    }
</script>

<style scoped>
    .container{
        height: 100%;
    }
    .ant-layout-footer {
        padding: 0 !important;
    }
    

    .header-btn{
        display: flex;
        justify-content: flex-end;
        align-items: center;
    }


</style>
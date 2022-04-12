import Vue from 'vue'
import VueRouter from 'vue-router'
import Login from '../views/Login.vue'
import Admin from '../views/Admin.vue'

import Index from '../components/admin/Index.vue'
import AddArt from '../components/article/AddArt.vue'
import ArtList from '../components/article/ArtList.vue'
import CateList from '../components/category/CateList.vue'
import UserList from '../components/user/UserList.vue'
import Profile from '../components/user/Profile'
Vue.use(VueRouter)

const routes = [{
  path: '/login',
  name: 'Login',
  component: Login,
  meta: {
    title: '后台登录',
  }
}, {
  path: '/',
  name: 'Admin',
  component: Admin,
  children: [{
      path: 'index',
      component: Index,
      meta: {
        title: '后台主页',
      }
    },
    {
      path: 'addart',
      component: AddArt,
      meta: {
        title: '写文章',
      }
    },
    {
      path: 'addart/:id',
      component: AddArt,
      props: true,
      meta: {
        title: '编辑文章',
      }
    },
    {
      path: 'artlist',
      component: ArtList,
      meta: {
        title: '文章列表页',
      }
    },
    {
      path: 'catelist',
      component: CateList,
      meta: {
        title: '分类列表页',
      }
    },
    {
      path: 'userlist',
      component: UserList,
      meta: {
        title: '用户列表页',
      }
    },
    {
      path: 'profile',
      component: Profile,
      meta: {
        title: '个人设置',
      }
    },

  ]
}]

const router = new VueRouter({
  routes
})
// 路由前置守卫
router.beforeEach((to, from, next) => {
  const token = window.sessionStorage.getItem("token");

  // 设置网页标题
  if (to.meta.title) {
    document.title = to.meta.title
  }

  if (to.path == "/login") {
    return next();
  }
  if (!token) {
    next("/login");
  } else {
    next();
  }
});
export default router
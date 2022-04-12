import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '@/views/Home.vue'
import ArticleList from '@/components/ArticleList'
import Detail from '@/components/Detail'
import CateList from '@/components/CateList'
import Search from '@/components/Search'

Vue.use(VueRouter)


//处理顶部按钮点击跳到相同视图而报错--------------------------------
//获取原型对象上的push函数
const originalPush = VueRouter.prototype.push
//修改原型对象中的push方法
VueRouter.prototype.push = function push(location) {
	return originalPush.call(this, location).catch(err => err)
}
//---------------------------------------------------------------


const routes = [
	{
		path: '/',
		name: 'Home',
		component: Home,
		children: [
			{
				path: '/',
				component: ArticleList,
				meta: {
					title: '首页 - 油虾条的个人博客'
				}
			},
			{
				path: 'article/:id',
				component: Detail,
				meta: {
					title: '文章详情'
				},
				props: true
			},
			{
				path: 'category/:cid',
				component: CateList,
				meta: {
					title: '分类信息'
				},
				props: true
			},
			{
				path: 'search/:keywords',
				component: Search,
				meta: {
					title: '搜索结果'
				},
				props: true
			},
		]
	},

]

const router = new VueRouter({
	routes
})


// 处理网页标题
router.beforeEach((to, from, next) => {
	if (to.meta.title) {
		document.title = to.meta.title ? to.meta.title : '加载中';
		next();
	}
});

export default router

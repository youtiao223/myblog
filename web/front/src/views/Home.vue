<template>
	<v-app app>
		<TopBar></TopBar>

		<v-main class="grey lighten-3">
			<v-container>
				<v-row>
					<v-col cols="3">
						<Nav></Nav>
					</v-col>
					<v-col>
						<v-sheet max-width="56vw" min-height="80vh" rounded="md">
							<router-view :key="key"></router-view>
						</v-sheet>
					</v-col>
				</v-row>
			</v-container>
		</v-main>

		<Footer></Footer>
	</v-app>
</template>

<script>
	import TopBar from '@/components/TopBar'
	import Footer from '@/components/Footer'
	import Nav from '@/components/Nav'
	export default {
		components: { TopBar, Footer, Nav },
		computed: {
			// 计算属性为router-view加上随机数key
			// 处理路由跳转而页面不刷新
			key() {
				return this.$route.path + Math.random();
			}
		}
	}
</script>
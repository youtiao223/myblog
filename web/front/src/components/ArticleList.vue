<template>
	<v-card class="pa-2 mb-2">
		<v-card-title>最新文章</v-card-title>

		<v-card elevation="3" class="article-info ma-4" v-for="article in articleList" :key="article.ID">
			<v-card-title>
				<router-link :to="'/article/'+ article.ID">{{article.title}}</router-link>
			</v-card-title>
			<v-card-text>
				<div class="mb-2">{{article.desc}}</div>
				<div class="mb-2">
					<v-chip color="orange" x-small outlined class="mr-2">
						{{article.flag}}
					</v-chip>
					<v-chip color="green" x-small outlined>
						{{article.Category.name}}
					</v-chip>
				</div>
				<div>
					<v-icon>{{'mdi-calendar-month'}}</v-icon>
					<span>{{article.CreatedAt | dateformat('YYYY-MM-DD HH:SS')}}</span>

				</div>
			</v-card-text>
		</v-card>
		<!-- 分页 -->
		<div class="text-center">
			<v-row justify="center">
				<v-col cols="8">
					<v-container class="max-width">
						<v-pagination circle :total-visible="7" v-model="queryParam.pageNum"
							:length="Math.ceil(this.total/queryParam.pageSize)" class="my-2" @input="getArticleList()"
							@next="getArticleList()" @previous="getArticleList()">
						</v-pagination>
					</v-container>
				</v-col>
			</v-row>
		</div>

	</v-card>
	<!-- <v-col>
	</v-col> -->
</template>
<script>
	export default {
		data() {
			return {
				articleList: [],
				queryParam: {
					pageSize: 5,
					pageNum: 1
				},
				total: 0,
			}
		},
		created() {
			this.getArticleList();
		},
		methods: {
			// 获取文章列表
			async getArticleList() {
				const { data: res } = await this.$axios.get('articles', {
					params: {
						pageSize: this.queryParam.pageSize,
						pageNum: this.queryParam.pageNum
					}
				});
				this.articleList = res.data;
				this.total = res.total;
			},


		},
	}
</script>
<style>
	.article-info a {
		text-decoration: none;
		color: #67b3ef;
	}

	.article-info a:hover {
		color: rgb(0, 150, 136);
	}
</style>
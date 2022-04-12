<template>
	<div>
		<div v-if="total == 0 && isLoad" class="d-flex justify-center align-center">
			<div>
				<v-alert class="ma-5" dense outlined type="error">抱歉，你搜索的文章标题不存在！</v-alert>
			</div>
		</div>

		<v-sheet>
			<v-card elevation="3" class="article-info ma-4" v-for="article in artList" :key="article.ID">
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
			<v-col>
				<div class="text-center">
					<v-row justify="center">
						<v-col cols="8">
							<v-container class="max-width">
								<v-pagination circle :total-visible="7" v-model="queryParam.pagenum"
									:length="Math.ceil(this.total/queryParam.pagesize)" class="my-2"
									@input="getArtList()">
								</v-pagination>
							</v-container>
						</v-col>
					</v-row>
				</div>
			</v-col>
		</v-sheet>
	</div>
</template>
<script>
	export default {
		props: ['keywords'],
		data() {
			return {
				artList: [],
				queryParam: {
					pagesize: 5,
					pagenum: 1
				},
				total: 0,
				isLoad: false
			}
		},
		created() {
			this.getArtList();
		},
		methods: {
			// 获取文章列表
			async getArtList() {
				const { data: res } = await this.$axios.get(`articles`, {
					params: {
						keywords: this.keywords,
						pagesize: this.queryParam.pagesize,
						pagenum: this.queryParam.pagenum
					}
				});

				this.artList = res.data;
				this.total = res.total;
				this.isLoad = true;
			}
		}
	}
</script>
<style>
</style>
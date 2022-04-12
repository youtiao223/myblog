<template>
	<v-card class="mx-auto" max-width="300">
		<v-img :src="profile.backGroundImg">
			<v-card-title>
				<v-col align="center">
					<v-avatar size="130" color="grey">
						<v-img :src="profile.avatar"></v-img>
					</v-avatar>
					<div class="ma-4 white--text text-h4">{{profile.name}}</div>
				</v-col>
			</v-card-title>
			<v-divider></v-divider>
		</v-img>

		<v-card-title>关于我</v-card-title>
		<v-card-text>
			<div class="text-center font-weight-bold">{{profile.motto}}</div>
			<p class="ma-2 grey--text">{{profile.selfIntroduction}}</p>
		</v-card-text>
		<v-divider color="lime lighten-1"></v-divider>

		<v-list nav dense>
			<v-list-item>
				<v-list-item-icon class="ma-3">
					<v-icon color="blue darken-2">{{'mdi-qqchat'}}</v-icon>
				</v-list-item-icon>
				<v-list-item-content class="grey--text">{{profile.qq}}</v-list-item-content>
			</v-list-item>
			<v-list-item>
				<v-list-item-icon class="ma-3">
					<v-icon color="yellow darken-2">{{'mdi-email'}}</v-icon>
				</v-list-item-icon>
				<v-list-item-content class="grey--text">{{profile.email}}</v-list-item-content>
			</v-list-item>
			<!-- <v-list-item>
				<v-list-item-icon class="ma-3">
					<v-icon color="blue darken-2">{{'mdi-youtube'}}</v-icon>
				</v-list-item-icon>
				<v-list-item-content class="grey--text">
					<a class="bili-link grey--text" :href="profile.bili" target="_blank">BiliBili</a>
				</v-list-item-content>
			</v-list-item> -->
			<v-list-item>
				<v-list-item-icon class="ma-3">
					<v-icon>{{'mdi-github'}}</v-icon>
				</v-list-item-icon>
				<v-list-item-content class="grey--text">
					<a class="bili-link grey--text" :href="profile.github" target="_blank">Github</a>
				</v-list-item-content>
			</v-list-item>
		</v-list>

	</v-card>
</template>
<script>
	export default {
		data() {
			return {
				profile: {
					id: 13,
				},

			}
		},
		created() {
			this.getProfileInfo();
		},
		methods: {
			// 获取个人信息
			async getProfileInfo() {
				const { data: res } = await this.$axios.get(`/user/profile/${this.profile.id}`);
				this.profile = res.data
			},
		},
	}
</script>
<style scoped>
	.bili-link {
		text-decoration: none;
	}
</style>
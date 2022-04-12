<template>
	<div>
		<a-card>
			<h3>个人设置</h3>
			<a-form-model>
				<!-- 作者名称 -->
				<a-form-model-item label="作者名称" prop="author">
					<a-input v-model="profile.author"></a-input>
				</a-form-model-item>
				<!-- 座右铭 -->
				<a-form-model-item label="座右铭" prop="motto">
					<a-input v-model="profile.motto"></a-input>
				</a-form-model-item>
				<!-- 个人介绍 -->
				<a-form-model-item label="个人介绍" prop="selfIntroduction">
					<a-textarea v-model="profile.selfIntroduction"></a-textarea>
				</a-form-model-item>
				<!-- QQ -->
				<a-form-model-item label="QQ" prop="qq">
					<a-input v-model="profile.qq"></a-input>
				</a-form-model-item>
				<!-- 微信 -->
				<a-form-model-item label="微信" prop="weChat">
					<a-input v-model="profile.weChat"></a-input>
				</a-form-model-item>
				<!-- 远程仓库 -->
				<a-form-model-item label="Github" prop="github">
					<a-input v-model="profile.github"></a-input>
				</a-form-model-item>
				<!-- 邮箱 -->
				<a-form-model-item label="邮箱" prop="email">
					<a-input v-model="profile.email"></a-input>
				</a-form-model-item>
				<!-- 背景图 -->
				<a-form-model-item label="背景图" prop="backGroundImg">
					<a-input v-model="profile.backGroundImg"></a-input>
				</a-form-model-item>
				<!-- 头像 -->
				<a-form-model-item label="头像" prop="avatar">
					<a-input v-model="profile.avatar"></a-input>
				</a-form-model-item>
				<!-- 提交 -->
				<a-form-model-item>
					<a-button type="danger" style="margin-right:15px" @click="updateProfile">更新</a-button>
				</a-form-model-item>
			</a-form-model>
		</a-card>
	</div>
</template>
<script>
	export default {
		data() {
			return {
				profile: {
					uid: 13,
					author: '',
					motto: '',
					selfIntroduction: '',
					qq: '',
					weChat: '',
					github: '',
					email: '',
					backGroundImg: '',
					avatar: ''
				}
			}
		},
		created() {
			this.getProfileInfo();
		},
		methods: {
			async getProfileInfo() {
				const { data: res } = await this.$axios.get(`/user/profile/${this.profile.uid}`);
				if (res.status !== 200) return this.$message.error(res.message)
				this.profile = res.data
			},
			// 更新
			async updateProfile() {
				const { data: res } = await this.$axios.put(`/user/profile/${this.profile.uid}`, this.profile)
				if (res.status !== 200) return this.$message.error(res.message)
				this.$message.success(`个人信息更新成功`)
				this.$router.push('/index')
			},
		},
	}
</script>
<style>

</style>
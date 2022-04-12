<!-- 登录页面 -->
<template>
  <div class="container">
    <div class="loginBox">
      <a-form-model
        :model="formdata"
        :rules="rules"
        ref="loginFormRef"
        class="loginForm"
      >
        <a-form-model-item prop="username" label="用户名">
          <a-input v-model="formdata.username" placeholder="请输入用户名">
            <a-icon
              slot="prefix"
              type="user"
              style="color: rgba(0, 0, 0, 0.25)"
            />
          </a-input>
        </a-form-model-item>
        <a-form-model-item prop="password" label="密码">
          <a-input
            v-model="formdata.password"
            placeholder="请输入密码"
            type="password"
            v-on:keyup.enter="login()"
          >
            <a-icon
              slot="prefix"
              type="lock"
              style="color: rgba(0, 0, 0, 0.25)"
            />
          </a-input>
        </a-form-model-item>
        <a-form-model-item class="loginBtn">
          <a-button type="primary" style="margin: 0 20px" @click="login()"
            >登录</a-button
          >
          <a-button type="info" style="margin: 0 20px" @click="reset()"
            >重置</a-button
          >
        </a-form-model-item>
      </a-form-model>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      formdata: {
        username: "",
        password: "",
      },
      rules: {
        username: [
          { required: true, message: "请输入用户名", trigger: "blur" },
          {
            min: 4,
            max: 12,
            message: "用户名在4到12字符之间",
            trigger: "blur",
          },
        ],
        password: [
          { required: true, message: "请输入密码", trigger: "blur" },
          { min: 6, max: 20, message: "密码在6到20字符之间", trigger: "blur" },
        ],
      },
    };
  },
  //生命周期 - 创建完成（访问当前this实例）
  created() {},
  //生命周期 - 挂载完成（访问DOM元素）
  mounted() {},
  //方法集
  methods: {
    // 取消（重置表单）
    reset() {
      this.$refs.loginFormRef.resetFields();
    },
    // 登录
    login() {
      this.$refs.loginFormRef.validate(async (val) => {
        if (!val) {
          this.$message.error("输入非法数据，请重新输入！");
          return;
        }
        const { data: res } = await this.$axios.post("login", this.formdata);
        // console.log(res);
        if (res.status != 200) {
          this.$message.error(res.message);
          return;
        }
        window.sessionStorage.setItem("token", res.token);
        this.$router.push("/index");
      });
    },
  },
};
</script>

<style scoped>
.container {
  height: 100%;
  background-color: #747d8c;
  /* display: flex;
       justify-content: center;
       align-items: center; */
}

.loginBox {
  width: 500px;
  height: 350px;
  background-color: azure;
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  border-radius: 5px;
}

.loginForm {
  width: 100%;
  position: absolute;
  bottom: 10%;
  padding: 0 20px;
  box-sizing: border-box;
}

.loginBtn {
  display: flex;
  justify-content: center;
}
</style>
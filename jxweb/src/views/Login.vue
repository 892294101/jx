<template>
  <div class="login_container">
    <div class="login_box">
      <el-form class="login_form" ref="loginFormRef" :rules="rules" :model="loginForm">
        <div class="title">松鼠后台管理系统</div>
        <el-form-item prop="username">
          <el-input ref="usernameRef" prefix-icon="el-icon-user-solid" placeholder="账号" v-model="loginForm.username" maxlength="20" clearable></el-input>
        </el-form-item>
        <el-form-item prop="password">
          <el-input placeholder="密码" prefix-icon="el-icon-key" minlength="5" maxlength="18" v-model="loginForm.password" clearable show-password></el-input>
        </el-form-item>
        <el-form-item prop="image">
          <el-input placeholder="验证码" @input="checkCodeValidate" prefix-icon="el-icon-view" style="width: 200px; float: left;" minlength="4" maxlength="4" v-model="loginForm.image" clearable/>
          <el-image class="captchaImg" style="width: 150px; float: left;" :src="image" @click="getCaptcha"/>
        </el-form-item>
        <el-form-item>
          <el-row :gutter="20">
            <el-col :span="12" :offset="0">
              <el-button type="primary" class="login-but" @click="loginBtn">登录</el-button>
            </el-col>
            <el-col :span="12" :offset="0">
              <el-button class="login-reset" type="info" @click="resetLoginForm">重置</el-button>
            </el-col>
          </el-row>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
export default {
  name: "Login",
  data() {
    return {
      image: '',
      rules: {
        username: [{required: true, message: "请输入账号", trigger: "blur"}],
        password: [{required: true, message: "请输入密码", trigger: "blur"}],
        image: [{required: true, message: "请输入验证码", trigger: "blur"}]
      },
      loginForm: {
        username: '',
        password: '',
        image: '',
        idKey: ''
      }
    }
  },
  methods: {
    // 获取验证码
    async getCaptcha() {
      const {data: res} = await this.$api.captcha()
      // console.log("获取验证码res数据：", res)
      if (res.code !== 200) {
        this.$message.error(res.message)
      } else {
        this.image = res.data.image
        this.loginForm.idKey = res.data.idKey
      }
    },
    // 登录
    loginBtn() {
      this.$refs.loginFormRef.validate(async valid => {
        if (valid) {
          const {data: res} = await this.$api.login(this.loginForm)
          // console.log("获取登录的res数据", res)
          if (res.code !== 200) {
            this.$message.error(res.message)
            await this.getCaptcha()
            this.$refs.usernameRef.focus()
            this.loginForm.image = ''
            this.loginForm.username = ''
            this.loginForm.password = ''
          } else {
            this.$message.success({message: "登录成功", center: true})
            this.$store.commit('saveSysAdmin', res.data.sysAdmin)
            this.$store.commit('saveToken', res.data.token)
            this.$store.commit('saveLeftMenuList', res.data.leftMenuList)
            this.$store.commit('savePermissionList', res.data.permissionList)
            await this.$router.push("/index")
          }
        } else {
          return false
        }
      })
    },
    // 重置
    resetLoginForm() {
      this.getCaptcha()
      this.$refs.loginFormRef.resetFields()
    },

    // 验证码输入实时效验值
    checkCodeValidate(value) {
      const reg = /^[0-9a-zA-Z]*$/
      if (reg.test(value)) {
        return true
      } else {
        this.loginForm.image = this.loginForm.image.replace(/[^a-zA-Z0-9]/g, "")
      }
    },

    // 处理键盘事件
    keyEnterHandle(event) {
      if (event.keyCode !== undefined) {
        if (event.keyCode == 13) {
          this.loginBtn()
        }
      }
    }

  },

  mounted() {
    // 打开页面，自动焦点到username输入框
    this.$refs.usernameRef.focus()
    // 监听键盘按键
    document.addEventListener("keydown", this.keyEnterHandle, true);
  },
  beforeDestroy() {
    document.removeEventListener("keydown", this.keyEnterHandle, true);
  },
  created() {
    this.getCaptcha()
  },
}
</script>

<style lang="less" scoped>

.login_container {
  background-image: url("../assets/image/background.png");
  background-size: cover;
  height: 100%;
}

.login_box {
  width: 400px;
  height: 330px;
  background-color: #fff;
  border-radius: 1px;
  position: absolute;
  left: 50%;
  top: 50%;
  transform: translate(-50%, -50%);
  border-radius: 8px;
  background-color: rgba(0, 0, 0, 0.5);
}

.login_form {
  position: absolute;
  bottom: 0;
  padding: 0 20px;
  box-sizing: border-box;

}

.title {
  color: white;
  font-size: 23px;
  line-height: 1.5;
  text-align: center;
  margin-bottom: 20px;
  font-weight: bold;
  font-style: italic;
}

.captchaImg {
  height: 38px;
  width: 100%;
  margin-left: 8px;
  border: 0px solid rgba(0, 0, 0, 0.5);
}

.login-but {
  width: 100%;
  font-size: large;
  opacity: 0.9;
  background-image: linear-gradient(to right, #74ebd5 0%, #9face6 100%);
}

.login-reset {
  width: 100%;
  font-size: large;
  opacity: 0.9;
}

::v-deep .el-input__inner {
  color: white;
  background-color: rgba(0, 0, 0, 0.5);
  box-sizing: border-box;
  border: 1px solid rgba(0, 0, 0, 0.5);
}


</style>
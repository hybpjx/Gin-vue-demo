<template>

    <nav class="container is-max-widescreen" role="" aria-label="navigation">

        <div class="register">
            <div class="box">

                <form action="" class="box">
                    <h3 class="title">注册</h3>
                    <div class="field">
                        <label class="label">姓名</label>
                        <div class="control">
                            <input class="input" type="text" placeholder="输入你的名称（选填）" v-model="user.name">
                        </div>
                        <!-- <p class="help is-danger">This username is available</p> -->
                    </div>

                    <div class="field">
                        <label class="label">手机号</label>
                        <div class="control has-icons-left has-icons-right">
                            <input class="input" type="text" placeholder="输入你的手机号" @blur="checkMobile"
                                v-model="user.telephone">
                            <span class="icon is-small is-left">
                                <i class="fas fa-user"></i>
                            </span>
                            <span class="icon is-small is-right">
                                <i class="fas fa-check"></i>
                            </span>
                            <div v-if="showTelephoneValidate">
                                <p class="help is-danger">手机长度不对 或者 手机格式不正确</p>
                            </div>

                        </div>
                    </div>

                    <div class="field">
                        <label class="label">密码</label>
                        <div class="control has-icons-left has-icons-right">
                            <input class="input" type="password" placeholder="请输入你的密码" @blur="checkPassword"
                                v-model="user.password">
                            <span class="icon is-small is-left">
                                <i class="fas fa-envelope"></i>
                            </span>
                            <span class="icon is-small is-right">
                                <i class="fas fa-exclamation-triangle"></i>
                            </span>

                            <div v-if="showPasswordValidate">
                                <p class="help is-danger">密码长度要大于或等于6位</p>
                            </div>

                        </div>
                    </div>


                    <input type="button" style="margin-top: 1em" class="button is-link is-fullwidth" @click="register"
                        value="注册">
                </form>


            </div>
        </div>

    </nav>

    <el-row>
        <el-button>默认按钮</el-button>
        <el-button type="primary">主要按钮</el-button>
        <el-button type="success">成功按钮</el-button>
        <el-button type="info">信息按钮</el-button>
        <el-button type="warning">警告按钮</el-button>
        <el-button type="danger">危险按钮</el-button>
    </el-row>



</template>
<script>
import { ref } from 'vue';
import axios from 'axios';
import { ElMessage } from 'element-plus';
export default {
    setup() {
        const user = ref({
            name: "",
            telephone: "",
            password: "",
        })
        // 判断手机号是否有效
        const showTelephoneValidate = ref(false);
        // 判断密码是否有效
        const showPasswordValidate = ref(false);
        // 点击事件的方法
        function register() {
            // 验证数据
            if (!(this.user.telephone)) {
                this.showTelephoneValidate = true;
                return
            } else if (!(this.user.password)) {
                this.showPasswordValidate = true;
                return
            }
            const api = "http://localhost:1016/api/auth/register"

            axios.post(api, { ...this.user }).then(res => {

                // 保存token


                // 跳转到主页
                console.log(res);
            }).catch(err => {


                const openCenter = () => {
                    ElMessage({
                        showClose: true,
                        message: '用户名或密码错误',
                        center: true,
                    })
                }

                openCenter()

            })

            console.log("注册成功")



        }
        function checkMobile() {
            // 1. 检查手机格式

            if (this.user.telephone?.length !== 11) {
                this.showTelephoneValidate = true;
            }
            else if (!/1[3-9]\d{9}/.test(this.user.telephone)) {
                console.log("手机号码格式错误");
                this.showTelephoneValidate = true;
            } else {
                this.showTelephoneValidate = false;
            }

        }
        function checkPassword() {
            // 1. 检查手机格式
            console.log()
            if (!(this.user.password?.length >= 6)) {
                this.showPasswordValidate = true;
            } else {
                this.showPasswordValidate = false
            }


        }



        return { user, register, showTelephoneValidate, showPasswordValidate, checkMobile, checkPassword }
    }

}
</script>
<style>

</style>

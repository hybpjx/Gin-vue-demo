<template>

    <nav class="container is-max-widescreen" role="" aria-label="navigation">

        <div class="login">
            <div class="box">

                <form action="" class="box">
                    <h3 class="title">登录</h3>


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


                    <input type="button" style="margin-top: 1em" class="button is-link is-fullwidth" @click="login"
                        value="登录">
                </form>


            </div>
        </div>

    </nav>




</template>
<script>
import { ref } from 'vue';
import { ElMessage } from 'element-plus';
import { useRouter } from "vue-router";
import { useStore } from 'vuex';
export default {
    setup() {

        // 使用vue router
        const $router = useRouter();
        // 从vuex 数据仓库中取东西
        const $store = useStore();

        const user = ref({
            telephone: "",
            password: "",
        })

        // 判断手机号是否有效
        const showTelephoneValidate = ref(false);
        // 判断密码是否有效
        const showPasswordValidate = ref(false);

        function checkMobile() {
            // 1. 检查手机格式
            if (!(this.user.telephone)) {
                this.showTelephoneValidate = true;
                return
            }
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
            if (!(this.user.password)) {
                this.showPasswordValidate = true;
                return
            };
            if (!(this.user.password?.length >= 6)) {
                this.showPasswordValidate = true;
            } else {
                this.showPasswordValidate = false
            }
        };



        function login() {

            // 请求登录
            $store.dispatch("userModule/login", this.user).then(() => {
                // 跳转主页
                $router.replace({ name: "home" })// 没有this
            }).catch((err) => {
                console.log(err)
                ElMessage.error(err.response.data.msg)
                return
            })

            console.log("登录成功")
        }


        return { user, login, showTelephoneValidate, showPasswordValidate, checkMobile, checkPassword }
    }

}
</script>
<style>

</style>

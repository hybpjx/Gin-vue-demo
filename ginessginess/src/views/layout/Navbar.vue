<template>
    <div>
        <nav class="navbar is-info" role="navigation" aria-label="main navigation">
            <div class="container">
                <div class="navbar-brand">
                    <a class="navbar-item" @click="$router.push({name:'home'})"> JustSoSo </a>

                    <a role="button" class="navbar-burger" aria-label="menu" aria-expanded="true"
                        data-target="navbarBasicExample">
                        <span aria-hidden="true"></span>
                        <span aria-hidden="true"></span>
                        <span aria-hidden="true"></span>
                    </a>
                </div>

                <div id="navbarBasicExample" class="navbar-menu">

                    <div class="navbar-end" v-if="userInfo">
                        <div class="navbar-item">

                            <div class="buttons">
                                <a class="button">
                                    <p>{{userInfo.name}}</p>
                                </a>
                                <a class="button is-primary" @click="$router.replace({ name: 'profile' })">
                                    <strong>个人主页</strong>
                                </a>
                                <a class="button is-light" @click="logout">
                                    登出
                                </a>
                            </div>
                        </div>
                    </div>



                    <div class="navbar-end" v-else>
                        <div class="navbar-item">
                            <div class="buttons">
                                <button class="button is-primary" v-if="$route.name!='register'"
                                    @click="$router.replace({ name: 'register' })">
                                    <strong>注册</strong>
                                </button>
                                <button class="button is-light" v-if="$route.name!='login'"
                                    @click="$router.replace({ name: 'login' })">
                                    登录
                                </button>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </nav>


    </div>
</template>


<script>

import { computed } from 'vue'
import { useStore } from 'vuex';
import { ElMessage } from 'element-plus';
import { useRouter } from "vue-router";
export default {

    setup() {
        // 使用vue router
        const $router = useRouter();
        // 从vuex 数据仓库中取东西
        const $store = useStore();



        const userInfo = computed(() => {
            //  存的话没问题 在这个页面取值
            let user = JSON.parse($store.state.userModule.userInfo)
            return user;
        })


        function logout() {

            // 请求登录
            $store.dispatch("userModule/logout").then(
                () => {
                    // 跳转主页
                    location.href = window.location.href;
                    // $router.push({ name: "home" })// 没有this
                    console.log("退出成功")
                }).catch((err) => {
                    console.log(err)
                    ElMessage.error("退出错误")
                })

        }




        return { userInfo, logout }
    }

};
</script>


<style>
.navbar {
    align-content: center;
    justify-content: center;
}
</style>

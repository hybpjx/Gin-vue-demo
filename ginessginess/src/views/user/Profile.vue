<template>

    <nav class="container is-max-widescreen" role="" aria-label="navigation">

        <h3 class="title is-3 left">你好 <span v-if="userInfo">{{userInfo.name}}</span></h3>


    </nav>

</template>

<script>
import { computed } from 'vue';
import { useStore } from 'vuex';
export default {

    setup() {
        // 从vuex 数据仓库中取东西
        const $store = useStore();

        const userInfo = computed(() => {
            //  存的话没问题 在这个页面取值
            let user = JSON.parse($store.state.userModule.userInfo)
            return user;
        })
        console.log(userInfo)

        return { userInfo }
    }

}
</script>

<style>

</style>

<template>
    <div>
        <nav class="navbar is-info" role="navigation" aria-label="main navigation">
            <div class="container">
                <div class="navbar-brand">
                    <a class="navbar-item" href="#"> OceanLearn </a>

                    <a role="button" class="navbar-burger" aria-label="menu" aria-expanded="false"
                        data-target="navbarBasicExample">
                        <span aria-hidden="true"></span>
                        <span aria-hidden="true"></span>
                        <span aria-hidden="true"></span>
                    </a>
                </div>

                <div id="navbarBasicExample" class="navbar-menu">


                    <div class="navbar-end">
                        <div class="navbar-item">
                            <div class="buttons">
                                <button class="button is-primary" @click="$router.replace({ name: 'register' })">
                                    <strong>注册</strong>
                                </button>
                                <button class="button is-light" @click="$router.replace({ name: 'login' })">
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
export default {

};
</script>


<style>
.navbar {
    align-content: center;
    justify-content: center;
}
</style>

<template>
  <div id="app">

    <div class="contains is-info">

      <Navbar />

      <!-- <div class="container">

      </div> -->

    </div>


    <router-view />
  </div>

</template>


<script>
import Navbar from "./views/layout/Navbar.vue";

export default {
  components: {
    Navbar
  }
}


</script>


<style lang="scss">

</style>

<template>
  <div class="g-sidenav-show bg-gray-100 ">
    <div class="min-height-300 bg-primary position-absolute w-100"></div>
    <Aside />
    <main class="main-content position-relative border-radius-lg">
      <Nav />
      <div class="container-fluid py-4">
        <router-view></router-view>

        <MyFooter />
      </div>
    </main>
    <Plugin />
  </div>
</template>

<script lang="ts" setup>
import Aside from "@/layout/components/aside.vue";
import MyFooter from "@/layout/components/footer.vue";
import Nav from "@/layout/components/nav.vue";
import Plugin from "@/layout/components/plugin.vue";
</script>

<style scoped></style>

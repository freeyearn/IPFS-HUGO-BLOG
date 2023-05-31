<template>
  <router-view />
</template>
<script setup>
// 解决eltable无法自适应宽度问题
import { onHandleElTableResize } from "@/utils/eltableResize.ts";
onHandleElTableResize();
</script>
<style>
</style>

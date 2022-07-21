<template>
  <!--<img alt="Vue logo" src="./assets/logo.png" />
  <HelloWorld msg="Hello Vue 3.0 + Vite" />-->
  <div class="common-layout">
    <el-container>
      <el-header>Header</el-header>
      <el-main>Main</el-main>
      <el-footer>Footer</el-footer>
    </el-container>
  </div>
</template>

<script>
import HelloWorld from './components/HelloWorld.vue'
import Test from './components/Test.vue'

export default {
  name: 'App',
  components: {
    HelloWorld,
    Test
  }
}
</script>
<style>
.el-header{
  background-color: var(--el-color-primary-light-7);
}
.el-main{
  height: 1000px;
}
.el-footer{
  background-color: var(--el-color-primary-light-7);
}
</style>
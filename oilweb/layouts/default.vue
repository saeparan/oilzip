<template>
  <div class="div-first">
    <my-header/>
    <nuxt/>
  </div>
</template>

<script>
import MyHeader from '~components/header.vue'
export default {
  components: {
    MyHeader
  }
}
</script>

<style>
  #__nuxt, .div-first {
    height: 100%;
    min-height: 100%;
  }
</style>

<style scoped>
  .div-first > div {
    height: 100%;
  }
</style>

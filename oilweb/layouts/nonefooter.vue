<template>
  <div class="div-first">
    <my-header/>
    <nuxt/>
    <my-footer/>
  </div>
</template>

<script>
import MyFooter from '~components/Footer.vue'
import MyHeader from '~components/header.vue'
export default {
  components: {
    MyFooter, MyHeader
  }
}
</script>

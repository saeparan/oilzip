<template>
  <section class="container text-center" style="height: 100%;">
    <h1 class="title">
      <i class="fa fa-bug"></i>
      {{ error.statusCode }}
    </h1>
    <h3 class="info text-danger">
      웹페이지를 잠시 표시할 수 없습니다.
    </h3>
    <h4>
      지금은 곤란하다, 조금만 기다려 달라. (MB)
    </h4>
  </section>
</template>
<script>
export default {
  props: ['error']
}
</script>

<style scoped>
.title
{
  margin-top: 15px;
  font-size: 5em;
}
.info
{
  font-weight: 500;
  margin: 0;
}
</style>

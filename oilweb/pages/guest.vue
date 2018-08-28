<template>
  <div>
    <section class="container">
      <div class="container">
        <div class="row">
            <div class="col-lg-12 text-center">
                <h3>방명록<hr></h3>
                <div class="col-lg-12">
                  하고 싶은 말을 자유롭게 적어주세요.
                </div>
                <div class="col-lg-4">
                  <input type="text" class="form-control" placeholder="닉네임">
                </div>
                <div class="col-lg-12 mt-xs">
                  <textarea class="form-control" placeholder="하고 싶은 말"></textarea>
                </div>
                <div class="col-lg-12 mt-xs text-right">
                  <button type="button" class="btn btn-info" v-on:click="test">쓰기</button>
                </div>
            </div>
        </div>
      </div>
    </section>
  </div>
</template>

<script>
export default {
  data: function() {
    return {
      done: true
    }
  },
  methods: {
    test() {

    }
  },
  mounted: function() {

  },
}
</script>

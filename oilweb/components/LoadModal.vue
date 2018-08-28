<template>
  <div class="modal" id="modal-spinner" tabindex="-1" role="dialog">
      <div class="modal-dialog modal-sm" role="document">
          <div class="modal-content text-center">
            <div class="mb-sm">
              <img src="/spinner.gif" class="mt-xs" width="100%" height="50%">
              <div class="pb-sm"><h4>{{message}}</h4></div>
            </div>
          </div>
      </div>
  </div>
</template>

<script>
export default {
  props: {
    message: {
      type: String,
      required: true
    }
  },
  mounted: function() {
    $('#modal-spinner').modal('show')
  },
}
</script>

<style scoped>

</style>

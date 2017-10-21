<template>
    <div>
        <video ref="video" autoplay playsinline></video>
        <span>peerId: {{ stream.peerId }}</span>
    </div>
</template>

<script>
export default {
  props: ['stream'],
  mounted() {
    let el = this.$refs.video
    el.srcObject = this.stream
    el.play()
  },
  watch: {
    stream: {
      deep: true,
      handler(val, old) {
        console.log('change stream')
      }
    }
  }
}
</script>
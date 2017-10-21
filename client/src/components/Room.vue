<template>
    <div id="room">
        <div style="background-color: red">
        </div>
        <div style="background-color: yellow"></div>
        <div style="background-color: green">
            <video-chat></video-chat>
        </div>
        <div style="background-color: blue"></div>
    </div>
</template>

<script>
  import VideoChat from './VideoChat.vue'
  export default {
    components: {
      VideoChat
    }
  }
</script>

<style scoped>
    #room {
        width: 100vw;
        height: 100vh;
        display: grid;
        grid-template-columns: 70% 30%;
        grid-template-rows: 20% 80%;
    }
</style>
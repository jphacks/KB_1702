<template>
  <div id="app">
    <!--<video-chat></video-chat>-->
    <room></room>
  </div>
</template>

<script>
import Persona from './components/Persona.vue'
import VideoChat from './components/VideoChat.vue'
import Room from './components/Room.vue'

export default {
  name: 'app',
  data () {
    return {
      msg: 'Welcome to Your Vue.js App',
      stream: null
    }
  },
  created() {
    const constraints = {
      audio: true,
      video: true,
    }
    navigator.mediaDevices.getUserMedia({video: true, audio: true}).then(stream => {
      console.log('fuga')
      this.stream = stream
    })
  },
  components: {
    VideoChat,
    Persona,
    Room
  }
}
</script>

<style>
  body {
    padding: 0;
    margin: 0;
  }
</style>

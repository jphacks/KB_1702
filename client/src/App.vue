<template>
  <div id="app">
    <video-chat></video-chat>
  </div>
</template>

<script>
import Persona from './components/Persona.vue'
import VideoChat from './components/VideoChat.vue'

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
    Persona
  }
}
</script>

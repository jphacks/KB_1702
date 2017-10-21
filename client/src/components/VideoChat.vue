<template>
    <div>
        <h1>ビデオチャット</h1>
        <div>
            <video ref="selfVideo"></video>
            <span v-if="skyWay.peer">peerId: {{ skyWay.peer.id }}</span>
        </div>
        <div id="other-videos">
            <persona v-for="stream in skyWay.otherStreams" :stream="stream"></persona>
        </div>
        <input type="text" v-model="roomName">
        <button @click="join">Join</button>
    </div>
</template>

<script>
import Persona from './Persona.vue'
import SkyWay from '../lib/SkyWay'

export default {
  data() {
    return {
      skyWay: {},
      roomName: '',
    }
  },
  mounted() {
    navigator.mediaDevices.getUserMedia({video: true, audio: true}).then(stream => {
      this.skyWay = new SkyWay(stream, {
        key: 'a559a530-2f4b-4c14-a9e3-72c9407503ed',
//        debug: 3,
      })
//      let skyWay = new SkyWay(stream, {
//        key: 'a559a530-2f4b-4c14-a9e3-72c9407503ed',
//        debug: 3,
//      })
//      this.$set(this.$data, 'hoge', skyWay)

      this.skyWay.joinOther = (stream) => {
//        console.log(stream)
      }

      let el = this.$refs.selfVideo
      el.srcObject = this.skyWay.stream
      el.play()
    })
  },
  methods: {
    join() {
      this.skyWay.joinRoom(this.roomName)
    }
  },
  components: {
    Persona
  }
}
</script>
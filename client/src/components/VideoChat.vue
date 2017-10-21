<template>
    <div>
        <div id="self-video">
            <video ref="selfVideo" controls></video>
            <span v-if="skyWay.peer">peerId: {{ skyWay.peer.id }}</span>
        </div>
        <div id="other-videos">
            <persona v-for="stream in skyWay.otherStreams" :stream="stream"></persona>
        </div>
        <div id="chat-actions">
            <button @click="endCall">End Call</button>
        </div>
        <input type="text" v-model="roomName">
        <button @click="join">Join</button>
    </div>
</template>

<script>
import Persona from './Persona.vue'
import SkyWay from '../lib/SkyWay'
import Speak from '../lib/Speak'

export default {
  data() {
    return {
      skyWay: {},
      roomName: '',
      speak: {},
    }
  },
  mounted() {
    navigator.mediaDevices.getUserMedia({video: true, audio: true}).then(stream => {
      this.skyWay = new SkyWay(stream, {
        key: 'a559a530-2f4b-4c14-a9e3-72c9407503ed',
//        debug: 3,
      })

      this.speak = new Speak(stream)



      this.skyWay.joinOther = (stream) => {
      }

      //自分のビデオを再生
      let el = this.$refs.selfVideo
      el.srcObject = this.skyWay.stream
      el.play()
    })
  },
  methods: {
    join() {
      this.skyWay.joinRoom(this.roomName)
    },
    endCall() {
      this.skyWay.endCall()
    }
  },
  components: {
    Persona
  }
}
</script>
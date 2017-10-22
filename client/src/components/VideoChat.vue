<template>
    <div id="video-chat">
        <div id="other-videos">
            <persona v-for="stream in skyWay.otherStreams" :stream="stream"></persona>
            <p>{{ minutes }}</p>
        </div>
        <div id="chat-actions">
            <div id="self-video">
                <video ref="selfVideo" controls muted></video>
            </div>
        </div>
    </div>
</template>

<script>
import Persona from './Persona.vue'
import SkyWay from '../lib/SkyWay'
import Speak from '../lib/Speak'
import SpeechRecognition from '../lib/SpeechRecognition'
export default {
  data() {
    return {
      skyWay: {},
      speak: {},
      minutes: ''
    }
  },
  mounted() {
    navigator.mediaDevices.getUserMedia({video: true, audio: true}).then(stream => {
      this.skyWay = new SkyWay(stream, {
        key: 'a559a530-2f4b-4c14-a9e3-72c9407503ed',
//        debug: 3,
      })

      this.skyWay.peer.on('open', () => {
        this.skyWay.joinRoom(this.roomName)
        this.skyWay.room.on('data', (res) => {
          this.handleOnData(res.data)
        })
        this.$emit('onload', this)
      })



      let speechRecognition = new SpeechRecognition()
      speechRecognition.onResult = (result) => {
        this.minutes += result
      }

      this.speak = new Speak(stream)
      this.speak.onStartSpeak = () => {

        console.log('StartSpeak')
      }


      this.speak.onEndSpeak = () => {
        console.log('EndSpeak')
      }

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
    },
    handleOnData({type, data}) {
      console.log(type, data)
      if(type == 'changeProgress') {
        this.$parent.updateProgress(data.progress)
      }
    },
  },
  computed: {
    roomName() {
      return window.location.href.split('/').reverse()[0]
    }
  },
  components: {
    Persona
  }
}
</script>

<style scoped lang="scss">
    #video-chat {
        height:100%;
        width: 100%;
    }
    #other-videos {
        display: grid;
        grid-template-columns: 20% 20% 20% 20%;
        grid-template-rows: 40% 40%;
        grid-column-gap: 5%;
        grid-row-gap: 5%;
        /*display: flex;*/
        /*justify-content: space-around;*/
        flex-wrap: wrap;
        height: 80%;
        width: 100%;
    }
    #chat-actions {
        height: 20%;
        width: 100%;

        #self-video {
            padding: 5px 0;
            height:100%;
            width: 30%;

            video {
                height:100%;
                width: auto;
            }
        }
    }
</style>

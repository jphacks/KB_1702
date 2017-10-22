<template>
    <div class="persona" :style="{ 'border-color': voltage }">
        <div class="video-area">
            <video ref="video" autoplay playsinline ></video>
            <img v-show="isSpeak" src="../assets/speech-bubble.svg" alt="hukidashi">
        </div>
    </div>
</template>

<script>
    import Speak from '../lib/Speak.js'

export default {
  props: ['stream'],
  data() {
    return {
      speak: {},
      isSpeak: false,
      speakCount: 0,
      interval: null,
      voltage: 'blue'
    }
  },
  mounted() {
    let el = this.$refs.video
    el.srcObject = this.stream
    el.play()

    this.speak = new Speak(this.stream)
    this.speak.onStartSpeak = () => {
      this.isSpeak = true
      this.speakCount++
    }

    this.speak.onEndSpeak = (spealTime) => {
      this.isSpeak = false
    }

    this.voltage = this.calcVoltage()
    this.interval = setInterval(() => {
      this.voltage = this.calcVoltage()
      this.speakCount = 0
    }, 1000 * 45)
  },
  methods: {
    calcVoltage() {
      if(this.speakCount > 10) {
        return 'red'
      } else if(this.speakCount > 7) {
        return 'orange'
      } else if(this.speakCount > 5) {
        return 'yellow'
      } else if(this.speakCount > 3) {
        return 'green'
      } else if(this.speakCount >= 0) {
        return 'blue'
      }
    }
  },
}
</script>

<style scoped lang="scss">
.persona {
  border: solid 4px gray;
  border-radius: 50%;
  overflow: hidden;
  width: 120px;
  height: 120px;
  box-sizing: border-box;
  .video-area {
    position: relative;
    video {
      transform: scale(1.7);
      width: 100%;
      height:100%;
    }
    img {
      display: none;
      position: absolute;
      top: 5px;
      right: 5px;
      width: 50%;
      height: 50%;
    }
  }
}

.is-speak {
  border-color: red;
}

.actions {
  display: flex;
  height: 100px;
}
.action-btn {
  flex: 1;
  height: 100%;
  text-align: center;
  img {
    width: auto;
    height: 50%;
  }
}

</style>

<template>
    <div class="persona" :style="{ 'border-color': voltage }">
        <div class="vide-area">
            <video ref="video" autoplay playsinline controls></video>
            <img v-show="isSpeak" src="../assets/speech-bubble.svg" alt="">
        </div>
        <!--<div class="actions">-->
            <!--<div class="action-btn">-->
                <!--<img src="../assets/like.svg" alt="">-->
            <!--</div>-->
            <!--<div class="action-btn">-->
                <!--<img src="../assets/thumb-up.svg" alt="">-->
            <!--</div>-->
        <!--</div>-->
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
        border-radius: 10px;
        overflow: hidden;
        /*display: inline-block;*/
        width: 100%;
        height: 100%;
        /*margin-left: 2%;*/
        /*margin-top: 6.66%;*/
        box-sizing: border-box;

        .vide-area {
            display: inline-block;
            position: relative;
            video {
                width: 100%;
                height:100%;
            }

            img {
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
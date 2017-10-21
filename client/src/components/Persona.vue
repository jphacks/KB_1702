<template>
    <div class="persona" :class="{'is-speak': isSpeak}">
        <video ref="video" autoplay playsinline controls></video>
        <div class="actions">
            <div class="action-btn">
                <img src="../assets/like.svg" alt="">
            </div>
            <div class="action-btn">
                <img src="../assets/thumb-up.svg" alt="">
            </div>
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
      isSpeak: false
    }
  },
  mounted() {
    let el = this.$refs.video
    el.srcObject = this.stream
    el.play()

    this.speak = new Speak(this.stream)
    this.speak.onStartSpeak = () => {
      this.isSpeak = true
    }

    this.speak.onEndSpeak = (spealTime) => {
      this.isSpeak = false
    }
  }
}
</script>

<style scoped lang="scss">
    .persona {
        border: solid 4px gray;
        border-radius: 10px;
        overflow: hidden;
        flex: 1;
        margin: 40px;


        video {
            width: 100%;
            height: auto;
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
        height: 100%;
        img {
            width: 100%;
            height: 100%;
        }
    }


</style>
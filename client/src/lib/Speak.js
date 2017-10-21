const SPEAK_VOLUME = 3
const SPEAK_TIME = 3

export default class Speak
{
  constructor(stream) {
    window.AudioContext = window.AudioContext || window.webkitAudioContext;
    let audioContext = new AudioContext();
    let mediaStreamSource = audioContext.createMediaStreamSource(stream)
    let processor = audioContext.createScriptProcessor(512);
    mediaStreamSource.connect(processor)
    processor.connect(audioContext.destination)

    this.speakCounter = 0
    //喋った判定
    let counter = {
      '1': 0,
      '10': 0,
      '100': 0,
      '1000': 0,
      '10000': 0,
    }

    this.lastSampling = new Date()

    let notSpeakingCount = 0;
    let speakFlag = false
    this.overFlag = false
    this.speakingFlag = false
    this.speakStartTime = null
    // setInterval(() => {
    //   console.log(counter)
    //   if(Speak.isSpeak()) {
    //     notSpeakingCount++
    //     if(notSpeakingCount > 3 && speakFlag) {
    //       notSpeakingCount = 0
    //       speakFlag = false
    //       console.log('EndSpeak')
    //     }
    //   } else {
    //     notSpeakingCount = 0
    //     if(!speakFlag) {
    //       console.log('StartSpeak')
    //       speakFlag = true
    //     }
    //     // console.log('Speaking')
    //   }
    //
    //   counter = {
    //     '1': 0,
    //     '10': 0,
    //     '100': 0,
    //     '1000': 0,
    //     '10000': 0,
    //   }
    // }, 1000 * 1)

    processor.onaudioprocess = (event) => {
      let volume = Speak.instrumentationVolume(event)
      this.overFlag = Speak.isOver(volume)
      if(this.isSamplingTiming()) {
        // console.log(this.speakingFlag)
        if(this.overFlag && this.isSpeak()) {
          //SPEAK_TIME秒の間にvolumeが規定値を超えた、かつspeakingFlagがfalse場合
          this.startSpeak()
        } else if(this.isSilence()) {
          this.endSpeak()
        }
        this.updateLastSampling()
      }

    }
  }

  static instrumentationVolume(event) {
    let buf = event.inputBuffer.getChannelData(0);
    let bufLength = buf.length;
    let sum = 0;
    let x;
    for (var i=0; i<bufLength; i++) {
      x = buf[i];

      sum += x * x;
    }
    return sum
  }

  //ボリュームが規定値を超えたか
  static isOver(volume) {
    return volume > SPEAK_VOLUME
  }

  isSpeak() {
    return !this.speakingFlag
  }

  isSilence() {
    return this.elapsedTime > SPEAK_TIME && this.speakingFlag
  }

  isSamplingTiming() {
    return this.elapsedTime > 1500
  }

  //最後のサンプリングからの経過時間
  get elapsedTime() {
    let now = new Date()
    return now.getTime() - this.lastSampling.getTime()
  }

  updateLastSampling() {
    this.lastSampling = new Date()
  }

  startSpeak() {
    this.onStartSpeak()
    this.speakStartTime = new Date().getTime()
    this.speakingFlag = true
  }

  endSpeak() {
    let now = new Date().getTime()
    this.onEndSpeak(now - this.speakStartTime)
    this.speakingFlag = false
  }

  onStartSpeak() {
    console.log('StartSpeak')
  }

  onEndSpeak() {
    console.log('EndSpeak')
  }
}
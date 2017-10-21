const SPEAK_VOLUME = 20
const SPEAK_TIME = 2

export default class Speak
{
  constructor(stream) {
    this.onStartSpeak = () => {}
    this.onEndSpeak = (speakTime) => {}

    window.AudioContext = window.AudioContext || window.webkitAudioContext;
    let audioContext = new AudioContext();
    let mediaStreamSource = audioContext.createMediaStreamSource(stream)
    let processor = audioContext.createScriptProcessor(512);
    mediaStreamSource.connect(processor)
    processor.connect(audioContext.destination)

    this.notSpeakCount = 0;
    this.lastSampling = new Date()
    this.overFlag = false
    this.speakingFlag = false
    this.speakStartTime = null

    this.intervalKey = setInterval(() => {
      if(this.overFlag) {
        this.notSpeakCountReset()
        if(this.isSpeak()) this.startSpeak()
      } else {
        this.notSpeakCountUp()
        if(this.isSilence()) this.endSpeak()
      }
      this.overFlag = false
    }, 1000)

    processor.onaudioprocess = (event) => {
      this.volume = Speak.instrumentationVolume(event)
      if(Speak.isOver(this.volume)) this.overFlag = true
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

  notSpeakCountUp() {
    this.notSpeakCount++
  }

  notSpeakCountReset() {
    this.notSpeakCount = 0
  }

  isSpeak() {
    return !this.speakingFlag
  }

  isSilence() {
    return this.notSpeakCount > SPEAK_TIME && this.speakingFlag
  }

  startSpeak() {
    console.debug('StartSpeak')
    this.onStartSpeak()
    this.speakStartTime = new Date().getTime()
    this.speakingFlag = true
  }

  endSpeak() {
    console.debug('EndSpeak')
    let now = new Date().getTime()
    this.notSpeakCountReset()
    this.onEndSpeak(now - this.speakStartTime)
    this.speakingFlag = false
  }


}
export default class SpeechRecognition
{
  constructor() {
    window.SpeechRecognition = window.SpeechRecognition || webkitSpeechRecognition
    this.boot = false
    this.onResult = () => {}

    this.setup()
  }

  setup() {
    console.log('setup')
    let recognition = new webkitSpeechRecognition()
    recognition.lang = 'ja';
    recognition.interimResults = true;
    recognition.continuous = true;

    recognition.onsoundend = () => {
      console.log('onsoundend')

    }

    recognition.onresult = (event) => {
      let transcript = '';
      let results = event.results;
      for (var i = event.resultIndex; i < results.length; i++) {
        if (results[i].isFinal) {
          transcript += results[i][0].transcript;
          this.onResult(transcript)
          this.boot = false

          this.setup()
          let interval = setInterval(() => {
            if(this.boot) {
              clearInterval(interval)
            } else {
              this.setup()
              this.boot = true
            }
          }, 1000)

        }
        else {
          transcript += results[i][0].transcript;
        }
      }
    }

    recognition.start()
  }

  start() {
    console.debug('recognition start')
    this.recognition.start()
  }

  stop() {
    this.recognition.stop()
    return new Promise((resolve, reject) => {
      this.recognition.onresult = (event) => {
        resolve(event)
      }
    })
  }
}





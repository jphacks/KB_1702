window.SpeechRecognition = window.SpeechRecognition || webkitSpeechRecognition

const recognition = new webkitSpeechRecognition()
recognition.lang = 'ja';

export default class SpeechRecognition
{
  static start() {
    recognition.start()
  }

  static stop() {
    return recognition.stop()
  }
}





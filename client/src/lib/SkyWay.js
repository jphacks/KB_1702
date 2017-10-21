export default class SkyWay
{
  constructor(stream, option) {
    this.peer = new Peer(option)
    this.stream = stream
    this.otherStreams = []
    this.joinOther = () => {}
  }

  joinRoom(roomName) {
    this.room = this.peer.joinRoom(roomName, {mode: 'sfu', stream: this.stream })

    this.room.on('stream', (stream) => {
      this.otherStreams.push(stream)
      //CallBack
      this.joinOther(stream)
    })
  }
}
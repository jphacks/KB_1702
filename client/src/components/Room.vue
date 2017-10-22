<template>
    <div id="room">
        <div style="background-color: red">
        </div>
        <div style="background-color: yellow">
          <Agenda :agenda="roomData.agenda" :progress="roomData.progress" @recive-next-progress="changeProgress" />
          <Decision :agenda="roomData.agenda" :progress="roomData.progress" />
        </div>
        <div style="background-color: green">
            <video-chat ref="video_chat" @onload="roomOnLoad"></video-chat>
        </div>
        <div style="background-color: blue"></div>
    </div>
</template>

<script>
  import VideoChat from './VideoChat.vue'
  import Agenda from './agenda/agenda.vue'
  import Decision from './decision/decision.vue'

  export default {
    data() {
      return {
        roomData: {
          id: "dwabdhjwabkjdbadkad",
          name: "うぇいサウンド",
          progress: 1,
          start_at: "1995-01-11T06:25:13+09:00",
          end_at: "1995-01-11T06:25:13+09:00",
          agenda: [
            {
              id: 1,
              title: "アイデア出し1",
              goal: "アイデアを10個出す",
              time: 10,
              start_at: "1995-01-11T06:25:13+09:00",
              end_at: "1995-01-11T06:25:13+09:00",
              child: [
                {
                  id: 2,
                  title: "アイデア出し2",
                  goal: "アイデアを10個出す",
                  time: 10,
                  start_at: "1995-01-11T06:25:13+09:00",
                  end_at: "1995-01-11T06:25:13+09:00"
                }
              ]
            },
            {
              id: 3,
              title: "アイデア出し3",
              goal: "アイデアを10個出す",
              time: 10,
              start_at: "1995-01-11T06:25:13+09:00",
              end_at: "1995-01-11T06:25:13+09:00"
            }
          ]
        },
        videoChatVm: null
      }
    },
    mounted() {
    },
    methods: {
      changeProgress(progress) {
        this.roomData.progress = progress;
        this.videoChatVm.skyWay.send('changeProgress', { progress })
      },
      roomOnLoad(videoChatVm) {
        this.videoChatVm = videoChatVm
      },
      updateProgress(progress) {
        console.log(progress)
        this.roomData.progress = progress
      }
    },
    components: {
      VideoChat,
      Agenda,
      Decision,
    },
  }
</script>

<style scoped>
    #room {
        width: 100vw;
        height: 100vh;
        display: grid;
        grid-template-columns: 70% 30%;
        grid-template-rows: 20% 80%;
    }
</style>

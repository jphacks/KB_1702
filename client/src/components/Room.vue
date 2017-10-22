<template>
  <div id="room">
    <div class="left">
      <TimeSchedule
        room-start="2017-10-21T16:40:00+09:00"
        title="JPHACKS いろいろ考える会"
      />
      <VideoChat />
    </div>
    <div class="right">
      <Agenda :agenda="roomData.agenda" :progress="roomData.progress" @recive-next-progress="changeProgress" />
      <Decision :agenda="roomData.agenda" :progress="roomData.progress" />
    </div>
  </div>
</template>

<script>
  import TimeSchedule from './timeSchedule/index.vue';
  import VideoChat from "./VideoChat.vue";
  import Agenda from "./agenda/agenda.vue";
  import Decision from "./decision/decision.vue";

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
              title: "アイデア出し",
              goal: "アイデアを10個出す",
              time: 10,
              start_at: "1995-01-11T06:25:13+09:00",
              end_at: "1995-01-11T06:25:13+09:00"
            },
            {
              id: 2,
              title: "アイデア出し",
              goal: "アイデアを10個出す",
              time: 10,
              start_at: "1995-01-11T06:25:13+09:00",
              end_at: "1995-01-11T06:25:13+09:00"
            },
            {
              id: 3,
              title: "アイデア出し",
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
      TimeSchedule
    },
  }
</script>

<style lang="scss" scoped>
  #room {
    display: flex;
    width: 100vw;
    .left {
      width: 50vw;
      background:-webkit-gradient(linear, left top, right bottom, from(rgba(74,185,236,0.7)), to(rgba(110,13,193,0.7)));
      background:-webkit-linear-gradient(left top, rgba(74,185,236,0.7), rgba(110,13,193,0.7));
      background:-moz-linear-gradient(left top, rgba(74,185,236,0.7), rgba(110,13,193,0.7));
      background:-o-linear-gradient(left top, rgba(74,185,236,0.7), rgba(110,13,193,0.7));
      background:linear-gradient(to right bottom, rgba(74,185,236,0.7), rgba(110,13,193,0.7));
    }
    .right {
      width: 50vw;
    }
  }
</style>

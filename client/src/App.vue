<template>
  <div>
    <!-- <TopView /> -->
    <Agenda :agenda="roomData.agenda" :progress="roomData.progress" v-on:recive-next-progress="changeProgress" />
    <Decision :agenda="roomData.agenda" :progress="roomData.progress" />
  </div>
</template>

<script>
import TopView from "./components/topView/index.vue";
import Agenda from "./components/agenda/agenda.vue";
import Decision from "./components/decision/decision.vue";

export default {
  name: "app",
  components: {
    // TopView,
    Agenda,
    Decision
  },
  methods: {
    changeProgress(progress) {
      this.roomData.progress = progress;
    }
  },
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
            title: "うぇい",
            goal: "アイデアを10個出す",
            time: 10,
            start_at: "1995-01-11T06:25:13+09:00",
            end_at: "1995-01-11T06:25:13+09:00",
          },
          {
            id: 2,
            title: "ぶい出し",
            goal: "アイデアを10個出す",
            time: 10,
            start_at: "1995-01-11T06:25:13+09:00",
            end_at: "1995-01-11T06:25:13+09:00"
          },
          {
            id: 3,
            title: "アイ",
            goal: "アイデアを10個出す",
            time: 10,
            start_at: "1995-01-11T06:25:13+09:00",
            end_at: "1995-01-11T06:25:13+09:00"
          }
        ]
      }
    };
  }
};
</script>

<style lang="scss">
#app {
  font-family: "Avenir", Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
}
ul {
  list-style-type: square;
  margin-left: 2em;
}
</style>

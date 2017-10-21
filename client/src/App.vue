<template>
  <div id="app">
    <TopView />
    <Agenda :agenda="roomData.agenda" :progress="roomData.progress" v-on:change-progress="changeProgress" />
  </div>
</template>

<script>
import TopView from "./components/topView/index.vue";
import Agenda from "./components/agenda/agenda.vue";
export default {
  name: "app",
  components: {
    // TopView,
    Agenda
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
            title: "アイデア出し",
            goal: "アイデアを10個出す",
            time: 10,
            start_at: "1995-01-11T06:25:13+09:00",
            end_at: "1995-01-11T06:25:13+09:00",
            child: [
              {
                id: 2,
                title: "アイデア出し",
                goal: "アイデアを10個出す",
                time: 10,
                start_at: "1995-01-11T06:25:13+09:00",
                end_at: "1995-01-11T06:25:13+09:00",
                child: [
                  {
                    id: 4,
                    title: "アイデア出し",
                    goal: "アイデアを10個出す",
                    time: 10,
                    start_at: "1995-01-11T06:25:13+09:00",
                    end_at: "1995-01-11T06:25:13+09:00"
                  }
                ]
              }
            ]
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
  list-style-type: disc;
  margin-left: 2em;
}
</style>

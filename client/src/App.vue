<template>
  <div id="app">
    <TimeSchedule
      title="JPHACKS つくるもの会議"
      room-start="2017-10-21T15:00:00+09:00"
    />
  </div>
</template>

<script>
import TimeSchedule from './components/timeSchedule/index.vue'

export default {
  name: 'app',
  components: {
    TimeSchedule
  }
}
</script>

<style lang="scss">
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
}
</style>

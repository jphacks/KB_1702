<template>
  <div id="app">
    <TimeSchedule />
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

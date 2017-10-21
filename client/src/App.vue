<template>
  <div id="app">
    <CreateRoom />
  </div>
</template>

<script>
import CreateRoom from './components/createRoom/index.vue'

export default {
  name: 'app',
  components: {
    CreateRoom
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

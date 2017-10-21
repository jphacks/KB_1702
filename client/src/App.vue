<template>
  <div id="app">
    <TopView />
  </div>
</template>

<script>
import TopView from './components/topView/index.vue'

export default {
  name: 'app',
  components: {
    TopView
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

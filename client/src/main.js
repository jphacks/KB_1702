import Vue from 'vue'
import App from './App.vue'

import Room from './components/Room.vue'
Vue.component('Room', Room)

new Vue({
  el: '#app',
  render: h => h(App)
})

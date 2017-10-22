import Vue from 'vue'
import Buefy from 'buefy'
import 'buefy/lib/buefy.css'

Vue.use(Buefy)

import Room from './components/Room.vue'
Vue.component('Room', Room)

import App from './App.vue'
Vue.component('App', App)

import CreateRoom from './components/createRoom/roomCreate.vue'
Vue.component('room-create', CreateRoom)

import AgendaItem from './components/agenda/agendaItem.vue'
Vue.component('AgendaItem', AgendaItem)

new Vue({
  el: '#app',
})

new Vue({
  el: '#room-create'
})

new Vue({
  el: '#room'
})
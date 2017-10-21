<template>
  <div class="agenda-field" :style="{marginLeft: agenda.margin + 'rem'}">

   <h1>タイトル: <input type="text" v-model="title" :value="title"></h1>

   <template v-if="agenda.child">
    <agenda-field
      v-for="childAgenda in agenda.child"
      :key="childAgenda.id"
      :agenda="childAgenda"
    />
   </template>
  </div>
</template>

<script>
export default {
  name: 'agenda-field',
  props: ['agenda'],
  data() {
    return {
      title: this.agenda.title,
      goal: this.agenda.goal
    }
  },
  methods: {
    deleteAgenda() {
      // agenda配列から自分を消す
      console.log(this.agenda.id)
    },
    addAgenda() {
      console.log("add agenda")
    }
  }
}
</script>

<style lang="scss">
.agenda-field {
}
</style>

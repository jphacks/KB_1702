<template>
  <div class="agenda-field" :style="{marginLeft: agenda.margin + 'rem'}">

    <div class="agenda-card">
      <div class="agenda-card-detail">
        <h1>タイトル: <input type="text" v-model="title" :value="title"></h1>
        <h1>ゴール: <input type="text" v-model="goal" :value="goal"></h1>
        <h1>目標時間: <input type="number" v-model="targettime" :value="targettime"></h1>
      </div>
      <a class="button is-danger delete-agenda" @click="destroyAgenda">×</a>
    </div>
    <a
      class="button is-primary child-agenda"
      @click="childAgenda"
      v-if="agenda.isParent"
    >詳細を追加</a>

  </div>
</template>

<script>
import agendaManager from '../../lib/AgendaManager';

export default {
  name: 'agenda-field',
  props: ['agenda'],
  data() {
    return {
      title: this.agenda.title,
      goal: this.agenda.goal,
      targettime: this.agenda.time
    }
  },
  methods: {
    destroyAgenda() {
      agendaManager.destroy(this.agenda.id);
    },
    childAgenda() {
      agendaManager.child(this.agenda.id)
    }
  }
}
</script>

<style lang="scss">
.agenda-field {
  box-shadow: 0px 0px 1px black;
  width: 300px;
  margin: 10px 0;
  padding: 10px;
  position: relative;
  .agenda-card {
    display: flex;
    align-items: center;
    .delete-agenda {
      position: absolute;
      top: 5px;
      right: 5px;
    }
    .agenda-card-detail {
      display: flex;
      flex-direction: column;
      align-items: flex-end;
    }
  }
  .child-agenda {
    margin-top: 10px;
  }
}
</style>

<template>
  <div class="agenda-field">

    <div class="agenda-card">
      <div class="agenda-card-detail">
        <p>タイトル: <input type="text" v-model="data.title" :value="data.title" placeholder="議題のタイトルを入力してください"></p>
        <p>ゴール: <input type="text" v-model="data.goal" :value="data.goal" placeholder="この議題のゴールを設定してください"></p>
        <p>目標時間: <input type="number" v-model="data.time" :value="data.time"></p>
      </div>
      <a class="button is-danger delete-agenda" @click="destroyAgenda">×</a>
    </div>

  </div>
</template>

<script>
import agendaManager from '../../lib/AgendaManager';

export default {
  name: 'agenda-field',
  props: ['agenda'],
  data() {
    return {
      data: this.agenda
    }
  },
  methods: {
    destroyAgenda() {
      agendaManager.destroy(this.agenda.id);
    }
  }
}
</script>

<style lang="scss">
.agenda-field {
  box-shadow: rgba(0, 0, 0, 0.12) 0px 1px 6px, rgba(0, 0, 0, 0.12) 0px 1px 4px;
  width: 90%;
  margin: 10px auto;
  padding: 10px;
  position: relative;
  .agenda-card {
    display: flex;
    align-items: center;
    .delete-agenda {
      position: absolute;
      top: 10px;
      right: 10px;
    }
    .agenda-card-detail {
      display: flex;
      flex-direction: column;
      align-items: flex-start;
      width: 100%;
      p {
        width: 100%;
        font-size: 1rem;
        margin: 10px 0;
      }
      input {
        display: block;
        width: 80%;
        font-size: 1.1rem;
        padding: 5px;
        outline: none;
      }
    }
  }
}
</style>

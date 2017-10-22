<template>
  <div id="create-room">

    <div class="field">
      <label>ルーム名</label>
      <input type="text" v-model="titleInput" :value="titleInput" />
    </div>

    <div class="field">
      <label>アジェンダを追加する</label>
      <template v-for="agenda in agendas">
        <AgendaField
          :agenda="agenda"
          :key="agenda.id"
        />
      </template>
      <button @click="create" class="button is-primary">新しくアジェンダを追加</button>
    </div>

  </div>
</template>

<script>
import AgendaField from './agendaField.vue';
import agendaManager from '../../lib/AgendaManager';

export default {
  name: 'create-room',
  components: {
    AgendaField
  },
  data() {
    return {
      titleInput: '',
      agendas: agendaManager.data
    }
  },
  methods: {
    create() {
      agendaManager.create();
    }
  }
}
</script>

<style lang="scss" scoped>
#create-room {
  .button {
    margin-top: 10px;
    display: block;
  }
}
</style>

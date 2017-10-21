<template>
  <div id="create-room">

    <div class="field">
      <label>ルーム名</label>
      <input type="text" v-model="titleInput" :value="titleInput" />
    </div>

    <div class="field">
      <label>アジェンダを追加する</label>
    </div>

  </div>
</template>

<script>

import AgendaField from './agendaField.vue';

const MOCK_AGENDAS = [
  {
    id: 1,
    title: 'アイデア出し',
    goal: 'アイデアを10個だす',
    time: 10,
    margin: 1,
    child: [
      {
        id: 2,
        title: 'アイデア出し',
        goal: 'アイデアを10個だす',
        time: 10,
        margin: 2,
        child: []
      }
    ]
  }
];

export default {
  name: 'create-room',
  components: {
    AgendaField
  },
  data() {
    return {
      titleInput: '',
      agendas: MOCK_AGENDAS
    }
  }
}
</script>

<style lang="scss">
#create-room {
}
</style>

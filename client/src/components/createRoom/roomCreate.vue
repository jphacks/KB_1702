<template>
  <div id="create-room">

    <div class="wrapper">
      <div class="field agenda-setting">
        <template v-for="agenda in agendas">
          <AgendaField
            :agenda="agenda"
            :key="agenda.id"
          />
        </template>
        <button @click="create" class="new-agenda-button">+</button>
      </div>

      <div class="buttons">
        <label>ルーム名</label>
        <input type="text" v-model="titleInput" :value="titleInput" />
        <button @click="startMTG" class="start-mtg">会議を始める！</button>
      </div>
    </div>

  </div>
</template>

<script>
import AgendaField from './agendaField.vue';
import agendaManager from '../../lib/AgendaManager';

export default {
  name: 'room-create',
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
    },
    startMTG() {
      const agendaData = JSON.parse(JSON.stringify(agendaManager.data));
      const formatedData = agendaManager.postFormat(this.titleInput, agendaData);
      console.log(formatedData)
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
  .wrapper {
    display: flex;
    padding-bottom: 50px;
    .new-agenda-button {
      background-color: #E91E63;
      color: white;
      width: 90%;
      text-align: center;
      margin: 0 auto 30px;
      border: 0;
      font-weight: bold;
      font-size: 1.5rem;
      padding: 10px 0;
      cursor: pointer;
    }
    .agenda-setting {
      width: 60%;
      height: 60vh;
      display: flex;
      flex-direction: column;
      padding: 50px 0;
    }
    .buttons {
      width: 40%;
      display: flex;
      flex-direction: column;
      position: fixed;
      top: 50%;
      right: 0;
      transform: translateY(-50%);
      button {
        display: block;
        width: 50%;
        margin: 30px auto;
        border: 0;
        padding: 20px;
        font-weight: bold;
        font-size: 1rem;
      }
      label {
        font-size: 1.5rem;
      }
      input {
        padding: 5px;
        font-size: 1.5rem;
        outline: none;
        display: block;
        width: 90%;
      }
      .start-mtg {
        background-color: #8E24AA;
        color: white;
        cursor: pointer;
      }
    }
  }
}
</style>

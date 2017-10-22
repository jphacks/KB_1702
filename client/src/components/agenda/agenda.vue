<template>
    <div class="agenda-list-wrapper">
      <ul id="agenda-list">
          <agendaItem 
              v-for="agendaItem in agenda" 
              :item="agendaItem" 
              :key="agendaItem.id" 
              :progress="progress">
          </agendaItem>
      </ul>
      <nextAgenda :progress="progress" v-on:to-next-progress="changeProgress"></nextAgenda>
    </div>
</template>

<script>
import AgendaItem from "./agendaItem.vue";
import NextAgenda from "./nextAgenda.vue";
import axios from "axios";

export default {
  props: ["agenda", "progress"],
  components: {
    AgendaItem,
    NextAgenda
  },
  methods: {
    changeProgress() {
      let url = "/api/agenda/";

      let path = window.location.pathname.split("/");
      let roomId = path[path.length - 1];
      url += roomId;
      url += "/next";

      axios
        .post(url, {
          finish_agenda_id: this.progress
        })
        .then(response => {
          this.$emit("recive-next-progress", response.data.nextProgress);
        })
        .catch(error => {
          // apiアクセスミスっても動いてるように見える
          this.$emit("recive-next-progress", this.progress + 1);
        });
    }
  }
};
</script>

<style lang="scss" scoped>
.agenda-list-wrapper {
  height: 100%;
}

#agenda-list {
  height: 100%;
  overflow-y: scroll;
}
</style>

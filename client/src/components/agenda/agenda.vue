<template>
    <ul class="agenda-list">
        <agendaItem 
            v-for="agendaItem in agenda" 
            :item="agendaItem" 
            :key="agendaItem.id" 
            :progress="progress">
        </agendaItem>
        <nextAgenda :progress="progress" v-on:to-next-progress="changeProgress"></nextAgenda>
    </ul>
</template>

<script>
import AgendaItem from "./agendaItem.vue";
import NextAgenda from "./nextAgenda.vue";

export default {
  props: ["agenda", "progress"],
  components: {
    AgendaItem,
    NextAgenda
  },
  methods: {
      changeProgress() {
          this.$emit('change-progress', this.progress + 1)
      }
  }
};
</script>

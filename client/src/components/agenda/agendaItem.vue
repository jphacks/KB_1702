<template>
  <li class="agenda-item" :data-agenda-id="item.id">
    <agendaInfo
      :title="item.title"
      :goal="item.goal"
      :time="item.time"
      :class="{progressing: isProgressing}"
    />
  </li>
</template>

<script>
import AgendaInfo from "./agendaInfo.vue";

export default {
  props: ["item", "progress"],
  components: {
    AgendaInfo
  },
  computed: {
    isProgressing() {
      return this.item.id === this.progress;
    }
  }
};
</script>

<style lang="scss">
.agenda-item {
  margin: 10px auto;
  width: 90%;
  box-shadow: rgba(0, 0, 0, 0.12) 0px 1px 6px, rgba(0, 0, 0, 0.12) 0px 1px 4px;
}
</style>


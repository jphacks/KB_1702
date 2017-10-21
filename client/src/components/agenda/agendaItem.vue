<template>
  <li class="agenda-item" :data-agenda-id="item.id">
        <agendaInfo :title="item.title" :goal="item.goal" :time="item.time" :class="{progressing: isProgressing}"></agendaInfo>
        <template v-if="item.child">
            <ul class="agenda-child-list">
                <agenda-item v-for="agenda in item.child" :item="agenda" :key="agenda.id" :progress="progress"></agenda-item>
            </ul>
        </template>
  </li>
</template>

<script>
import AgendaInfo from "./agendaInfo.vue";

export default {
  props: ["item", "progress"],
  components: {
    AgendaInfo
  },
  //   data() {
  //       return {
  //           isProgressing: false
  //       }
  //   },
  //   mounted() {
  //       if (this.item.id === this.progress) {
  //           this.isProgressing = true
  //       }
  //   },
  computed: {
    isProgressing() {
      return this.item.id === this.progress;
    }
  }
};
</script>

<style lang="scss">
.progressing {
  div {
    font-size: 150%;
    color: #2f2f2f;
    font-weight: bolder;
  }
}
</style>


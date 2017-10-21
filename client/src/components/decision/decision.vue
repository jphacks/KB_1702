<template>
  <ul class="decision">
    <DecisionItem
        v-for="decisionItem in agenda"
        :item="decisionItem"
        :key="decisionItem.id"
        :progress="progress">
    </DecisionItem>
  </ul>
</template>

<script>
import DecisionItem from "./decisionItem.vue";

export default {
  props: ["agenda", "progress"],
  components: { DecisionItem },
  created() {
    console.log(this.progress);
  }
};
</script>

<style>

</style>

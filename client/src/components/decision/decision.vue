<template>
  <ul class="decision">
    <DecisionItem
        :progress="progress">
    </DecisionItem>
  </ul>
</template>

<script>
import DecisionItem from "./decisionItem.vue";

export default {
  props: ["agenda", "progress"],
  components: { DecisionItem },
};
</script>

<style>

</style>

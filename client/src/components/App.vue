<template>
  <div>
    <TopView />
  </div>
</template>

<script>
import TopView from "./topView/index.vue";

export default {
  name: "app",
  components: {
    TopView
  }
};
</script>

<style lang="scss">
ul {
  list-style-type: square;
  margin-left: 2em;
}
</style>

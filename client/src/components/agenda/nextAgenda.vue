<template>
  <button class="button is-info" @click="next">NEXT > </button>
</template>

<script>
export default {
  props: ["progress"],
  methods: {
    next() {
      let element = document.getElementById("agenda-list");
      element.scrollTop = element.scrollHeight - element.clientHeight - 10;
      this.$emit("to-next-progress");
    }
  }
};
</script>

<style>

</style>

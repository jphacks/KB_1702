<template>
  <button class="next-button" @click="next">次へ</button>
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

<style lang="scss" scoped>
.next-button {
  width: 70px;
  border: 1px solid #303F9F;
  border-radius: 10px;
  padding: 10px;
  font-size: 0.9rem;
  background-color: transparent;
  color: #303F9F;
  margin: 10px 5%;
}
</style>

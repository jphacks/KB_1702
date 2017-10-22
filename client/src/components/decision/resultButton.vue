<template>
  <button class="next-button-fix" @click="fix">次へ</button>
</template>

<script>
export default {
    methods: {
        fix() {
            this.$emit('fix-decision')
        }
    }
}
</script>

<style>
.next-button-fix {
  width: 70px;
  border: 1px solid #303F9F;
  border-radius: 10px;
  padding: 10px;
  font-size: 0.9rem;
  background-color: transparent;
  color: #303F9F;
  position: absolute;
  bottom: 0;
  right: 5%;
}
</style>

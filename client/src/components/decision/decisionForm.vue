<template>
  <b-field label="決定事項">
      <b-input placeholder="決定事項を入力してください" v-model="decision"></b-input>
  </b-field>
</template>

<script>
export default {
  data() {
    return {
      decision: ""
    };
  },
  updated() {
      this.$emit('update-decision', this.decision)
  }
};
</script>

<style>

</style>

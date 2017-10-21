<template>
  <button class="button is-info" @click="next">NEXT > </button>
</template>

<script>
export default {
  props: ["progress"],
  methods: {
    next() {
        this.$emit('to-next-progress')
    }
  }
};
</script>

<style>

</style>

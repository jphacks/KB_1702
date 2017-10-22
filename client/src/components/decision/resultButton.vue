<template>
  <button class="button is-primary" @click="fix">Fix</button>
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

</style>

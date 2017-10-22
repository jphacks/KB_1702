<template>
  <div class="agenda-info">
    <div class="agenda-info-title">{{title}}</div>
    <div class="agenda-info-goal">{{goal}}</div>
    <div class="agenda-info-time">{{time}} min</div>
  </div>
</template>

<script>
export default {
  props: ["title", "time", "goal", "isProgressing"]
};
</script>

<style lang="scss" scoped>
.agenda-info {
  & div {
    font-size: 1em;
  }
  & .agenda-info-goal {
    margin-left: 1em;
  }
  & .agenda-info-time {
    margin-left: 1em;
  }
}

.progressing {
  background-color: lightsteelblue;
  .label {
    font-size: 150%;
    color: #2f2f2f;
    font-weight: bolder;
  }
}
</style>

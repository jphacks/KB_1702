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
  width: 100%;
  padding: 10px;
  font-weight: bold;
  letter-spacing: 1px;
  .agenda-info-title {
    font-size: 1.1rem;
    &:before {
      content: '議題： ';
    }
  }
  .agenda-info-goal {
    font-size: 1.1rem;
    &:before {
      content: '達成： ';
    }
  }
  .agenda-info-time {
    font-size: 1.1rem;
    &:before {
      content: '目標： ';
    }
  }
}
</style>

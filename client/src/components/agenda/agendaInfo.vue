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

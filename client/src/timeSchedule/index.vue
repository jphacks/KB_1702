<template>
  <div id="time-schedule">
    <h1 class="meeting-title">{{ title }}</h1>
    <div class="elapsed-goal">
      <p class="title">経過時間 / 目標時間</p>
      <p class="time" v-if="isElapsedTimeOver">
        <span class="js-elapsed-time over">{{ time.elapsed }}</span> / {{ time.goal }}
      </p>
      <p class="time" v-else>
        <span class="js-elapsed-time">{{ time.elapsed }}</span> / {{ time.goal }}
      </p>
    </div>
    <div class="total">
      <p class="title">現在の会議時間</p>
      <p class="time">{{ time.total }}</p>
    </div>
  </div>
</template>

<script>
export default {
  name: 'time-schedule',
  data() {
    return {
      isElapsedTimeOver: false,
      title: 'JPHACKS つくるもの会議',
      time: {
        elapsed: '01:00',
        goal: '01:00',
        total: '12:00'
      }
    }
  }
}
</script>

<style lang="scss">
#time-schedule {
  width: 50%;
  .meeting-title {
    margin: 0;
    font-size: 1.5rem;
  }
  .elapsed-goal {
    width: 50%;
    display: flex;
    flex-direction: column;
    align-items: center;
  }
  .total {
    width: 50%;
    display: flex;
    flex-direction: column;
    align-items: center;
  }
  .title {
    margin: 0;
    font-size: 0.7rem;
    letter-spacing: 1.5px;
  }
  .time {
    margin: 0;
    font-size: 1.5rem;
    letter-spacing: 1px;
  }

  .js-elapsed-time.over {
    color: red;
    font-size: 3rem;
  }
}
</style>

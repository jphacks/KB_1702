<template>
  <div id="time-schedule">
    <h1 class="meeting-title">{{ title }}</h1>
    <div class="times">
      <div class="elapsed-goal">
        <p class="title">経過時間 / 目標時間</p>
        <p class="time" v-if="isElapsedTimeOver">
          <span class="js-elapsed-time over">{{ time.agenda.elapsed }}</span> / {{ time.agenda.goal }}
        </p>
        <p class="time" v-else>
          <span class="js-elapsed-time">{{ time.agenda.elapsed }}</span> / {{ time.agenda.goal }}
        </p>
      </div>
      <div class="total">
        <p class="title">現在の会議時間</p>
        <p class="time">{{ time.total }}</p>
      </div>
    </div>
  </div>
</template>

<script>
import logger from '../../lib/Logger';
import formatUtil from '../../lib/FormatUtil';
import moment from 'moment';

const UNIX_TIME_INTERVAL = 32400;

export default {
  name: 'time-schedule',
  props: ['roomStart', 'title'],
  data() {
    return {
      timerManage: {
        timerId: null,
        count: 0
      },
      isElapsedTimeOver: false,
      hoge: 'hoge',
      time: {
        agenda: {
          elapsed: '00:00',
          goal: '00:00'
        },
        total: '00:00'
      }
    }
  },
  // life cycle
  mounted() {
    logger.cycle("mounted");
    this.startRoomTimer();
  },
  beforeDestroy() {
    logger.biglog("beforeDestroy");
    clearTimeout(this.timerManage.timerId);
  },
  methods: {
    startRoomTimer() {
      let count = 0;
      const diff = Math.trunc(moment().diff(moment(this.roomStart)) / 1000);
      this.timerManage.timerId = setInterval(() => {
        count += 1;
        this.time.total = moment.unix(diff + count - UNIX_TIME_INTERVAL).format('HH:mm');
      }, 1000);
    }
  }
}
</script>

<style lang="scss">
.times-flex-column {
  display: flex;
  flex-direction: column;
  align-items: center;
}
#time-schedule {
  width: 50%;
  .meeting-title {
    margin: 0;
    font-size: 1.5rem;
  }
  .times {
    width: 50%;
    display: flex;
    flex-direction: column;
    align-items: flex-start;
    .elapsed-goal,
    .total {
      @extend .times-flex-column;
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
}
</style>

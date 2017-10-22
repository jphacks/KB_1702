<template>
  <div id="time-schedule">
    <h1 class="meeting-title">{{ title }}</h1>
    <div class="times">
      <div class="elapsed-goal">
        <p class="title">経過時間 / 目標時間</p>
        <p class="time">
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

// constants
const UNIX_TIME_INTERVAL = 32400;
const agendaStartTime = "2017-10-21T16:40:00+09:00";
const agendaHopeTime = 10;

export default {
  name: 'time-schedule',
  props: ['roomStart', 'title'],
  data() {
    return {
      timerManage: {
        roomTimerId: null,
        agendaTimerId: null,
        count: 0
      },
      isElapsedTimeOver: false,
      time: {
        agenda: {
          elapsed: '00:00',
          goal: this.formatGoalTime(agendaHopeTime)
        },
        total: '00:00'
      }
    }
  },
  // life cycle
  mounted() {
    logger.cycle("mounted");
    this.startRoomTimer();
    this.agendaTimer(agendaStartTime);
  },
  beforeDestroy() {
    logger.biglog("beforeDestroy");
    clearTimeout(this.timerManage.roomTimerId);
    clearTimeout(this.timerManage.agendaTimerId);
  },
  methods: {
    startRoomTimer() {
      let count = 0;
      const diff = Math.trunc(moment().diff(moment(this.roomStart)) / 1000);
      this.timerManage.roomTimerId = setInterval(() => {
        count += 1;
        this.time.total = moment.unix(diff + count - UNIX_TIME_INTERVAL).format('HH:mm');
      }, 1000);
    },
    formatGoalTime(time) {
      const hour = Math.trunc(time / 60);
      const min = Math.trunc(time % 60);
      return `${formatUtil.zeroSuppress(hour)}:${formatUtil.zeroSuppress(min)}`;
    },
    agendaTimer(startTime) {
      let count = 0;
      const diff = Math.trunc(moment().diff(moment(startTime)) / 1000);
      this.timerManage.agendaTimerId = setInterval(() => {
        count += 1;
        this.time.agenda.elapsed = moment.unix(diff + count - UNIX_TIME_INTERVAL).format('HH:mm');
        // this.isElapsedTimeOver = moment().isAfter(moment(startTime).add(agendaHopeTime, 'minutes'));
      }, 1000);
    }
  }
}
</script>

<style lang="scss" scoped>
.times-flex-column {
  display: flex;
  flex-direction: column;
  align-items: center;
}
#time-schedule {
  width: 99.5%;
  height: 20vh;
  background-color: #fff;
  .meeting-title {
    padding: 10px;
    font-size: 2rem;
  }
  .times {
    width: 50%;
    display: flex;
    align-items: center;
    .elapsed-goal,
    .total {
      margin: 0 20px;
      @extend .times-flex-column;
    }
    .title {
      margin: 0;
      font-size: 0.7rem;
      letter-spacing: 1.5px;
      transform: translateY(5px);
      color: rgba(0,0,0,0.5);
    }
    .time {
      margin: 0;
      font-size: 1.5rem;
      letter-spacing: 1px;
    }
  }
}
</style>

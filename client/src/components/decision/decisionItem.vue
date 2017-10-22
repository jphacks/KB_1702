<template>
  <div class="decision-item">
    <h1>アイデア出し</h1>
    <textarea></textarea>
    <ResultButton v-on:fix-decision="postDecision"></ResultButton>
  </div>
</template>

<script>
import axios from "axios";
import ResultButton from './resultButton.vue';

export default {
  props: ["item", "progress"],
  data() {
    return {
      decision: ""
    };
  },
  components: {
    ResultButton,
  },
  methods: {
    updateDecision(decision) {
      this.decision = decision;
    },
    postDecision() {
      let url = "/api/agenda/";

      let path = window.location.pathname.split("/");
      let roomId = path[path.length - 1];
      url += roomId;
      url += "/decision";

      axios.post(url, {
        room_id: roomId,
        result: this.decision
      });
    }
  }
};
</script>

<style lang="scss">
.decision-item {
  display: flex;
  flex-direction: column;
  position: relative;
  h1 {
    margin: 10px auto;
    width: 90%;
    font-size: 1.1rem;
    font-weight: bold;
  }
  textarea {
    width: 90%;
    height: 20vh;
    margin: 0 auto 50px;
  }
}
</style>

<template>
  <li class="decision-item">
    <DecisionForm :class="{progressing: isProgressing}" v-on:update-decision="updateDecision">
    </DecisionForm>
    <ResultButton v-on:fix-decision="postDecision"></ResultButton>
  </li>
</template>

<script>
import DecisionForm from "./decisionForm.vue";
import ResultButton from "./resultButton.vue";
import axios from "axios";

export default {
  props: ["item", "progress"],
  components: {
    DecisionForm,
    ResultButton
  },
  data() {
    return {
      decision: ""
    };
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
.progressing {
  background-color: lightsteelblue;
  .label {
    font-size: 150%;
    color: #2f2f2f;
    font-weight: bolder;
  }
}
</style>

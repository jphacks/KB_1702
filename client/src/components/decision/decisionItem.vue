<template>
  <li class="decision-item">
    <DecisionForm :class="{progressing: isProgressing}" v-on:update-decision="updateDecision">
    </DecisionForm>
    <ResultButton v-on:fix-decision="postDecision"></ResultButton>
    <template v-if="item.child">
        <ul class="decision-child-list">
            <DecisionItem 
            v-for="decision in item.child" 
            :item="decision" 
            :key="decision.id" 
            :progress="progress">
            </DecisionItem>
        </ul>
    </template>
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
  computed: {
    isProgressing() {
      return this.item.id === this.progress;
    }
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

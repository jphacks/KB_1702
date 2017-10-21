<template>
    <div class="settings">
        <div class="setting-icon" @click="settings">
            <img src="../../assets/settings.png" alt="icon">
        </div>

        <div class="setting-dialog" :class="{open: openFlg}" v-show="openFlg">
            <b-field label="URL">
                <b-input id="room-url" :value="roomUrl" readonly></b-input>
                <button class="room-url-copy button is-primary" @click="copyUrl">Copy</button>
            </b-field>
        </div>
    </div>
</template>

<script>
export default {
  name: "settings",
  data() {
    return {
      openFlg: false,
      roomUrl: ""
    };
  },
  mounted: function() {
    this.roomUrl = window.location;
  },
  methods: {
    settings() {
      if (!this.openFlg) {
        this.openFlg = !this.openFlg;
      } else {
        this.openFlg = !this.openFlg;
      }
    },
    copyUrl() {
      document.querySelector("#room-url>input").select();
      document.execCommand("copy");
    }
  }
};
</script>

<style lang="scss" scoped>
.setting-icon {
  img {
    width: 30px;
    height: 30px;
    user-select: none;
  }
}

.open {
  background-color: lightgray;
  width: 40%;
  padding: 10px 20px;
  button {
    display: inline-block;
    width: 4em;
  }
}
</style>

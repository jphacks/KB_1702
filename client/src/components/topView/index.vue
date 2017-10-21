<template>
  <div id="topView">
    <h1>Google Hangout</h1>
    <button>会議を始める</button>
  </div>
</template>

<script>
export default {
  name: 'topView'
}
</script>

<style lang="scss">
body, html {
  width: 100%;
  height: 100%;
  box-sizing: border-box;
}

* {
  padding: 0;
  margin: 0;
}

#topView {
  width: 100%;
  height: 100vh;
  background-image: url('../../assets/bg.jpg');
  background-size: cover;
  background-position: left bottom;
  h1 {
    color: white;
    font-weight: bold;
    font-size: 4rem;
    position: fixed;
    bottom: 10px;
    left: 50px;
  }
  button {
    font-size: 1.5rem;
    width: 20%;
    border: 1px solid #D81B60;
    padding: 20px;
    position: fixed;
    top: 50%;
    left: 50px;
    color: #D81B60;
    background-color: transparent;
    transform: translateY(-50%);
    cursor: pointer;
    &:hover {
      color: white;
      background-color: #D81B60;
    }
  }
}
</style>

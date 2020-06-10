<template>
  <div id="app">
    <img alt="Vue logo" src="./assets/logo.png">
    <TimeLine msg="Image Chat"/>
  </div>
</template>

<script>
import TimeLine from './components/TimeLine.vue'

export default {
  name: 'app',
  components: {
    TimeLine
  }
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>

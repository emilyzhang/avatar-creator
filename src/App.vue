<template>
  <!-- <img alt="Vue logo" src="./assets/hair1.png"> -->
  <Creator msg="Welcome to Your Vue.js App"/>
</template>

<script>
import Creator from './components/Creator.vue'

export default {
  name: 'App',
  components: {
    Creator,
  }
}
</script>

<style>
@import url('https://fonts.googleapis.com/css2?family=Spartan:wght@100;200;300;400;500;600&display=swap');

html,
body {
  padding: 0;
  margin: 0;
  font-family: 'Spartan', sans-serif, -apple-system, BlinkMacSystemFont, Segoe UI, Oxygen,
    Ubuntu, Cantarell, Fira Sans, Droid Sans, Helvetica Neue, sans-serif;
  background-image: url(https://images-wixmp-ed30a86b8c4ca887773594c2.wixmp.com/intermediary/f/a63491be-b5f0-49b8-bcf0-baf70fc55e02/da7yz1u-90fcafd2-1da5-43a7-9270-dc6a48b1c5a1.png);
  background-repeat: no-repeat;
  background-attachment: fixed;
  background-size: cover;
  background-blend-mode: soft-light;
  background-color: rgba(255, 211, 185, 0.288);
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}
</style>

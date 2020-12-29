<template>
  <!-- <img alt="Vue logo" src="./assets/hair1.png"> -->
  <header class="hello">avatar creator</header>
  <Creator/>
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
  background-color: rgba(255, 211, 185, 0.188);
  text-align: center;
}
.hello {
  border: 2px dashed white;
  text-align: center;
  display: inline-block;
  margin: 50px auto;
  font-size: 50px;
  color: #77A2BC;
  padding: 30px;
  text-shadow: 5 5 30px white;
  -webkit-background-clip: text;
  -moz-background-clip: text;
  background-clip: text;
  -webkit-text-fill-color: transparent; 
  -moz-text-fill-color: transparent;
  background-image: linear-gradient(270deg, #646292, #AEE1E2, white, #AEE1E2, #646292);
}
.hello:hover {
  border: 2px dashed white;
  text-shadow: 0 0 30px #728CA7;
  color: white; 
  -webkit-background-clip: none;
  -moz-background-clip: none;
  background-clip: none;
  -webkit-text-fill-color: white; 
  -moz-text-fill-color: white;
  background-image: white;
}
</style>

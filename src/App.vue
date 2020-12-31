<template>
  <!-- <img alt="Vue logo" src="./assets/hair1.png"> -->
  <header class="hello">avatar creator</header>
  <Creator />
</template>

<script>
import Creator from "./components/Creator.vue";

export default {
  name: "App",
  components: {
    Creator,
  },
};
</script>

<style>
@import url("https://fonts.googleapis.com/css2?family=Spartan:wght@100;200;300;400;500;600&display=swap");

*,
*::before,
*::after {
  box-sizing: border-box;
}

html,
body {
  padding: 0;
  margin: 0;
  font-family: "Spartan", sans-serif, -apple-system, BlinkMacSystemFont,
    Segoe UI, Oxygen, Ubuntu, Cantarell, Fira Sans, Droid Sans, Helvetica Neue,
    sans-serif;
  background-image: url(https://images-wixmp-ed30a86b8c4ca887773594c2.wixmp.com/intermediary/f/a63491be-b5f0-49b8-bcf0-baf70fc55e02/da7yz1u-90fcafd2-1da5-43a7-9270-dc6a48b1c5a1.png);
  background-repeat: no-repeat;
  background-attachment: fixed;
  background-size: cover;
  background-blend-mode: soft-light;
  background-color: rgba(255, 211, 185, 0.188);
  text-align: center;
  height: 100%;
}

button, select {
  font-family: inherit;
  font-size: 1em;
  padding: 8px 8px;
  border-radius: 0px;
  margin: 10px;
  /* background: #aee1e2; */
  outline: none;
  border: none;
  color: #77a2bc;
  box-shadow: 0px 0px 70px 0px #646292;
}
button:hover { 
  color: #646292;
}
button:active {
  transform: scale(0.98); 
  box-shadow: 3px 2px 22px 1px #646292;
  outline: none;
}

.hello {
  border: 2px solid #aee1e2;
  text-align: center;
  display: inline-block;
  margin: 50px auto;
  font-size: 50px;
  color: #77a2bc;
  padding: 20px;
  text-shadow: 5 5 30px white;
  -webkit-background-clip: text;
  -moz-background-clip: text;
  background-clip: text;
  -webkit-text-fill-color: transparent;
  -moz-text-fill-color: transparent;
  background-image: linear-gradient(
    270deg,
    #646292,
    #aee1e2,
    white,
    #aee1e2,
    #646292
  );
}
.hello:hover {
  text-shadow: 0 0 30px #728ca7;
  color: white;
  -webkit-background-clip: none;
  -moz-background-clip: none;
  background-clip: none;
  -webkit-text-fill-color: white;
  -moz-text-fill-color: white;
  background-image: white;
}
</style>

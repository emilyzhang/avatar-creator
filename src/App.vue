<template>
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

button,
select {
  font-family: inherit;
  font-size: 1em;
  padding: 8px 8px;
  border-radius: 0px;
  margin: 10px;
  /* background: #aee1e2; */
  outline: none;
  border: 2px solid #77a2bc;
  background: #faf6ed;
  color: #77a2bc;
  /* box-shadow: 0px 0px 70px 0px #646292; */
}
button:hover {
  color: #4b669d;
  box-shadow: 0px 0px 20px 0px #aee1e2;
}
button:active {
  transform: scale(0.98);
}
</style>

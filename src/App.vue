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

/* RESET */
html,
body,
div,
span,
object,
iframe,
h1,
h2,
h3,
h4,
h5,
h6,
p,
blockquote,
pre,
abbr,
address,
cite,
code,
del,
dfn,
em,
img,
ins,
kbd,
q,
samp,
small,
strong,
sub,
sup,
var,
b,
i,
dl,
dt,
dd,
ol,
ul,
li,
fieldset,
form,
label,
legend,
table,
caption,
tbody,
tfoot,
thead,
tr,
th,
td,
article,
aside,
canvas,
details,
figcaption,
figure,
footer,
header,
hgroup,
menu,
nav,
section,
summary,
time,
mark,
audio,
video {
  margin: 0;
  padding: 0;
  border: 0;
  outline: 0;
  font-size: 100%;
  vertical-align: baseline;
  background: transparent;
}

body {
  line-height: 1;
}

article,
aside,
details,
figcaption,
figure,
footer,
header,
hgroup,
menu,
nav,
section {
  display: block;
}

nav ul {
  list-style: none;
}

blockquote,
q {
  quotes: none;
}

blockquote:before,
blockquote:after,
q:before,
q:after {
  content: "";
  content: none;
}

a {
  margin: 0;
  padding: 0;
  font-size: 100%;
  vertical-align: baseline;
  background: transparent;
}

/* change colours to suit your needs */
ins {
  background-color: #c4f3f4;
  color: #77a2bc;
  text-decoration: none;
}

/* change colours to suit your needs */
mark {
  background-color: #c4f3f4;
  color: #77a2bc;
  font-style: italic;
  font-weight: bold;
}

del {
  text-decoration: line-through;
}

abbr[title],
dfn[title] {
  border-bottom: 1px dotted;
  cursor: help;
}

table {
  border-collapse: collapse;
  border-spacing: 0;
}

/* change border colour to suit your needs */
hr {
  display: block;
  height: 1px;
  border: 0;
  border-top: 2px solid #aee1e2;
  margin: 1em 0;
  padding: 0;
}

input,
select {
  vertical-align: middle;
}

/* RESET COMPLETE */

html,
body {
  padding: 0;
  margin: 0;
  font-family: "Spartan", sans-serif, -apple-system, BlinkMacSystemFont,
    Segoe UI, Oxygen, Ubuntu, Cantarell, Fira Sans, Droid Sans, Helvetica Neue,
    sans-serif;
  text-align: center;
  height: 100%;
}

#app {
  background-attachment: fixed;
  background-size: cover;
  background-image: url("/assets/background.png");
  background-blend-mode: soft-light;
  background-color: rgba(252, 216, 196, 0.103);
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
/* button:active {
  transform: scale(0.98);
} */

/* Push Animation */
@-webkit-keyframes hvr-push {
  50% {
    -webkit-transform: scale(0.8);
    transform: scale(0.8);
  }
  100% {
    -webkit-transform: scale(1);
    transform: scale(1);
  }
}
@keyframes hvr-push {
  50% {
    -webkit-transform: scale(0.8);
    transform: scale(0.8);
  }
  100% {
    -webkit-transform: scale(1);
    transform: scale(1);
  }
}
.hvr-push {
  display: inline-block;
  vertical-align: middle;
  -webkit-transform: perspective(5px) translateZ(0);
  transform: perspective(5px) translateZ(0);
  box-shadow: 0 0 1px rgba(0, 0, 0, 0);
}
.hvr-push:active {
  -webkit-animation-name: hvr-push;
  animation-name: hvr-push;
  -webkit-animation-duration: 0.3s;
  animation-duration: 0.3s;
  -webkit-animation-timing-function: linear;
  animation-timing-function: linear;
  -webkit-animation-iteration-count: 1;
  animation-iteration-count: 1;
}
</style>

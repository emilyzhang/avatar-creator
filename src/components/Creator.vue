<template>
  <div class="grid-container">
    <div class="creator">
      <canvas id="pixi"></canvas>
    </div>
    <div class="color-picker">
      <div id="picker"></div>
      <div class="feature-select" @change="featureSelect($event)">
        <select>
          <option value="eyeColor">eyeColor</option>
          <option value="face">face</option>
          <option value="eye">eye</option>
          <option value="eyebrows">eyebrows</option>
          <option value="hair">hair</option>
          <option value="hairColor">hairColor</option>
          <option value="eyeLight">eyeLight</option>
          <option value="mouth">mouth</option>
          <option value="nose">nose</option>
          <option value="background">background</option>
        </select>
      </div>
    </div>
    <div class="select-color">
      <div id="color-square"></div>
      <!-- <button @click="select('hairColor')">hair color</button>
      <button @click="select('eyeColor')">eye color</button> -->
    </div>
  </div>
  <svg
    xmlns="http://www.w3.org/2000/svg"
    fill-rule="evenodd"
    clip-rule="evenodd"
  >
    <defs>
      <g
        id="handle"
        transform="scale(0.6)"
        fill="#728ca7"
        stroke-width="2"
        stroke="white"
      >
        <!-- this is where the handle svg content starts -->
        <!-- <circle
          r="10"
          fill="#728ca7"
          stroke-width="2"
          stroke="white"
        ></circle> -->
        <path
          d="M12 21.593c-5.63-5.539-11-10.297-11-14.402 0-3.791 3.068-5.191 5.281-5.191 1.312 0 4.151.501 5.719 4.457 1.59-3.968 4.464-4.447 5.726-4.447 2.54 0 5.274 1.621 5.274 5.181 0 4.069-5.136 8.625-11 14.402m5.726-20.583c-2.203 0-4.446 1.042-5.726 3.238-1.285-2.206-3.522-3.248-5.719-3.248-3.183 0-6.281 2.187-6.281 6.191 0 4.661 5.571 9.429 12 15.809 6.43-6.38 12-11.148 12-15.809 0-4.011-3.095-6.181-6.274-6.181"
        />
        <!-- this is where the handle svg content ends -->
      </g>
    </defs>
  </svg>
</template>

<script>
import * as PIXI from "pixi.js";
import { ColorOverlayFilter } from "@pixi/filter-color-overlay";
import { PixiConsole, PixiConsoleConfig } from "pixi-console";
import iro from "@jaames/iro";

const avatar = {};
let stage;
let colorPicker;
let selectedFeature = "eyeColor";

// const featuresEnum = Object.freeze({
//   "face": 0,
//   "eyebrows": 1,
//   "hairColor": 2,
//   "hair": 3,
//   "eyeColor": 4,
//   "eyeTop": 5,
//   "eyeLight": 6,
//   "mouth": 7,
//   "nose": 8,
// })

const hexToRGB = (hex) =>
  hex
    .replace(
      /^#?([a-f\d])([a-f\d])([a-f\d])$/i,
      (m, r, g, b) => "#" + r + r + g + g + b + b
    )
    .substring(1)
    .match(/.{2}/g)
    .map((x) => parseInt(x, 16));

const changeColor = (selectedFeature) => {
  if (selectedFeature) {
    var hex = colorPicker.color.hexString;
    console.log("HEX", hex);
    const selectedColor = hexToRGB(hex);
    console.log(selectedColor);
    console.log("attempting to change color");
    console.log("avatar", avatar);
    if (!stage) {
      console.log("stage not ready yet");
    }
    avatar[selectedFeature].filters = [
      new ColorOverlayFilter([
        selectedColor[0] / 255,
        selectedColor[1] / 255,
        selectedColor[2] / 255,
      ]),
    ];
  }
};

export default {
  name: "Creator",
  components: {},
  methods: {
    select(selection) {
      console.log(selection);
      changeColor(selection);
    },
    featureSelect(event) {
      console.log(event.target.value);
      selectedFeature = event.target.value;
      // changeColor(selectedFeature);
    },
    attachConsole(stage) {
      // customize default values of PixiConsole
      const consoleConfig = new PixiConsoleConfig();
      consoleConfig.consoleWidth = 420;
      consoleConfig.consoleHeight = 420;

      let pixiConsole = PixiConsole.getInstance();
      if (!PixiConsole.getInstance()) {
        pixiConsole = new PixiConsole(consoleConfig);
      }
      stage.addChild(pixiConsole);
    },
    drawPixi() {
      var canvas = document.getElementById("pixi");

      const app = new PIXI.Application({
        width: 420,
        height: 420,
        antialias: true,
        transparent: true,
        // backgroundColor: "0xffffff",
        view: canvas,
      });
      stage = app.stage;
      // this.attachConsole(stage)

      // const renderer = PIXI.autoDetectRenderer(420, 420);
      // var screenRect = new PIXI.Rectangle(0,0, renderer.width/renderer.resolution, renderer.height/renderer.resolution);
      // stage.filterArea = screenRect;

      const hairContainer = new PIXI.Container();
      const hair = new PIXI.Sprite.from("../assets/hair1.png");
      const nose = new PIXI.Sprite.from("../assets/nose0.png");
      const mouth = new PIXI.Sprite.from("../assets/mouth0.png");
      const eye = new PIXI.Sprite.from("../assets/eye1top.png");
      const eyeColor = new PIXI.Sprite.from("../assets/eye1color.png");
      const eyeLight = new PIXI.Sprite.from("../assets/eye1light.png");
      const eyebrows = new PIXI.Sprite.from("../assets/eyebrows0.png");
      const hairColor = new PIXI.Sprite.from("../assets/hair1color.png");
      const face = new PIXI.Sprite.from("../assets/face1.png");
      const background = new PIXI.Sprite(PIXI.Texture.WHITE);
      background.width = 420;
      background.height = 420;
      background.tint = 0xffffff;
      face.filters = [new ColorOverlayFilter([1, 0.845, 0.7])];
      eyebrows.filters = [new ColorOverlayFilter([0, 0.8, 0.83])];
      hairColor.filters = [new ColorOverlayFilter([0, 0.8, 0.83])];
      // hairColor.filters = [new ColorOverlayFilter([0.55, 0.4, 0.3])];
      hair.alpha = 0.5;
      hair.filters = [new ColorOverlayFilter([0.8, 1, 0.9])];
      eyeColor.filters = [new ColorOverlayFilter([0.2, 0.2, 0.25])];
      stage.addChild(background);
      hairContainer.addChild(hairColor);
      hairContainer.addChild(hair);
      // console.log(hair);
      // console.log(face);
      stage.addChild(face);
      stage.addChild(eyebrows);
      stage.addChild(hairContainer);
      stage.addChild(eyeColor);
      stage.addChild(eye);
      stage.addChild(eyeLight);
      stage.addChild(mouth);
      stage.addChild(nose);
      avatar.face = face;
      avatar.eyebrows = eyebrows;
      avatar.hair = hair;
      avatar.hairColor = hairColor;
      avatar.eyebrows = eyebrows;
      avatar.eyeLight = eyeLight;
      avatar.eye = eye;
      avatar.mouth = mouth;
      avatar.nose = nose;
      avatar.eyeColor = eyeColor;
      avatar.background = background;
    },
  },

  mounted() {
    this.drawPixi();
    colorPicker = new iro.ColorPicker("#picker", {
      // Set the size of the color picker
      width: 200,
      // Set the initial color to pure red
      color: "#f00",
      display: "inline-block",
      borderColor: "white",
      padding: 0,
      borderWidth: 2,
      // handleRadius: 10,
      handleSvg: "#handle",
      handleProps: { x: -8, y: -8 },
      layout: [
        {
          component: iro.ui.Wheel,
        },
        {
          component: iro.ui.Slider,
        },
      ],
    });
    colorPicker.on("color:change", function (color) {
      console.log("New active color:", color.hexString);
      changeColor(selectedFeature);
    });
  },
};
</script>

<style>
.grid-container {
  display: grid;
  /* grid-template-areas: "lside avatar colorpicker random rside"; */
  grid-template-areas: "lside avatar colorpicker rside";
  grid-template-rows: 424px;
  /* grid-template-columns: 20px 425px 300px auto 20px; */
  grid-template-columns: 1fr 425px 300px 1fr;
  grid-gap: 15px;
  align-content: center;
}
.creator {
  border: 2px dashed white;
  /* margin: auto 100px auto; */
  /* padding: 4px; */
  width: 420px;
  grid-area: avatar;
}
.select-color {
  /* border: 2px dashed white; */
  /* margin: auto auto 100px; */
  /* padding: 4px; */
  grid-area: random;
}
.color-picker {
  border: 2px dashed white;
  grid-area: colorpicker;
}
#picker {
  margin: 40px;
}
</style>

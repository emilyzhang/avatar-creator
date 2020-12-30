<template>
  <div class="grid-container">
    <div class="creator">
      <canvas id="pixi"></canvas>
    </div>
    <div class="color-picker">
    <div id="picker"></div>
    </div>
    <div class="select-color">
      <div id="color-square"></div>
      <button @click="select('hairColor')">hair color</button>
      <button @click="select('eyeColor')">eye color</button>
    </div>
  </div>
</template>

<script>
import * as PIXI from "pixi.js";
import { ColorOverlayFilter } from "@pixi/filter-color-overlay";
import { PixiConsole, PixiConsoleConfig } from "pixi-console";
import iro from '@jaames/iro';


const avatar = {};
let stage;
let colorPicker;


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

export default {
  name: "Creator",
  components: {},
  methods: {
    select (selection) {
      console.log(selection)
      this.changeColor(selection);
    },
    changeColor(selectedFeature) {
      if (selectedFeature) {
      var hex = colorPicker.color.hexString;
      console.log("HEX", hex)
      const selectedColor = hexToRGB(hex);
      console.log(selectedColor);
      console.log("attempting to change color");
      console.log("avatar", avatar);
      if (!stage) {
        console.log("stage not ready yet");
      }
      avatar[selectedFeature].filters = [new ColorOverlayFilter([selectedColor[0]/255, selectedColor[1]/255, selectedColor[2]/255])]
      }
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
      const eyeTop = new PIXI.Sprite.from("../assets/eye1top.png");
      const eyeColor = new PIXI.Sprite.from("../assets/eye1color.png");
      const eyeLight = new PIXI.Sprite.from("../assets/eye1light.png");
      const eyebrows = new PIXI.Sprite.from("../assets/eyebrows0.png");
      const hairColor = new PIXI.Sprite.from("../assets/hair1color.png");
      const face = new PIXI.Sprite.from("../assets/face1.png");
      face.filters = [new ColorOverlayFilter([1, 0.845, 0.7])];
      eyebrows.filters = [new ColorOverlayFilter([0, 0.8, 0.83])];
      hairColor.filters = [new ColorOverlayFilter([0, 0.8, 0.83])];
      eyeColor.filters = [new ColorOverlayFilter([0.2, 0.2, 0.25])];
      hair.tint = 0xbeebee;
      hairContainer.addChild(hairColor);
      hairContainer.addChild(hair);
      // console.log(hair);
      // console.log(face);
      stage.addChild(face);
      stage.addChild(eyebrows);
      stage.addChild(hairContainer);
      stage.addChild(eyeColor);
      stage.addChild(eyeTop);
      stage.addChild(eyeLight);
      stage.addChild(mouth);
      stage.addChild(nose);
      avatar.face = face;
      avatar.eyebrows = eyebrows;
      avatar.hair = hair;
      avatar.hairColor = hairColor;
      avatar.eyebrows = eyebrows;
      avatar.eyeLight = eyeLight;
      avatar.mouth = mouth;
      avatar.nose = nose;
      avatar.eyeColor = eyeColor;
    },
  },

  mounted() {
    this.drawPixi();
    colorPicker = new iro.ColorPicker('#picker', {
    // Set the size of the color picker
    width: 300,
    // Set the initial color to pure red
    color: "#f00",
    display: "inline-block",
    borderColor: "#ffffff",
    padding: 0,
    borderWidth: 2,
  });
  },
};
</script>

<style>
.grid-container {
  display: grid;
  grid-template-areas: "lside avatar colorpicker random rside";
  grid-template-rows: auto;
  grid-template-columns: 20px 425px 500px auto 20px;
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
  border: 2px dashed white;
  /* margin: auto auto 100px; */
  /* padding: 4px; */
  grid-area: random;
}
.color-picker {
  border: 2px dashed white;
  grid-area: colorpicker;
}
#picker {
  margin: 20px;
}
</style>

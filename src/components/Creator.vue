<template>
  <div class="grid-container">
    <div class="creator">
      <canvas id="pixi"></canvas>
    </div>
    <div class="color-picker">
      <div id="picker"></div><a v-bind:href="downloadURL" download="avatar.png"
          >
      <button class="save-button" @click="downloadAvatar">
        <svg
            class="download-svg"
            fill="currentColor"
            viewBox="0 0 20 20"
            xmlns="http://www.w3.org/2000/svg"
          >
            <path
              d="M7.707 10.293a1 1 0 10-1.414 1.414l3 3a1 1 0 001.414 0l3-3a1 1 0 00-1.414-1.414L11 11.586V6h5a2 2 0 012 2v7a2 2 0 01-2 2H4a2 2 0 01-2-2V8a2 2 0 012-2h5v5.586l-1.293-1.293zM9 4a1 1 0 012 0v2H9V4z"
            ></path></svg
        >
      </button></a>
      <div class="feature-select" @change="featureSelect($event)">
        <select>
          <option value="eyeColor">eyeColor</option>
          <option value="face">face</option>
          <option value="eyeShape">eyeShape</option>
          <option value="eyebrows">eyebrows</option>
          <option value="hair">hair</option>
          <option value="hairColor">hairColor</option>
          <option value="mouthColor">mouthColor</option>
          <!-- <option value="nose">nose</option> -->
          <option value="background">background</option>
        </select>
      </div>
    </div>
    <div class="select-color">
      <div id="color-square"></div>
      <!-- <button @click="select('hairColor')">hair color</button>
      <button @click="select('eyeColor')">eye color</button> -->
    </div>
    <div class="feature-box"></div>
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
import * as PIXI from "pixi.js-legacy";
import { ColorOverlayFilter } from "@pixi/filter-color-overlay";
import iro from "@jaames/iro";

const avatar = {};
let stage;
let colorPicker;
let selectedFeature = "eyeColor";
let renderer;
// PIXI.settings.FAIL_IF_MAJOR_PERFORMANCE_CAVEAT = false;

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
    if (selectedFeature === "hairColor") {
      avatar.hairBack.filters = [
        new ColorOverlayFilter([
          selectedColor[0] / 255,
          selectedColor[1] / 255,
          selectedColor[2] / 255,
        ]),
      ];
    }
  }
};

export default {
  name: "Creator",
  components: {},
  data () {
    return {
      downloadURL: ""
    }
  },
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
    downloadAvatar() {
      console.log("downloadingAvatar", this.downloadURL)
      this.downloadURL = renderer.view.toDataURL("image/png", 1);
      console.log("after downloadingAvatar", this.downloadURL)
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
        preserveDrawingBuffer: true,
      });
      stage = app.stage;
      renderer = app.renderer;
      console.log("renderer", renderer);
      console.log("app", app);
      // this.attachConsole(stage)

      // const hairContainer = new PIXI.Container();
      const hair = new PIXI.Sprite.from("../assets/hair1shape.png");
      const hairBack = new PIXI.Sprite.from("../assets/hair1back.png");
      const nose = new PIXI.Sprite.from("../assets/nose6.png");
      const mouthLine = new PIXI.Sprite.from("../assets/mouth1line.png");
      const mouthColor = new PIXI.Sprite.from("../assets/mouth1color.png");
      const eyeWhite = new PIXI.Sprite.from("../assets/eye2white.png");
      const eyeShape = new PIXI.Sprite.from("../assets/eye2shape.png");
      const eyeColor = new PIXI.Sprite.from("../assets/eye2color.png");
      const eyeLight = new PIXI.Sprite.from("../assets/eye2glare.png");
      const eyebrows = new PIXI.Sprite.from("../assets/eyebrows0.png");
      const hairColor = new PIXI.Sprite.from("../assets/hair1color.png");
      const face = new PIXI.Sprite.from("../assets/face1.png");
      const background = new PIXI.Sprite(PIXI.Texture.WHITE);
      background.width = 420;
      background.height = 420;
      background.tint = 0xffffff;
      face.filters = [new ColorOverlayFilter([0.9, 0.75, 0.6])];
      eyebrows.filters = [new ColorOverlayFilter([0, 0.6, 0.63])];
      hairColor.filters = [new ColorOverlayFilter([0, 0.8, 0.83])];
      hairBack.filters = [new ColorOverlayFilter([0, 0.8, 0.83])];
      mouthColor.filters = [new ColorOverlayFilter([0.8, 0.4, 0.4])];
      // hairColor.filters = [new ColorOverlayFilter([0.55, 0.4, 0.3])];
      hair.alpha = 0.5;
      mouthColor.alpha = 0.5;
      eyebrows.alpha = 0.8;
      hair.filters = [new ColorOverlayFilter([0.8, 1, 0.9])];
      eyeColor.filters = [new ColorOverlayFilter([0.2, 0.2, 0.25])];
      stage.addChild(background);
      // hairContainer.addChild(hairColor);
      // hairContainer.addChild(hair);
      // console.log(hair);
      // console.log(face);
      // stage.addChild(hairBack);
      stage.addChild(face);
      stage.addChild(eyeWhite);
      stage.addChild(eyeColor);
      stage.addChild(eyeShape);
      stage.addChild(eyeLight);
      stage.addChild(eyebrows);
      // stage.addChild(hairContainer);
      stage.addChild(hairColor);
      stage.addChild(hair);
      stage.addChild(mouthColor);
      stage.addChild(mouthLine);
      stage.addChild(nose);
      avatar.face = face;
      avatar.eyebrows = eyebrows;
      avatar.hair = hair;
      avatar.hairColor = hairColor;
      avatar.eyebrows = eyebrows;
      avatar.eyeLight = eyeLight;
      avatar.eyeShape = eyeShape;
      avatar.mouthColor = mouthColor;
      avatar.nose = nose;
      avatar.eyeColor = eyeColor;
      avatar.background = background;
      avatar.hairBack = hairBack;

      this.downloadURL = renderer.view.toDataURL("image/png", 1);
      console.log(this.downloadURL);
    },
  },

  mounted() {
    this.drawPixi();
    colorPicker = new iro.ColorPicker("#picker", {
      // Set the size of the color picker
      width: 150,
      // Set the initial color to pure red
      color: "#beebee",
      display: "inline-block",
      borderColor: "#aee1e2",
      padding: 0,
      borderWidth: 2,
      // handleRadius: 10,
      handleSvg: "#handle",
      handleProps: { x: -7.3, y: -8 },
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
  grid-template-areas:
    "lside avatar colorpicker rside"
    "lside featurebox featurebox rside";
  grid-template-rows: 424px 400px;
  /* grid-template-columns: 20px 425px 300px auto 20px; */
  grid-template-columns: 1fr 425px 300px 1fr;
  grid-gap: 15px;
  align-content: center;
}
.feature-box {
  grid-area: featurebox;
  background-color: rgba(185, 255, 228, 0.281);
}
.creator {
  border: 2px solid #aee1e2;
  /* margin: auto 100px auto; */
  /* padding: 4px; */
  width: 424px;
  grid-area: avatar;
  background-color: rgba(185, 255, 228, 0.281);
}
.select-color {
  /* border: 2px dashed white; */
  /* margin: auto auto 100px; */
  /* padding: 4px; */
  grid-area: random;
}
.color-picker {
  border: 2px solid #aee1e2;
  background-color: rgba(185, 255, 228, 0.281);
  grid-area: colorpicker;
  display: inline-block;
}
#picker {
  margin: 10px;
  background-color: white;
  border-radius: 10px;
  padding: 5px 5px;
}
.download-svg {
  height: 30px;
}
.save-button {
  margin: 10px;
  padding: 5px;
}
</style>

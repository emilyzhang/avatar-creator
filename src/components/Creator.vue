<template>
  <div class="grid-container">
    <header class="hello">avatar creator</header>
    <div class="creator">
      <canvas id="pixi"></canvas>
    </div>
    <div class="color-picker">
      <a v-bind:href="downloadURL" download="avatar.png">
        <button class="save-button hvr-push" @click="downloadAvatar">
          ♡ save avatar
        </button></a
      >
      <div id="picker"></div>
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
    <div class="feature-box">
      <div class="feature-select-buttons">
        <button
          class="feature-button"
          :class="{ 'active-button': feature === currentFeature }"
          :key="feature"
          v-for="feature in Object.keys(choices)"
          @click="currentFeature = feature"
        >
          {{ feature }}
        </button>
      </div>
      <div class="thumbnail-box">
        <div v-if="choices[currentFeature]">
          <img
            class="thumbnail"
            :key="f"
            v-for="f in Object.keys(choices[currentFeature])"
            @click="selectNewFeature(choices[currentFeature][f])"
            :src="thumbnailFilePath(choices[currentFeature][f])"
          />
        </div>
      </div>
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
        <path
          d="M12 21.593c-5.63-5.539-11-10.297-11-14.402 0-3.791 3.068-5.191 5.281-5.191 1.312 0 4.151.501 5.719 4.457 1.59-3.968 4.464-4.447 5.726-4.447 2.54 0 5.274 1.621 5.274 5.181 0 4.069-5.136 8.625-11 14.402m5.726-20.583c-2.203 0-4.446 1.042-5.726 3.238-1.285-2.206-3.522-3.248-5.719-3.248-3.183 0-6.281 2.187-6.281 6.191 0 4.661 5.571 9.429 12 15.809 6.43-6.38 12-11.148 12-15.809 0-4.011-3.095-6.181-6.274-6.181"
        />
      </g>
    </defs>
  </svg>
</template>

<script>
import * as PIXI from "pixi.js-legacy";
import { ColorOverlayFilter } from "@pixi/filter-color-overlay";
import iro from "@jaames/iro";
const path = require("path");

const avatar = {};
let colorPicker;
let selectedFeature = "eyeColor";
// PIXI.settings.FAIL_IF_MAJOR_PERFORMANCE_CAVEAT = false;

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
  console.log(this.avatarState, this.choices);
  if (selectedFeature) {
    var hex = colorPicker.color.hexString;
    console.log("HEX", hex);
    const selectedColor = hexToRGB(hex);
    console.log(selectedColor);
    console.log("attempting to change color");
    console.log("avatar", avatar);
    if (!this.pixiApp?.stage) {
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
  computed: {},
  data() {
    return {
      choices: {},
      currentFeature: "hair",
      downloadURL: "",
      pixiApp: {},
      avatarState: {},
    };
  },
  methods: {
    select(selection) {
      console.log(selection);
      changeColor(selection);
    },
    featureFilePath(featureType, id, layer) {
      return path.join(
        "..",
        "assets",
        "features",
        featureType + id + layer + ".png"
      );
    },
    thumbnailFilePath(feature) {
      if (feature.thumb) {
        return path.join(
          "..",
          "assets",
          "features",
          this.currentFeature + feature.id + feature.thumb + ".png"
        );
      } else if (feature.isSingleLayer) {
        return path.join(
          "..",
          "assets",
          "features",
          this.currentFeature + feature.id + ".png"
        );
      }
      return path.join(
        "..",
        "assets",
        "features",
        this.currentFeature + feature.id + "line" + ".png"
      );
    },
    selectNewFeature(feature) {
      const currentAvatarFeature = this.avatarState.features[
        this.currentFeature
      ];
      if (currentAvatarFeature?.choice?.isSingleLayer) {
        this.avatarState.features[this.currentFeature].choice = feature;
        this.avatar;
      }
      // if (feature.isSingleLayer) {
      //   console.log(this.currentFeature)
      //   if (currentChoice.isSingleLayer) {
      //     if
      //     console.log("straight up swap!")
      //   }
      // } else {
      //   console.log(feature)
      // }
    },
    featureSelect(event) {
      console.log(event.target.value);
      selectedFeature = event.target.value;
      // changeColor(selectedFeature);
    },
    downloadAvatar() {
      this.downloadURL = this.pixiApp.renderer.view.toDataURL("image/png", 1);
    },
    setupCanvas() {
      var canvas = document.getElementById("pixi");
      this.pixiApp = new PIXI.Application({
        width: 420,
        height: 420,
        antialias: true,
        transparent: true,
        // backgroundColor: "0xffffff",
        view: canvas,
        preserveDrawingBuffer: true,
      });
      this.pixiApp.stage.sortableChildren = true;
      console.log("setup", this.pixiApp, this.pixiApp.stage);
      this.drawInitialAvatar();
    },
    newSprite(featureName, featureID, layerName, color, alpha, position) {
      const sprite = new PIXI.Sprite.from(
        this.featureFilePath(featureName, featureID, layerName)
      );
      if (color) {
        const rgb = hexToRGB(color);
        sprite.filters = [
          new ColorOverlayFilter([rgb[0] / 255, rgb[1] / 255, rgb[2] / 255]),
        ];
      }
      if (alpha) {
        sprite.alpha = alpha;
      }
      sprite.zIndex = position;
      return sprite;
    },
    drawInitialAvatar() {
      const app = this.pixiApp;
      const featureList = this.avatarState.features;
      const newSprite = this.newSprite;
      Object.keys(featureList).map(function (name) {
        const feature = featureList[name];
        if (feature.choice.isSingleLayer) {
          const sprite = newSprite(
            name,
            feature.choice.id,
            "",
            feature.color,
            feature.alpha,
            feature.choice.position
          );
          feature.sprite = sprite;
          app.stage.addChild(sprite);
        } else {
          Object.keys(feature.layers).map(function (layerName) {
            const layer = feature.layers[layerName];
            let alpha;
            if (feature.choice.alpha) {
              alpha = feature.choice.alpha[layerName];
            }
            const sprite = newSprite(
              name,
              feature.choice.id,
              layerName,
              layer.color,
              alpha,
              feature.choice.positions[layerName]
            );
            layer.sprite = sprite;
            app.stage.addChild(sprite);
          });
        }
      });
    },
    setupChoices() {
      this.choices.hair = Object.freeze({
        wavy: {
          name: "wavy",
          id: 1,
          isSingleLayer: false,
          thumb: "line",
          positions: {
            line: 10,
            color: 9,
            back: 0,
          },
          alpha: {
            line: 0.5,
          },
        },
      });
      this.choices.eye = Object.freeze({
        neutralFemale: {
          name: "neutralFemale",
          id: 2,
          isSingleLayer: false,
          positions: {
            line: 4,
            color: 3,
            white: 2,
            glare: 5,
          },
        },
      });
      this.choices.face = Object.freeze({
        neutralFemale: {
          name: "neutralFemale",
          id: 1,
          isSingleLayer: true,
          position: 1,
        },
      });
      this.choices.nose = Object.freeze({
        snub: {
          name: "snub",
          id: 0,
          isSingleLayer: true,
          position: 5,
        },
        simple: {
          name: "simple",
          id: 1,
          isSingleLayer: true,
          position: 5,
        },
        button: {
          name: "button",
          id: 2,
          isSingleLayer: true,
          position: 5,
        },
        wide: {
          name: "wide",
          id: 3,
          isSingleLayer: true,
          position: 5,
        },
        upturned: {
          name: "upturned",
          id: 4,
          isSingleLayer: true,
          position: 5,
        },
        downturned: {
          name: "downturned",
          id: 5,
          isSingleLayer: true,
          position: 5,
        },
        medium: {
          name: "medium",
          id: 6,
          isSingleLayer: true,
          position: 5,
        },
        elegant: {
          name: "elegant",
          id: 7,
          isSingleLayer: true,
          position: 5,
        },
        droopy: {
          name: "droopy",
          id: 8,
          isSingleLayer: true,
          position: 5,
        },
      });
      this.choices.eyebrows = Object.freeze({
        neutral: {
          name: "neutral",
          id: 0,
          isSingleLayer: true,
          position: 8,
        },
        worried: {
          name: "worried",
          id: 1,
          isSingleLayer: true,
          position: 8,
        },
        thick: {
          name: "thick",
          id: 2,
          isSingleLayer: true,
          position: 8,
        },
        surprised: {
          name: "surprised",
          id: 3,
          isSingleLayer: true,
          position: 8,
        },
        resting: {
          name: "resting",
          id: 4,
          isSingleLayer: true,
          position: 8,
        },
      });
      this.choices.mouth = Object.freeze({
        simpleSmile: {
          name: "simpleSmile",
          id: 0,
          isSingleLayer: true,
          alpha: 0.5,
        },
        neutralFemale: {
          name: "neutralFemale",
          id: 1,
          isSingleLayer: false,
          positions: {
            line: 7,
            color: 6,
          },
          alpha: {
            line: 0.5,
          },
        },
      });
    },
    setupInitialAvatar() {
      this.avatarState = {
        background: {
          isTransparent: false,
          color: "#beebee",
        },
        features: {
          face: {
            choice: this.choices.face.neutralFemale,
            color: "#d4b485",
          },
          hair: {
            choice: this.choices.hair.wavy,
            layers: {
              color: { color: "#3dc6db" },
              line: { color: "#cef6f7" },
              back: { color: "#3dc6db" },
            },
          },
          eye: {
            choice: this.choices.eye.neutralFemale,
            layers: {
              line: {},
              color: { color: "#646292" },
              white: {},
              glare: {},
            },
          },
          nose: {
            choice: this.choices.nose.medium,
            layers: {},
          },
          eyebrows: {
            choice: this.choices.eyebrows.resting,
            color: "#beebee",
          },
          mouth: {
            choice: this.choices.mouth.neutralFemale,
            layers: {
              line: {},
              color: {
                color: "#ba6665",
              },
            },
          },
        },
      };
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
      this.pixiApp = app;
      const stage = app.stage;
      const renderer = app.renderer;
      // this.attachConsole(stage)

      // const hairContainer = new PIXI.Container();
      const hair = new PIXI.Sprite.from("../assets/hair1line.png");
      const hairBack = new PIXI.Sprite.from("../assets/hair1back.png");
      const nose = new PIXI.Sprite.from("../assets/nose6.png");
      const mouthLine = new PIXI.Sprite.from("../assets/mouth1line.png");
      const mouthColor = new PIXI.Sprite.from("../assets/mouth1color.png");
      const eyeWhite = new PIXI.Sprite.from("../assets/eye2white.png");
      const eyeShape = new PIXI.Sprite.from("../assets/eye2line.png");
      const eyeColor = new PIXI.Sprite.from("../assets/eye2color.png");
      const eyeLight = new PIXI.Sprite.from("../assets/eye2glare.png");
      const eyebrows = new PIXI.Sprite.from("../assets/eyebrows0.png");
      const hairColor = new PIXI.Sprite.from("../assets/hair1color.png");
      const face = new PIXI.Sprite.from("../assets/face1.png");
      const background = new PIXI.Sprite(PIXI.Texture.WHITE);
      background.width = 420;
      background.height = 420;
      background.tint = 0xfaf6ed;
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
    this.setupChoices();
    this.setupInitialAvatar();
    this.setupCanvas();
    // this.drawPixi();
    colorPicker = new iro.ColorPicker("#picker", {
      // Set the size of the color picker
      width: 250,
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
.hello {
  border: 2px solid #aee1e2;
  text-align: center;
  display: inline-block;
  margin: 25px auto;
  font-size: 30px;
  color: #77a2bc;
  padding: 15px;
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
  grid-area: header;
}
.hello:hover {
  color: #faf6ed;
  text-shadow: 0 0 30px #728ca7;
  -webkit-background-clip: none;
  -moz-background-clip: none;
  background-clip: none;
  -webkit-text-fill-color: #faf6ed;
  -moz-text-fill-color: #faf6ed;
  background-image: #faf6ed;
}
.grid-container {
  display: grid;
  /* grid-template-areas: "lside avatar colorpicker random rside"; */
  grid-template-areas:
    " header header header header header"
    "lside avatar colorpicker featurebox rside"
    "lside savebox savebox featurebox rside";
  grid-template-rows: 117px min-content 150px;
  /* grid-template-columns: 20px 425px 300px auto 20px; */
  grid-template-columns: 0.1fr min-content min-content 1fr 0.1fr;
  grid-gap: 11px;
  align-content: center;
}
.feature-box {
  grid-area: featurebox;
  display: grid;
  grid-template-areas:
    "header"
    "features";
  grid-template-rows: 100px 1fr;
  /* grid-template-columns: 20px 425px 300px auto 20px; */
  grid-template-columns: 1;
  border: 2px solid #aee1e2;
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
  margin: 0px 8px;
  /* background-color: #FAF6ED; */
  /* padding: 3px 3px; */
}
.feature-select-buttons {
  display: inline;
  margin: 10px;
  padding: 5px;
  text-align: left;
  overflow-y: auto;
  max-height: 100px;
  border: 2px solid #aee1e2;
}
.feature-button {
  font-size: 15px;
  padding: 7px;
  margin: 0px 4px 4px 0px;
  background: rgba(27, 27, 36, 0.623);
  color: #cef6f7;
  /* border:none; */
  /* background: none; */
}
.feature-button:hover {
  color: #faf6ed;
}
.feature-button:active {
  transform: scale(0.9);
}
.thumbnail-box {
  padding: 10px;
  overflow-y: auto;
  text-align: left;
  border: 2px solid #aee1e2;
  margin: 10px;
  margin-top: 0px;
  background: rgba(240, 252, 255, 0.596);
}
.thumbnail {
  height: 70px;
  width: 70px;
  border: 2px solid #728ca7;
  padding: 2px;
  margin-left: 5px;
  object-fit: cover;
}
.thumbnail:active {
  transform: scale(0.9);
}
.active-button {
  color: #646292;
  background: rgba(240, 252, 255, 0.596);
}
.active-button:hover {
  color: #646292;
}
</style>

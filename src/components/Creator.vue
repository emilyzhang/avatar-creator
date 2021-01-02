<template>
  <div class="grid-container">
    <header class="hello">avatar creator</header>
    <div class="creator">
      <canvas id="pixi"></canvas>
    </div>
    <div class="color-picker">
      <a
        class="download-avatar"
        v-bind:href="downloadURL"
        download="avatar.png"
      >
        <button class="save-button hvr-push" @click="downloadAvatar">
          ♡ save avatar
        </button></a
      >
      <div id="picker"></div>
      <div class="feature-select" @change="featureSelect($event)">
        Changing: {{ currFeature }}
        <select
          class="select-dropdown"
          v-if="
            avatarState.features[currFeature] &&
            !avatarState.features[currFeature].choice.isSingleLayer
          "
        >
          <option
            :key="layer"
            v-for="layer in Object.keys(
              avatarState.features[currFeature].choice.layers
            ).filter(
              (layer) =>
                !avatarState.features[currFeature].choice.layers[layer]
                  .isNonEditable
            )"
            :selected="currLayer === layer"
          >
            {{ layer }}
          </option>
        </select>
      </div>
    </div>
    <!-- <div class="select-color">
      <div id="color-square"></div>
    </div> -->
    <div class="feature-box">
      <div class="feature-select-buttons">
        <button
          class="feature-button"
          :class="{ 'active-button': feature === currFeature }"
          :key="feature"
          v-for="feature in Object.keys(choices)"
          @click="changeFeatureCategory(feature)"
        >
          {{ feature }}
        </button>
      </div>
      <div class="thumbnail-box">
        <div v-if="choices[currFeature]">
          <img
            class="thumbnail"
            :key="f"
            v-for="f in Object.keys(choices[currFeature])"
            @click="selectNewFeature(choices[currFeature][f])"
            :src="thumbnailFilePath(choices[currFeature][f])"
            :class="{
              'selected-thumbnail':
                choices[currFeature][f].id ==
                avatarState.features[currFeature].choice.id,
            }"
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

let colorPicker;
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

export default {
  name: "Creator",
  components: {},
  computed: {},
  data() {
    return {
      choices: {},
      currFeature: "hair",
      downloadURL: "",
      pixiApp: {},
      avatarState: {
        features: {},
      },
      currLayer: "color",
    };
  },
  methods: {
    console() {
      return window.console;
    },
    changeFeatureCategory(feature) {
      const avatarFeatures = this.avatarState.features;
      this.currFeature = feature;
      if (
        !avatarFeatures[feature].choice.isSingleLayer &&
        avatarFeatures[feature].layers[this.currLayer] === undefined
      ) {
        this.currLayer = Object.keys(avatarFeatures[feature].layers)[0];
      }
    },
    changeLayerColor() {
      const avatarFeatures = this.avatarState.features;
      if (this.currLayer && avatarFeatures[this.currFeature]) {
        var hex = colorPicker.color.hexString;
        // console.log("HEX", hex);
        const selectedColor = hexToRGB(hex);
        // console.log("avatar", avatar);
        if (!this.pixiApp?.stage) {
          console.log("stage not ready yet");
        }
        if (avatarFeatures[this.currFeature].choice.isSingleLayer) {
          avatarFeatures[this.currFeature].sprite.filters = [
            new ColorOverlayFilter([
              selectedColor[0] / 255,
              selectedColor[1] / 255,
              selectedColor[2] / 255,
            ]),
          ];
          avatarFeatures[this.currFeature].color = hex;
        } else {
          const currentFeature = avatarFeatures[this.currFeature].layers;
          currentFeature[this.currLayer].sprite.filters = [
            new ColorOverlayFilter([
              selectedColor[0] / 255,
              selectedColor[1] / 255,
              selectedColor[2] / 255,
            ]),
          ];
          currentFeature[this.currLayer].color = hex;
          // set other layer colors, if applicable
          const otherLayer =
            avatarFeatures[this.currFeature].choice.layers[this.currLayer]
              ?.setOtherLayer;
          if (otherLayer) {
            currentFeature[otherLayer].sprite.filters = [
              new ColorOverlayFilter([
                selectedColor[0] / 255,
                selectedColor[1] / 255,
                selectedColor[2] / 255,
              ]),
            ];
            currentFeature[otherLayer].color = hex;
          }
        }
      }
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
          this.currFeature + feature.id + feature.thumb + ".png"
        );
      } else if (feature.isSingleLayer) {
        return path.join(
          "..",
          "assets",
          "features",
          this.currFeature + feature.id + ".png"
        );
      }
      return path.join(
        "..",
        "assets",
        "features",
        this.currFeature + feature.id + "line" + ".png"
      );
    },
    selectNewFeature(newFeature) {
      const pixiApp = this.pixiApp;
      const currFeature = this.currFeature;
      const currentAvatarFeature = this.avatarState.features[currFeature];
      const newSprite = this.newSprite;
      if (
        !newFeature.isSingleLayer &&
        newFeature.layers[this.currLayer]?.zIndex === undefined
      ) {
        this.currLayer = Object.keys(
          this.avatarState.features[currFeature].layers
        )[0];
      }
      if (currentAvatarFeature?.choice.id !== newFeature.id) {
        // first remove old feature, if it exists
        if (currentAvatarFeature?.choice) {
          // console.log(currentAvatarFeature);
          const oldFeature = currentAvatarFeature.choice;
          // first remove oldfeature
          if (oldFeature.isSingleLayer) {
            pixiApp.stage.removeChild(currentAvatarFeature.sprite);
          } else {
            Object.keys(currentAvatarFeature.layers).map(function (layerName) {
              const layer = currentAvatarFeature.layers[layerName];
              pixiApp.stage.removeChild(layer.sprite);
            });
          }
        }
        // Add new feature
        if (newFeature.isSingleLayer) {
          const sprite = newSprite(
            currFeature,
            newFeature.id,
            "",
            currentAvatarFeature.color,
            newFeature.alpha,
            newFeature.zIndex
          );
          currentAvatarFeature.sprite = sprite;
          pixiApp.stage.addChild(sprite);
        } else {
          Object.keys(newFeature.layers).map(function (layerName) {
            delete currentAvatarFeature.color;
            delete currentAvatarFeature.sprite;
            const oldLayer = currentAvatarFeature.layers[layerName];
            // console.log("old layer", oldLayer);
            let newLayer = newFeature.layers[layerName];
            const sprite = newSprite(
              currFeature,
              newFeature.id,
              layerName,
              oldLayer?.color,
              newLayer.alpha,
              newLayer.zIndex
            );
            if (currentAvatarFeature.layers[layerName] === undefined) {
              currentAvatarFeature.layers[layerName] = {};
            }
            currentAvatarFeature.layers[layerName].sprite = sprite;
            pixiApp.stage.addChild(sprite);
          });
        }
        currentAvatarFeature.isSingleLayer = newFeature.isSingleLayer;
        currentAvatarFeature.choice = newFeature;
        console.log(this.avatarState);
      }
    },
    featureSelect(event) {
      console.log(event.target.value);
      this.currLayer = event.target.value;
      // changeLayerColor(currLayer);
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
      // console.log("setup", this.pixiApp, this.pixiApp.stage);
      this.drawInitialAvatar();
    },
    newSprite(featureName, featureID, layerName, color, alpha, zIndex) {
      // console.log(
      //   "newsprite: ",
      //   featureName,
      //   featureID,
      //   layerName,
      //   color,
      //   alpha,
      //   zIndex
      // );
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
      // console.log(featureName, alpha, layerName)
      sprite.zIndex = zIndex;
      // sprite.anchor.set(-0.5,-0.5);
      // if(featureName ==="hair") {
      // sprite.scale.set(1, 0.8);
      // }
      // sprite.position.set(-4, 14);
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
            feature.choice.alpha,
            feature.choice.zIndex
          );
          feature.sprite = sprite;
          app.stage.addChild(sprite);
        } else {
          Object.keys(feature.layers).map(function (layerName) {
            const layer = feature.layers[layerName];
            const sprite = newSprite(
              name,
              feature.choice.id,
              layerName,
              layer.color,
              feature.choice.layers[layerName].alpha,
              feature.choice.layers[layerName].zIndex
            );
            layer.sprite = sprite;
            app.stage.addChild(sprite);
          });
        }
      });
    },
    setupChoices() {
      this.choices.hair = Object.freeze({
        short: {
          name: "short",
          id: 0,
          isSingleLayer: false,
          thumb: "line",
          layers: {
            line: { zIndex: 100, alpha: 0.455 },
            color: { zIndex: 90, setOtherLayer: "back" },
            // backline: { zIndex: 1, isNonEditable: true },
            backcolor: { zIndex: 0, isNonEditable: true },
          },
        },
        wavy: {
          name: "wavy",
          id: 1,
          isSingleLayer: false,
          thumb: "line",
          layers: {
            line: { zIndex: 100, alpha: 0.455 },
            color: { zIndex: 90, setOtherLayer: "backcolor" },
            // backline: { zIndex: 1, isNonEditable: true },
            backcolor: { zIndex: 0, isNonEditable: true },
          },
        },
        shortBob: {
          name: "shortBob",
          id: 3,
          isSingleLayer: false,
          thumb: "line",
          layers: {
            line: { zIndex: 100, alpha: 0.455 },
            color: { zIndex: 90, setOtherLayer: "backcolor" },
            // backline: { zIndex: 1, isNonEditable: true },
            backcolor: { zIndex: 0, isNonEditable: true },
          },
        },
      });
      this.choices.eye = Object.freeze({
        happyFemale: {
          name: "happyFemale",
          id: 1,
          isSingleLayer: false,
          layers: {
            line: { zIndex: 40 },
            color: { zIndex: 30 },
            glare: { zIndex: 50 },
          },
        },
        neutralFemale: {
          name: "neutralFemale",
          id: 2,
          isSingleLayer: false,
          layers: {
            line: { zIndex: 40 },
            color: { zIndex: 30 },
            glare: { zIndex: 50 },
            white: { zIndex: 20 },
            // pupil: {zIndex: 32, alpha: 0.5}
          },
        },
        sideGlance: {
          name: "sideGlance",
          id: 4,
          isSingleLayer: false,
          layers: {
            line: { zIndex: 40, alpha: 0.8 },
            color: { zIndex: 30 },
            glare: { zIndex: 50 },
            pupil: { zIndex: 32, alpha: 0.7 },
            white: { zIndex: 21 },
            shade: { zIndex: 31, alpha: 0.7 },
          },
        },
        intense: {
          name: "intense",
          id: 5,
          isSingleLayer: false,
          layers: {
            line: { zIndex: 40, alpha: 0.8 },
            color: { zIndex: 30 },
            glare: { zIndex: 50 },
            pupil: { zIndex: 33, alpha: 0.7 },
            white: { zIndex: 21 },
            shade: { zIndex: 31, alpha: 0.7 },
            light: { zIndex: 32, alpha: 0.5 }
          },
        },
        brightEyes: {
          name: "brightEyes",
          id: 6,
          isSingleLayer: false,
          layers: {
            line: { zIndex: 40, alpha: 0.8 },
            color: { zIndex: 30 },
            glare: { zIndex: 50 },
            pupil: { zIndex: 32, alpha: 0.7 },
            white: { zIndex: 21 },
            shade: { zIndex: 31, alpha: 0.5 },
          },
        },
      });
      this.choices.face = Object.freeze({
        pointyChin: {
          name: "pointyChin",
          id: 0,
          isSingleLayer: false,
          layers: {
            line: { zIndex: 10, alpha: 0.2 },
            color: { zIndex: 9 },
          },
        },
        uChin: {
          name: "uChin",
          id: 1,
          isSingleLayer: false,
          layers: {
            line: { zIndex: 10, alpha: 0.2 },
            color: { zIndex: 9 },
          },
        },
        vline: {
          name: "vline",
          id: 2,
          isSingleLayer: false,
          layers: {
            line: { zIndex: 10, alpha: 0.2 },
            color: { zIndex: 9 },
          },
        },
        thinPoint: {
          name: "thinPoint",
          id: 3,
          isSingleLayer: false,
          layers: {
            line: { zIndex: 10, alpha: 0.2 },
            color: { zIndex: 9 },
          },
        },
        square: {
          name: "square",
          id: 5,
          isSingleLayer: false,
          layers: {
            line: { zIndex: 10, alpha: 0.2 },
            color: { zIndex: 9 },
          },
        },
        round: {
          name: "round",
          id: 6,
          isSingleLayer: false,
          layers: {
            line: { zIndex: 10, alpha: 0.2 },
            color: { zIndex: 9 },
          },
        },
        angular: {
          name: "angular",
          id: 7,
          isSingleLayer: false,
          layers: {
            line: { zIndex: 10, alpha: 0.2 },
            color: { zIndex: 9 },
          },
        },
        bumpy: {
          name: "bumpy",
          id: 8,
          isSingleLayer: false,
          layers: {
            line: { zIndex: 10, alpha: 0.2 },
            color: { zIndex: 9 },
          },
        },
        melon: {
          name: "melon",
          id: 9,
          isSingleLayer: false,
          layers: {
            line: { zIndex: 10, alpha: 0.2 },
            color: { zIndex: 9 },
          },
        },
        long: {
          name: "long",
          id: 10,
          isSingleLayer: false,
          layers: {
            line: { zIndex: 10, alpha: 0.2 },
            color: { zIndex: 9 },
          },
        },
      });
      this.choices.nose = Object.freeze({
        snub: {
          name: "snub",
          id: 0,
          isSingleLayer: true,
          zIndex: 50,
          alpha: 0.3,
        },
        simple: {
          name: "simple",
          id: 1,
          isSingleLayer: true,
          zIndex: 50,
          alpha: 0.3,
        },
        button: {
          name: "button",
          id: 2,
          isSingleLayer: true,
          zIndex: 50,
          alpha: 0.3,
        },
        wide: {
          name: "wide",
          id: 3,
          isSingleLayer: true,
          zIndex: 50,
          alpha: 0.3,
        },
        upturned: {
          name: "upturned",
          id: 4,
          isSingleLayer: true,
          zIndex: 50,
          alpha: 0.3,
        },
        downturned: {
          name: "downturned",
          id: 5,
          isSingleLayer: true,
          zIndex: 50,
          alpha: 0.3,
        },
        medium: {
          name: "medium",
          id: 6,
          isSingleLayer: true,
          zIndex: 50,
          alpha: 0.3,
        },
        elegant: {
          name: "elegant",
          id: 7,
          isSingleLayer: true,
          zIndex: 50,
          alpha: 0.3,
        },
        droopy: {
          name: "droopy",
          id: 8,
          isSingleLayer: true,
          zIndex: 50,
          alpha: 0.3,
        },
        subtle: {
          name: "subtle",
          id: 10,
          isSingleLayer: true,
          zIndex: 50,
          alpha: 0.3,
        },
        casual: {
          name: "casual",
          id: 11,
          isSingleLayer: true,
          zIndex: 50,
          alpha: 0.3,
        },
        regular: {
          name: "regular",
          id: 12,
          isSingleLayer: true,
          zIndex: 50,
          alpha: 0.3,
        },
        small: {
          name: "small",
          id: 13,
          isSingleLayer: true,
          zIndex: 50,
          alpha: 0.3,
        },
      });
      this.choices.eyebrows = Object.freeze({
        neutral: {
          name: "neutral",
          id: 0,
          isSingleLayer: true,
          zIndex: 80,
        },
        worried: {
          name: "worried",
          id: 1,
          isSingleLayer: true,
          zIndex: 80,
        },
        thick: {
          name: "thick",
          id: 2,
          isSingleLayer: true,
          zIndex: 80,
        },
        surprised: {
          name: "surprised",
          id: 3,
          isSingleLayer: true,
          zIndex: 80,
        },
        resting: {
          name: "resting",
          id: 4,
          isSingleLayer: true,
          zIndex: 80,
        },
        angry: {
          name: "angry",
          id: 5,
          isSingleLayer: true,
          zIndex: 80,
        },
        eyebrowLift: {
          name: "eyebrowLift",
          id: 9,
          isSingleLayer: true,
          zIndex: 80,
        },
        worriedThin: {
          name: "worriedThin",
          id: 12,
          isSingleLayer: true,
          zIndex: 80,
        },
        neutralRelaxed: {
          name: "neutralRelaxed",
          id: 13,
          isSingleLayer: true,
          zIndex: 80,
        },
        concerned: {
          name: "concerned",
          id: 14,
          isSingleLayer: true,
          zIndex: 80,
        },
        square: {
          name: "square",
          id: 15,
          isSingleLayer: true,
          zIndex: 80,
        }, straight: {
          name: "straight",
          id: 17,
          isSingleLayer: true,
          zIndex: 80,
        },
      });
      this.choices.mouth = Object.freeze({
        simpleSmile: {
          name: "simpleSmile",
          id: 0,
          isSingleLayer: true,
          alpha: 0.5,
          zIndex: 70,
        },
        neutralFemale: {
          name: "neutralFemale",
          id: 1,
          isSingleLayer: false,
          layers: {
            line: { zIndex: 70, alpha: 0.5 },
            color: { zIndex: 60 },
          },
        },
        unhappy: {
          name: "unhappy",
          id: 2,
          isSingleLayer: true,
          alpha: 0.5,
          zIndex: 70,
        },
        openMouth: {
          name: "openMouth",
          id: 3,
          isSingleLayer: false,
          layers: {
            line: { zIndex: 70, alpha: 0.5 },
            teeth: { zIndex: 66 },
            inside: { zIndex: 60 },
          },
        },
        neutral: {
          name: "neutral",
          id: 4,
          isSingleLayer: true,
          alpha: 0.5,
          zIndex: 70,
        },
        zonedOut: {
          name: "zonedOut",
          id: 5,
          isSingleLayer: false,
          layers: {
            line: { zIndex: 70, alpha: 0.5 },
            inside: { zIndex: 60, alpha: 0.7},
          },
        },
        happySmile: {
          name: "happySmile",
          id: 6,
          isSingleLayer: true,
          alpha: 0.7,
          zIndex: 70,
        },
        smallMouth: {
          name: "smallMouth",
          id: 7,
          isSingleLayer: false,
          layers: {
            line: { zIndex: 70, alpha: 0.5 },
            color: { zIndex: 60},
          },
        },
        lipsSmile: {
          name: "lipsSmile",
          id: 8,
          isSingleLayer: false,
          layers: {
            line: { zIndex: 70, alpha: 0.5 },
            color: { zIndex: 60},
          },
        },
      });
      this.choices.shoulders = Object.freeze({
        neutral: {
          name: "neutral",
          id: 0,
          isSingleLayer: false,
          thumb: "line",
          layers: {
            line: { zIndex: 8, alpha: 0.2 },
            color: { zIndex: 7, setOtherLayer: "back" },
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
            choice: this.choices.face.vline,
            layers: {
              line: {},
              color: {
                color: "#d4b485",
              },
            },
          },
          hair: {
            choice: this.choices.hair.shortBob,
            layers: {
              color: { color: "#73bbbf" },
              line: { color: "#cef6f7" },
              backcolor: { color: "#73bbbf" },
              // backline: { color: "#cef6f7" },
            },
          },
          eye: {
            choice: this.choices.eye.neutralFemale,
            layers: {
              line: { color: "#151261" },
              color: { color: "#7278a6" },
              white: {},
              glare: {},
            },
          },
          nose: {
            choice: this.choices.nose.small,
          },
          eyebrows: {
            choice: this.choices.eyebrows.straight,
            color: "#73bbbf",
          },
          mouth: {
            choice: this.choices.mouth.lipsSmile,
            layers: {
              line: {},
              color: {
                color: "#ba6665",
              },
            },
          },
          shoulders: {
            choice: this.choices.shoulders.neutral,
            layers: {
              line: {},
              color: {
                color: "#d4b485",
              },
            },
          },
        },
      };
    },
  },
  mounted() {
    PIXI.settings.ROUND_PIXELS = true;
    this.setupChoices();
    this.setupInitialAvatar();
    this.setupCanvas();
    colorPicker = new iro.ColorPicker("#picker", {
      // Set the size of the color picker
      width: 150,
      // Set the initial color to pure red
      color: "#beebee",
      display: "inline-block",
      borderColor: "#beebee",
      padding: 0,
      borderWidth: 2,
      // handleRadius: 10,
      handleSvg: "#handle",
      handleProps: { x: -7.3, y: -8 },
      layoutDirection: "vertical",
      layout: [
        {
          component: iro.ui.Wheel,
        },
        {
          component: iro.ui.Slider,
        },
        // {
        //   component: iro.ui.Slider,
        //   options: {
        //     sliderType: "saturation",
        //   },
        // },
        // {
        //   component: iro.ui.Slider,
        // options: {
        //   sliderType: "hue",
        // },
        // },
      ],
    });
    const changeLayerColor = this.changeLayerColor;
    colorPicker.on("color:change", function () {
      // console.log("New active color:", color.hexString);
      changeLayerColor();
    });
  },
};
</script>

<style>
.hello {
  border: 2px solid #beebee;
  text-align: center;
  display: inline-block;
  margin: auto;
  font-size: 30px;
  color: #77a2bc;
  padding: 10px;
  text-shadow: 5 5 30px white;
  -webkit-background-clip: text;
  -moz-background-clip: text;
  background-clip: text;
  -webkit-text-fill-color: transparent;
  -moz-text-fill-color: transparent;
  background-image: linear-gradient(
    270deg,
    #646292,
    #beebee,
    white,
    #beebee,
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
  grid-template-areas:
    "l header header header r"
    "l avatar colorpicker featurebox r"
    "l zIndexbox zIndexbox zIndexbox r";
  grid-template-rows: 100px 430px 150px;
  grid-template-columns: 1fr auto minmax(auto, 170px) minmax(auto, 450px) 1fr;
  grid-gap: 10px;
  align-content: center;
}
.creator {
  border: 2px solid #beebee;
  /* margin: auto 100px auto; */
  /* padding: 4px; */
  width: 424px;
  grid-area: avatar;
  background-color: rgba(185, 255, 228, 0.281);
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
  border: 2px solid #beebee;
  background-color: rgba(185, 255, 228, 0.281);
}
.select-color {
  /* border: 2px dashed white; */
  /* margin: auto auto 100px; */
  /* padding: 4px; */
  grid-area: random;
}
.download-avatar {
  order: 0;
}
.save-button {
  /* margin: 10px auto; */
  border-color: #beebee;
}
.color-picker {
  border: 2px solid #beebee;
  background-color: rgba(185, 255, 228, 0.281);
  grid-area: colorpicker;
  display: flex;
  flex-direction: column;
  height: fit-content;
  /* justify-content: space-between; */
}
#picker {
  /* margin: 0px 10px; */
  padding: 7px;
  order: 0;
  /* background-color: #FAF6ED; */
  /* padding: 3px 3px; */
}
.select-dropdown {
  margin: 3px;
  padding: 3px;
  border: 2px solid #beebee;
  order: 7;
  /* border:none; */
  /* text-align-last: center; */
  /* text-align: center; */
}
.feature-select {
  margin: 10px;
  padding: 10px;
  /* display:inline; */
  border: 2px solid #beebee;
  color: #728ca7;
  background: rgba(240, 252, 255, 0.596);
  font-size: 13px;
  min-width: 142px;
  line-height: 1.4;
}
.feature-select-buttons {
  display: inline;
  margin: 10px;
  padding: 5px;
  text-align: left;
  overflow-y: auto;
  max-height: 100px;
  border: 2px solid #beebee;
}
.feature-button {
  font-size: 15px;
  padding: 7px;
  margin: 0px 4px 4px 0px;
  background: rgba(0, 23, 124, 0.623);
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
  border: 2px solid #beebee;
  margin: 10px;
  margin-top: 0px;
  background: rgba(240, 252, 255, 0.788);
  max-height: 500px;
}
.thumbnail {
  height: 70px;
  width: 70px;
  border: 2px solid #8aaec2;
  padding: 2px;
  margin-left: 5px;
  object-fit: cover;
  background: #faf6ed;
}
.selected-thumbnail {
  border-color: #646292;
}
.thumbnail:active {
  transform: scale(0.9);
}
.active-button {
  color: #646292;
  background: rgba(240, 252, 255, 0.788);
}
.active-button:hover {
  color: #646292;
}
</style>

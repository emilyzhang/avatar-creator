<template>
  <div class="creator">
    <canvas id="pixi"></canvas>
  </div>
  <!-- <v-input-colorpicker  v-model="color" /> -->
</template>

<script>
import * as PIXI from "pixi.js";
import { ColorOverlayFilter } from "@pixi/filter-color-overlay";
// import InputColorPicker from 'vue-native-color-picker';


export default {
  name: "Creator",
  props: {
    msg: String,
    selected: Object,
  },
  components: {
    // "v-input-colorpicker": InputColorPicker
  },
  methods: {
    drawPixi() {
      var canvas = document.getElementById("pixi");

      const app = new PIXI.Application({
        width: 500,
        height: 1000,
        antialias: true,
        transparent: true,
        view: canvas,
      });

      const hair = new PIXI.Sprite.from("./assets/hair1.png");
      const face = new PIXI.Sprite.from("./assets/face1.png");
      face.filters = [new ColorOverlayFilter([0.2, 0.3, 0.1])];
      app.stage.addChild(face);
      app.stage.addChild(hair);
    },
    setup() {
    }
  },

  mounted() {
    this.drawPixi();
  },
};
</script>

<style>
.creator {
  font-size: 50px;
  color: #77a2bc;
  border: 2px dashed #77a2bc;
}
.creator:hover {
  border: 2px dashed white;
  text-shadow: 0 0 30px #728ca7;
  color: white;
}
</style>

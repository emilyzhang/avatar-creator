<template>
  <div class="creator">
    <canvas id="pixi"></canvas>
  </div>
</template>

<script>
import * as PIXI from "pixi.js";
import { ColorOverlayFilter } from "@pixi/filter-color-overlay";


export default {
  name: "Creator",
  components: {
  },
  methods: {
    drawPixi() {
      var canvas = document.getElementById("pixi");
      console.log(canvas)

      const app = new PIXI.Application({
        width: 444,
        height: 444,
        antialias: true,
        // transparent: true,
        backgroundColor: "0xffffff",
        view: canvas,
      });
      const hair = new PIXI.Sprite.from("../assets/hair1.png");
      const nose = new PIXI.Sprite.from("../assets/nose0.png");
      const mouth = new PIXI.Sprite.from("../assets/mouth0.png");
      const eyeTop = new PIXI.Sprite.from("../assets/eye1top.png");
      const eyeColor = new PIXI.Sprite.from("../assets/eye1color.png");
      const eyeLight = new PIXI.Sprite.from("../assets/eye1light.png");
      const eyebrows = new PIXI.Sprite.from("../assets/eyebrows0.png");
      const hairColor = new PIXI.Sprite.from("../assets/hair1color.png");
      const face = new PIXI.Sprite.from("../assets/face1.png");
      face.filters = [new ColorOverlayFilter([1, 0.8, 0.7])];
      eyebrows.filters = [new ColorOverlayFilter([0, 0.8, 0.83])];
      hairColor.filters = [new ColorOverlayFilter([0, 0.8, 0.83])];
      eyeColor.filters = [new ColorOverlayFilter([0, 0.18, 0.3])];
      app.stage.addChild(face);
      app.stage.addChild(eyebrows);
      app.stage.addChild(hairColor);
      app.stage.addChild(hair);
      app.stage.addChild(eyeColor);
      app.stage.addChild(eyeTop);
      app.stage.addChild(eyeLight);
      app.stage.addChild(mouth);
      app.stage.addChild(nose);
    },
  },

  mounted() {
    this.drawPixi()
  },
};
</script>

<style>
.creator {
  border: 2px dashed white;
  margin: auto;
  padding: 4px;
  width: 450px;
}
</style>

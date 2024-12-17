<script setup>
import { onMounted, ref } from "vue";
import anime from "animejs";

const svg = ref(null);

onMounted(() => {
  anime({
    targets: svg.value.querySelector("path"),
    strokeDashoffset: [anime.setDashoffset, 0],
    easing: "easeInOutSine",
    duration: 2000,
  });
});
</script>

<template>
<div class="playground">
  <svg ref="svg" width="200" height="200">
    <path d="M10 80 C40 10, 65 10, 95 80 S150 150, 180 80" fill="none" stroke="blue" stroke-width="4" />
  </svg>
</div>
</template>

<style scoped>
.playground {
  width: 100%;
  height: 100%;
}
</style>
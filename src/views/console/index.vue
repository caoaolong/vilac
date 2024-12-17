<script setup>
import { ref, onMounted, onBeforeUnmount } from "vue";
import { Terminal } from '@xterm/xterm';
import { FitAddon } from '@xterm/addon-fit';
import "@xterm/xterm/css/xterm.css";

const terminalContainer = ref(null);
const terminal = new Terminal();
const fitAddon = new FitAddon();
const resizeObserver = null;

onMounted(() => {
  terminal.loadAddon(fitAddon);
  terminal.open(terminalContainer.value);
  fitAddon.fit();
  terminal.onData((data) => {
    terminal.write(`vilac> ${data}\r\n`);
  });
  const resizeObserver = new ResizeObserver((e) => {
    console.log(e)
    fitAddon.fit();
  });
  resizeObserver.observe(terminalContainer.value);
});

onBeforeUnmount(() => {
  resizeObserver.disconnect();
  terminal.dispose();
});
</script>

<template>
  <div ref="terminalContainer" class="terminal"></div>
</template>

<style scoped>
.terminal {
  width: 100%;
  height: 100%;
  text-align: left;
}
</style>
<script setup>
import {ref} from 'vue'
const code = ref(`
void main(string[] args)
{
    printf("%s\\n", "Hello,World!");
}
`)

const codeText = ref({
  lineHeight: '24px',
  fontSize: '16px',
})
</script>

<template>
  <n-layout-content>
    <div class="breakpoint"/>
    <n-code :code="code" language="c" show-line-numbers :style="codeText"/>
  </n-layout-content>
</template>

<style scoped>
* {
  text-align: left;
}
.breakpoint {
  background-color: #FCFC4444;
  width: 100%;
  height: 24px;
  position: absolute;
  top: 24px;
}
</style>
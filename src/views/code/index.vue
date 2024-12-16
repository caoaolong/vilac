<script setup>
import {ref} from 'vue'
const code = ref(`
void main(string[] args)
{
    printf("%s\\n", "Hello,World!");
}
`)
</script>

<template>
  <n-layout-content>
    <n-code :code="code" language="c" show-line-numbers/>
  </n-layout-content>
</template>

<style scoped>
* {
  text-align: left;
}
</style>
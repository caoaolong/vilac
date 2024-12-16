<script setup>
import { darkTheme } from 'naive-ui'
import hljs from 'highlight.js/lib/core'
import c from 'highlight.js/lib/languages/c'

hljs.registerLanguage('c', c)
</script>

<template>
  <n-config-provider :theme="darkTheme" :hljs="hljs">
    <router-view/>
  </n-config-provider>
</template>

<style scoped>
</style>

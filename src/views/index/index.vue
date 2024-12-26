<script setup>
import {ref} from 'vue'
import PlaygroundView from '@/views/playground/index.vue'
import ConsoleView from '@/views/console/index.vue'
import CodeView from '@/views/code/index.vue'
import DockerHeader from '@/components/DockHeader/index.vue'
import Toolbar from '@/components/Toolbar/index.vue'
import {Terminal, Chalkboard} from '@vicons/fa'

const horizontalSize = ref(0.3)
const verticalSize = ref(0.5)

const codeView = ref()

const stepInto = () => {
  if (!codeView.value) return
  codeView.value.stepInto()
}

const stepOut = () => {
  if (!codeView.value) return
  codeView.value.stepOut()
}

const stepPass = () => {
  if (!codeView.value) return
  codeView.value.stepPass()
}
</script>

<template>
  <n-layout has-sider class="full-view">

    <toolbar
        @show-code="verticalSize = 0" @hide-code="verticalSize = 0.5"
        @show-console="verticalSize = 0.5" @hide-console="verticalSize = 1"
        @show-playground="horizontalSize = 0.3" @hidePlayground="horizontalSize = 1"
        @step-into="stepInto" @step-out="stepOut" @step-pass="stepPass"/>

    <n-split direction="horizontal" v-model:size="horizontalSize">
      <template #1>
        <n-split direction="vertical" v-model:size="verticalSize">
          <template #1>
            <code-view ref="codeView"/>
          </template>
          <template #2>
            <docker-header title="Console" :icon="Terminal"/>
            <console-view/>
          </template>
        </n-split>
      </template>
      <template #2>
        <n-layout>
          <docker-header title="Playground" :icon="Chalkboard"/>
          <playground-view/>
        </n-layout>
      </template>
    </n-split>
  </n-layout>
</template>

<style scoped>
.full-view {
  width: 100vw;
  height: 100vh;
  overflow: hidden;
}
</style>
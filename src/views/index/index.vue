<script setup>
import PlaygroundView from '@/views/playground/index.vue'
import ConsoleView from '@/views/console/index.vue'
import CodeView from '@/views/code/index.vue'
import DockerHeader from '@/components/DockHeader/index.vue'
import {Code, Terminal, Chalkboard} from '@vicons/fa'
</script>

<template>
  <n-layout has-sider class="full-view">
    <n-split direction="horizontal" :default-size="0.3">
      <template #1>
        <n-split direction="vertical">
          <template #1>
            <n-layout>
              <docker-header title="Code" :icon="Code"/>
              <code-view/>
            </n-layout>
          </template>
          <template #2>
            <n-layout>
              <docker-header title="Console" :icon="Terminal"/>
              <n-layout-content>
                <console-view/>
              </n-layout-content>
            </n-layout>
          </template>
        </n-split>
      </template>
      <template #2>
        <n-layout>
          <docker-header title="Playground" :icon="Chalkboard"/>
          <n-layout-content>
            <playground-view/>
          </n-layout-content>
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
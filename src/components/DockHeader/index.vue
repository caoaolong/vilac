<script setup>
import { Close } from '@vicons/ionicons5'
defineProps({
  title: null,
  icon: null,
  toolbars: Array
})
const emits = defineEmits(['close', 'toolbar'])
</script>

<template>
  <n-layout-header class="dock-header">
    <div class="dock-header-title dock-header-left">
      <n-icon :size="18" :component="icon"/>
      <h4>{{ title }}</h4>
    </div>
    <div class="dock-header-title dock-header-right">
      <n-popover v-for="item in toolbars" :delay="500">
        <template #trigger>
          <n-icon  :size="16" :component="item.icon" @click="emits('toolbar', item.action)"/>
        </template>
        <span>{{ item.label }}</span>
      </n-popover>
      <n-divider vertical/>
      <n-icon :size="18" :component="Close" @click="emits('close')"/>
    </div>
  </n-layout-header>
</template>

<style scoped>
.dock-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}
.dock-header-title {
  width: 49%;
  display: flex;
  flex-direction: row;
  justify-content: left;
  align-items: center;
  height: 32px;
  padding-left: 10px;
  padding-right: 10px;
  cursor: pointer;
  gap: 8px;
}

.dock-header-left {
  justify-content: left;
}

.dock-header-right {
  justify-content: right;
}
</style>
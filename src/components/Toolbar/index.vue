<template>
  <div class="draggable" :style="style">
    <!-- 仅上边框可拖动 -->
    <div class="drag-handle" @mousedown.stop="startDrag">
      Toolbar
    </div>
    <!-- 其他内容 -->
    <div class="content">
      <n-space>
        <n-popover v-for="button in breakButtons" :delay="500">
          <template #trigger>
            <n-button size="small" quaternary circle>
              <template #icon>
                <n-icon :component="button.icon" />
              </template>
            </n-button>
          </template>
          <span>{{button.tooltip}}</span>
        </n-popover>
        <n-divider vertical />
        <n-popover v-for="button in programButtons" :delay="500">
          <template #trigger>
            <n-button size="small" quaternary circle :type="button.type">
              <template #icon>
                <n-icon :component="button.icon" />
              </template>
            </n-button>
          </template>
          <span>{{button.tooltip}}</span>
        </n-popover>
        <n-divider vertical />
        <n-popover v-for="button in viewButtons" :delay="500">
          <template #trigger>
            <n-button size="small" quaternary circle :type="button.type" @click="button.click(button)">
              <template #icon>
                <n-icon :component="button.icon" />
              </template>
            </n-button>
          </template>
          <span>{{button.tooltip}}</span>
        </n-popover>
      </n-space>
    </div>
  </div>
</template>

<script setup>
import {ref} from 'vue';
import {
  ArrowImport20Filled,
  ArrowDownload24Filled,
  ArrowExportUp24Filled,
  Play24Filled,
  BorderOutsideThick24Filled,
  ArrowClockwise24Filled
} from '@vicons/fluent'

import {Code, Terminal, Chalkboard} from '@vicons/fa'

const emits = defineEmits(['showCode', 'hideCode', 'showConsole', 'hideConsole', 'showPlayground', 'hidePlayground'])

const breakButtons = [
  {
    tooltip: "步过",
    icon: ArrowImport20Filled
  },
  {
    tooltip: "步入",
    icon: ArrowDownload24Filled
  },
  {
    tooltip: "步出",
    icon: ArrowExportUp24Filled
  }
]

const programButtons = [
  {
    type: "success",
    tooltip: "恢复程序",
    icon: Play24Filled
  },
  {
    type: "error",
    tooltip: "结束调试",
    icon: BorderOutsideThick24Filled
  },
  {
    type: "warning",
    tooltip: "重新运行",
    icon: ArrowClockwise24Filled
  }
]

const viewButtons = [
  {
    type: "primary",
    tooltip: "Code View",
    icon: Code,
    show: true,
    click: (button) => {
      button.show = !button.show
      if (button.show) {
        emits('hideCode')
      } else {
        emits('showCode')
      }
    }
  },
  {
    type: "primary",
    tooltip: "Console View",
    icon: Terminal,
    show: true,
    click: (button) => {
      button.show = !button.show
      if (button.show) {
        emits('showConsole')
      } else {
        emits('hideConsole')
      }
    }
  },
  {
    type: "primary",
    tooltip: "Playground View",
    icon: Chalkboard,
    show: true,
    click: (button) => {
      button.show = !button.show
      if (button.show) {
        emits('showPlayground')
      } else {
        emits('hidePlayground')
      }
    }
  }
]

const style = ref({
  position: 'absolute',
  left: '100px',
  top: '100px',
  zIndex: 10
});

let isDragging = ref(false);
let offset = {x: 0, y: 0};

const startDrag = (e) => {
  isDragging.value = true;
  offset.x = e.clientX - parseInt(style.value.left);
  offset.y = e.clientY - parseInt(style.value.top);

  // 监听 mousemove 和 mouseup
  window.addEventListener('mousemove', onMouseMove);
  window.addEventListener('mouseup', onMouseUp);
};

const onMouseMove = (e) => {
  if (!isDragging.value) return;

  // 计算新的位置
  let newLeft = e.clientX - offset.x;
  let newTop = e.clientY - offset.y;

  // 获取窗口宽度和高度
  const windowWidth = window.innerWidth;
  const windowHeight = window.innerHeight;

  // 获取div的宽度和高度
  const divWidth = 455; // 与样式中的宽度一致
  const divHeight = 70; // 与样式中的高度一致

  // 边界检查
  if (newLeft < 0) newLeft = 0;
  if (newTop < 0) newTop = 0;
  if (newLeft + divWidth > windowWidth) newLeft = windowWidth - divWidth;
  if (newTop + divHeight > windowHeight) newTop = windowHeight - divHeight;

  // 更新样式
  style.value.left = `${newLeft}px`;
  style.value.top = `${newTop}px`;
};

const onMouseUp = () => {
  isDragging.value = false;

  // 移除 mousemove 和 mouseup 事件监听
  window.removeEventListener('mousemove', onMouseMove);
  window.removeEventListener('mouseup', onMouseUp);
};
</script>

<style scoped>
.draggable {
  width: 450px;
  height: 65px;
  border: 1px solid #ccc;
  background-color: #101014;
  position: relative;
  border-radius: 4px;
}

.drag-handle {
  width: 100%;
  height: 20px; /* 上边框的高度 */
  background-color: #151515;
  display: flex;
  justify-content: center;
  align-items: center;
  cursor: grab;
  border-top-left-radius: 4px;
  border-top-right-radius: 4px;
  user-select: none;
}

.content {
  padding: 10px;
  color: white;
  user-select: none; /* 禁止选择 */
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: center;
}
</style>

<script setup>
import {ref, shallowRef} from 'vue'
import {Code, File} from "@vicons/fa";
import { useMessage } from 'naive-ui'
import DockerHeader from "@/components/DockHeader/index.vue";

const message = useMessage()
const fileInput = ref(null)

const debugging = ref(false)

const vm = ref({
  debug: {
    index: 0
  },
  code: {
    lineHeight: 24,
    lineCount: 0
  }
})

const toolbars = shallowRef([
  {
    label: 'Open',
    icon: File,
    action: 'open'
  }
])

const codeContent = ref(null)

const codeText = ref({
  lineHeight: `${vm.value.code.lineHeight}px`,
  fontSize: '16px',
})

const lineHighLight = ref({
  top: `${vm.value.debug.index}px`
})

const handleOpenFile = () => {
  fileInput.value.click()
}

const handleToolbar = (event) => {
  if (event === 'open') {
    handleOpenFile()
  }
}

const handleFileChange = (event) => {
  let file = event.target.files[0]
  if (!file) return;

  if (!file.name.endsWith('.vl') && !file.name.endsWith('.VL')) {
    message.error('Format error!')
  }

  const reader = new FileReader()
  reader.onload = function(event) {
    const fileContent = event.target.result;
    try {
      const jsonData = JSON.parse(fileContent)
      vm.value.code.lineCount = jsonData.lines.length
      for (const item of jsonData.lines) {
        codeContent.value += item.value + '\n'
      }
      debugging.value = true
    } catch (error) {
      message.error(error)
    }
  };
  reader.onerror = () => message.error(event.target.error)
  reader.readAsText(file)
}

const stepInto = () => {
  if (!codeContent.value) return
  console.log('a')
}

const stepOut = () => {
  if (!codeContent.value) return
  console.log('bv')
}

const stepPass = () => {
  if (!codeContent.value) {
    message.info('先打开文件!')
    return;
  }
  if (!debugging.value) {
    message.info('已经结束了!')
    return;
  }
  const {debug, code} = vm.value
  debug.index ++
  if (debug.index >= code.lineCount) {
    debugging.value = false
  }
  lineHighLight.value.top = `${debug.index * code.lineHeight}px`
}

defineExpose({
  stepInto, stepOut, stepPass,
})
</script>

<template>
  <input type="file" style="display: none;" @change="handleFileChange" ref="fileInput"/>
  <docker-header title="Code" :toolbars="toolbars" :icon="Code" @toolbar="handleToolbar"/>
  <n-layout-content>
    <div v-if="debugging" class="breakpoint" :style="lineHighLight"/>
    <n-code v-if="codeContent" :code="codeContent" language="c" show-line-numbers :style="codeText"/>
    <n-empty v-else description="请选择打开vl文件" style="margin-top: 25%;">
      <template #icon>
        <n-icon :component="Code">
        </n-icon>
      </template>
    </n-empty>
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
  top: 0;
}
</style>
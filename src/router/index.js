import {createMemoryHistory, createRouter} from 'vue-router'
import Index from '../views/index/index.vue'

const routes = [
    {
        path: '/',
        name: 'Index',
        component: Index,
        children: []
    }
]

const router = createRouter({history: createMemoryHistory(), routes})

export default router
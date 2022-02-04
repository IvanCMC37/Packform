import {createRouter, createWebHistory} from 'vue-router'

import Orders from '../views/Orders'
import Home from '../views/Home'

const routes = [
    {
        path : '/',
        name: 'Home',
        component: Home,
    }
    ,{
        path : '/orders',
        name: 'Orders',
        component: Orders,
    },
]

const router = createRouter({
    history: createWebHistory(process.env.BASE_URL),
    routes,
})

export default router
import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: '/',
      component: () => import('../views/Home.vue'),
    },
    {
      path: '/agriculture',
      component: () => import('../views/Agriculture.vue'),
    },
    {
      path: '/inspection',
      component: () => import('../views/Inspection.vue'),
    },
    {
      path: '/supply-chain',
      component: () => import('../views/SupplyChain.vue'),
    },
    {
      path: '/blockchain-viz',
      component: () => import('../views/BlockchainViz.vue'),
    },
  ],
})

export default router 
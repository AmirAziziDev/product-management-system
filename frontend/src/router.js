import { createRouter, createWebHistory } from 'vue-router'
import Products from './components/Products.vue'

const routes = [
  {
    // Root path redirects to products - just for demo purposes
    path: '/',
    redirect: '/product'
  },
  {
    path: '/product',
    name: 'Products',
    component: Products
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router

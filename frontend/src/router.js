import {createRouter, createWebHistory} from 'vue-router'
import ProductsView from './views/ProductView.vue'

const routes = [
  {
    // Root path redirects to products - just for demo purposes
    path: '/',
    redirect: '/product'
  },
  {
    path: '/product',
    name: 'Products',
    component: ProductsView
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router

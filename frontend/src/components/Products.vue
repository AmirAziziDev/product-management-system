<script setup>
import { ref, onMounted } from 'vue'
import axios from '../../api/axios.js'

const products = ref([])
const loading = ref(false)
const currentPage = ref(1)
const pageSize = ref(20)
const totalItems = ref(0)

const headers = [
  { title: 'ID', align: 'start', sortable: false, key: 'id' },
  { title: 'Code', align: 'start', sortable: false, key: 'code' },
  { title: 'Name', align: 'start', sortable: false, key: 'name' },
  { title: 'Description', align: 'start', sortable: false, key: 'description' },
  { title: 'Created At', align: 'start', sortable: false, key: 'created_at' }
]

const fetchProducts = async (page = 1) => {
  loading.value = true
  try {
    const response = await axios.get(`/api/v1/products?page=${page}&page_size=${pageSize.value}`)
    products.value = response.data.data || []
    totalItems.value = response.data.meta?.total || 0
    currentPage.value = response.data.meta?.page || page
  } catch (error) {
    console.error('Error fetching products:', error)
    products.value = []
    totalItems.value = 0
  } finally {
    loading.value = false
  }
}

const handlePageChange = (page) => {
  currentPage.value = page
  fetchProducts(page)
}

const formatDate = (dateString) => {
  if (!dateString) return ''
  return new Date(dateString).toLocaleDateString('en-US', {
    year: 'numeric',
    month: 'short',
    day: 'numeric',
    hour: '2-digit',
    minute: '2-digit'
  })
}


onMounted(() => {
  fetchProducts()
})
</script>

<template>
  <v-container>
    <v-row>
      <v-col>
        <h1 class="text-h4 mb-4">Products</h1>

        <v-card>
          <v-card-title>
            <v-row align="center" justify="space-between">
              <v-col>
                <span class="text-h6">Product List</span>
              </v-col>
            </v-row>
          </v-card-title>

          <v-data-table
            :headers="headers"
            :items="products"
            :loading="loading"
            :items-per-page="pageSize"
            :server-items-length="totalItems"
            :page="currentPage"
            @update:page="handlePageChange"
            class="elevation-1"
          >
            <template v-slot:item.description="{ item }">
              <span v-if="item.description">{{ item.description }}</span>
              <span v-else class="text-grey">—</span>
            </template>

            <template v-slot:item.created_at="{ item }">
              {{ formatDate(item.created_at) }}
            </template>

            <template v-slot:no-data>
              <v-alert
                :value="!loading && products.length === 0"
                color="info"
                icon="mdi-information"
                class="ma-4"
              >
                No products found.
              </v-alert>
            </template>
          </v-data-table>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

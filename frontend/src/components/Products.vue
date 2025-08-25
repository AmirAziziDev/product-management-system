<script setup>
import { ref, onMounted } from 'vue'
import axios from '../../api/axios.js'

const products = ref([])
const loading = ref(false)
const totalItems = ref(0)
const page = ref(1)
const pageSize = ref(20)

const headers = [
  { title: 'SKU', align: 'start', sortable: false, key: 'sku' },
  { title: 'Name', align: 'start', sortable: false, key: 'name' },
  { title: 'Product Type', align: 'start', sortable: false, key: 'product_type_name' },
  { title: 'Description', align: 'start', sortable: false, key: 'description' },
  { title: 'Created At', align: 'start', sortable: false, key: 'created_at' },
]

async function fetchProducts(p = page.value, ipp = pageSize.value) {
  loading.value = true
  try {
    const { data } = await axios.get('/api/v1/products', {
      params: {
        page: p,
        page_size: ipp,
      },
    })
    products.value = data?.data ?? []
    totalItems.value = data?.meta?.total ?? 0
    page.value = data?.meta?.page ?? p
  } catch (err) {
    console.error('Error fetching products:', err)
    products.value = []
    totalItems.value = 0
  } finally {
    loading.value = false
  }
}

function onPageChange(newPage) {
  page.value = newPage
  fetchProducts(newPage, pageSize.value)
}

function onItemsPerPageChange(newIpp) {
  pageSize.value = newIpp
  // reset to first page when page size changes (common UX)
  page.value = 1
  fetchProducts(1, newIpp)
}

function formatDate(dateString) {
  if (!dateString) return ''
  return new Date(dateString).toLocaleDateString('sv-SE', {
    year: 'numeric',
    month: '2-digit',
    day: 'numeric',
    hour: '2-digit',
    minute: '2-digit',
    second: '2-digit',
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
            <span class="text-h6">Product List</span>
          </v-card-title>

          <v-data-table-server
            :headers="headers"
            :items="products"
            :loading="loading"
            :items-length="totalItems"
            :page="page"
            :items-per-page="pageSize"
            :items-per-page-options="[10, 20, 50, 100]"
            @update:page="onPageChange"
            @update:items-per-page="onItemsPerPageChange"
            class="elevation-1"
          >
            <template #item.sku="{ item }">
              {{ item.product_type?.code }}.{{ item.code }}
            </template>

          <template #item.product_type_name="{ item }">
            {{ item.product_type?.name ?? '—' }}
          </template>

          <template #item.description="{ item }">
            <span v-if="item.description">{{ item.description }}</span>
            <span v-else class="text-grey">—</span>
          </template>

          <template #item.created_at="{ item }">
            {{ formatDate(item.created_at) }}
          </template>

          <template #no-data>
            <v-alert
              :value="!loading && products.length === 0"
              color="info"
              icon="mdi-information"
              class="ma-4"
            >
              No products found.
            </v-alert>
          </template>
          </v-data-table-server>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

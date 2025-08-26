<script setup>
import { ref, onMounted } from 'vue'
import { fetchProducts as fetchProductsAPI } from '@/api/products'

const products = ref([])
const loading = ref(false)
const totalItems = ref(0)
const page = ref(1)
const pageSize = ref(20)

const headers = [
  { title: 'SKU', align: 'start', sortable: false, key: 'sku', width: '100px' },
  { title: 'Name', align: 'start', sortable: false, key: 'name', width: '250px' },
  { title: 'Product Type', align: 'start', sortable: false, key: 'product_type_name', width: '200px' },
  { title: 'Colors', align: 'start', sortable: false, key: 'colors', width: '200px' },
  { title: 'Description', align: 'start', sortable: false, key: 'description', width: '300px' },
  { title: 'Created At', align: 'start', sortable: false, key: 'created_at', width: '200px' },
]

async function fetchProducts(p = page.value, ipp = pageSize.value) {
  loading.value = true
  try {
    const result = await fetchProductsAPI(p, ipp)
    products.value = result.products
    totalItems.value = result.meta.total
    page.value = result.meta.page
  } catch (err) {
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
  page.value = 1
  fetchProducts(1, newIpp)
}

function formatDate(dateString) {
  if (!dateString) return ''
  return new Date(dateString).toLocaleDateString('sv-SE', {
    year: 'numeric', month: '2-digit', day: 'numeric',
    hour: '2-digit', minute: '2-digit', second: '2-digit',
  })
}

onMounted(fetchProducts)
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
            <!-- SKU -->
            <template #item.sku="{ item }">
              {{ item.product_type?.code }}.{{ item.code }}
            </template>

            <!-- Product type name -->
            <template #item.product_type_name="{ item }">
              {{ item.product_type?.name ?? '—' }}
            </template>

            <!-- Colors (comma-separated) -->
            <template #item.colors="{ item }">
              <span v-if="item.colors?.length">
                {{ item.colors.map(c => c.name).join(', ') }}
              </span>
              <span v-else class="text-grey">—</span>
            </template>

            <!-- Description -->
            <template #item.description="{ item }">
              <span v-if="item.description">{{ item.description }}</span>
              <span v-else class="text-grey">—</span>
            </template>

            <!-- Created at -->
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

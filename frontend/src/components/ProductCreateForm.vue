<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { fetchProductTypes, fetchColors, createProduct } from '@/api/products'

const router = useRouter()

const form = ref({
  code: '',
  name: '',
  description: '',
  product_type_id: null,
  color_ids: []
})

const productTypes = ref([])
const colors = ref([])
const loading = ref(false)
const submitting = ref(false)
const showSuccess = ref(false)
const successMessage = ref('')

const rules = {
  required: value => !!value || 'This field is required',
  number: value => !isNaN(value) && value !== '' || 'Must be a number'
}

const isFormValid = computed(() => {
  return form.value.code && 
         form.value.name && 
         form.value.product_type_id && 
         form.value.color_ids.length > 0
})

async function loadProductTypes() {
  loading.value = true
  try {
    productTypes.value = await fetchProductTypes()
  } catch (error) {
    console.error('Failed to load product types:', error)
    productTypes.value = []
  } finally {
    loading.value = false
  }
}

async function loadColors() {
  try {
    colors.value = await fetchColors()
  } catch (error) {
    console.error('Failed to load colors:', error)
    colors.value = []
  }
}

async function handleSubmit() {
  if (!isFormValid.value) return
  
  submitting.value = true
  try {
    await createProduct(form.value)
    successMessage.value = 'Product created successfully!'
    showSuccess.value = true
    
    // Redirect to products page after showing success message
    setTimeout(() => {
      router.push('/product')
    }, 1500)
  } catch (error) {
    console.error('Failed to create product:', error)
  } finally {
    submitting.value = false
  }
}

onMounted(() => {
  loadProductTypes()
  loadColors()
})
</script>

<template>
  <v-card>
    <v-card-title>
      <span class="text-h5">Create New Product</span>
    </v-card-title>
    
    <v-card-text>
      <v-form @submit.prevent="handleSubmit">
        <v-row>
          <v-col cols="12">
            <v-text-field
              v-model="form.code"
              label="Product Code"
              type="number"
              :rules="[rules.required, rules.number]"
              variant="outlined"
              density="compact"
            />
          </v-col>
          
          <v-col cols="12">
            <v-text-field
              v-model="form.name"
              label="Product Name"
              :rules="[rules.required]"
              variant="outlined"
              density="compact"
            />
          </v-col>
          
          <v-col cols="12">
            <v-textarea
              v-model="form.description"
              label="Description (Optional)"
              variant="outlined"
              density="compact"
              rows="3"
            />
          </v-col>
          
          <v-col cols="12">
            <v-select
              v-model="form.product_type_id"
              :items="productTypes"
              item-title="name"
              item-value="id"
              label="Product Type"
              :rules="[rules.required]"
              :loading="loading"
              variant="outlined"
              density="compact"
            />
          </v-col>
          
          <v-col cols="12">
            <v-select
              v-model="form.color_ids"
              :items="colors"
              item-title="name"
              item-value="id"
              label="Colors"
              multiple
              chips
              closable-chips
              :rules="[value => value.length > 0 || 'Select at least one color']"
              variant="outlined"
              density="compact"
            />
          </v-col>
        </v-row>
        
        <v-row class="mt-4">
          <v-col>
            <v-btn
              type="submit"
              color="primary"
              variant="elevated"
              block
              :disabled="!isFormValid || submitting"
              :loading="submitting"
            >
              Create Product
            </v-btn>
          </v-col>
        </v-row>
      </v-form>
    </v-card-text>
  </v-card>

  <v-snackbar
    v-model="showSuccess"
    color="success"
    :timeout="3000"
    top
  >
    {{ successMessage }}
    <template v-slot:actions>
      <v-btn
        variant="text"
        @click="showSuccess = false"
      >
        Close
      </v-btn>
    </template>
  </v-snackbar>
</template>
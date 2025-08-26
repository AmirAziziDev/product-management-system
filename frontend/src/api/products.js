import axios from 'axios'

export async function fetchProducts(page = 1, pageSize = 20) {
  try {
    const { data } = await axios.get('/api/v1/products', {
      params: {
        page,
        page_size: pageSize,
      },
    })

    return {
      products: data?.data ?? [],
      meta: {
        total: data?.meta?.total ?? 0,
        page: data?.meta?.page ?? page,
        pageSize
      }
    }
  } catch (error) {
    console.error('Error fetching products:', error)
    throw error
  }
}

export async function fetchProductTypes() {
  try {
    const { data } = await axios.get('/api/v1/product-types')
    return data?.data ?? []
  } catch (error) {
    console.error('Error fetching product types:', error)
    throw error
  }
}

export async function fetchColors() {
  try {
    const { data } = await axios.get('/api/v1/colors')
    return data?.data ?? []
  } catch (error) {
    console.error('Error fetching colors:', error)
    throw error
  }
}

export async function createProduct(productData) {
  try {
    const payload = {
      code: parseInt(productData.code),
      name: productData.name,
      product_type_id: productData.product_type_id,
      color_ids: productData.color_ids
    }

    if (productData.description && productData.description.trim()) {
      payload.description = productData.description
    }

    const { data } = await axios.post('/api/v1/products', payload)
    return data
  } catch (error) {
    console.error('Error creating product:', error)
    throw error
  }
}

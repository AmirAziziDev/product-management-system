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
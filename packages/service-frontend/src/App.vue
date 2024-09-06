<script setup>
  import axios from 'axios'
  import { onMounted, ref } from 'vue';

  const fetchData = async () => {
    try {
      const response = await axios.get('http://localhost:4000/sample')
      return response.data
    } catch (error) {
      console.error(error)
    }
  }

  const message = ref('')

  onMounted(async() => {
    const result = await fetchData()
    message.value = result.message
  })
</script>

<template>
  <main>
    <div class="wrapper">
      {{  message }}
    </div>
  </main>
</template>

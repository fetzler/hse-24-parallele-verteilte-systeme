<template>
  <h1>Ich bin hier</h1>
  <button @click="add">Klick</button>
  <h1>{{ database }}</h1>

</template>

<script setup>
  import { ref } from 'vue'
  import axios from 'axios'

  const database = ref(null)

  function add() {
    axios.get("http://backend:8080/todos").then(res => { database.value = res.data})
  }
</script>

<style>
/* CSS-Styling hier */
</style>
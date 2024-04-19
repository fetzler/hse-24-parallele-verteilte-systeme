<template>
  <h1>Ich bin hier</h1>
  <button @click="add">Klick</button>
  <button @click="remove">Remove</button>
  <h1>{{ database }}</h1>

</template>

<script setup>
  import { ref } from 'vue'
  import axios from 'axios'

  const database = ref(null)

  function add() {
    axios.get("/todos").then(res => { 
    database.value = res.data;
    }).catch(error => {
      console.error('Error fetching todos:', error);
    });
  }
  function remove() {
    axios.post("/todos/mach").catch(error => {
      console.error('Error posting todos', error);
    })
  }
</script>

<style>
/* CSS-Styling hier */
</style>
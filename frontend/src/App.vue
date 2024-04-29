<template>
  <div>
    <h1>TODO List</h1>
    <div>
      <input type="text" v-model="newTodoText" placeholder="Enter your todo">
      <button @click="addTodo">Add Todo</button>
    </div>
    <ul>
      <li v-for="(todo, index) in todos" :key="todo.id">
        <span>{{ todo.title }}</span>
        <button @click="markAsDone(todo.id)">Done!</button>
      </li>
    </ul>
  </div>
</template>

<script>
import { ref, onMounted } from 'vue';
import axios from 'axios'

export default {
  setup() {
    const newTodoText = ref('');
    const todos = ref([]);

    const addTodo = async () => {
      if (newTodoText.value.trim() !== '') {
        try {
          const response = await fetch(`/todos/${newTodoText.value}`, {
            method: 'POST',
          });
          if (response.ok) {
            getTodos()
            newTodoText.value = '';
          } else {
            console.error('Failed to add todo:', response.statusText);
          }
        } catch (error) {
          console.error('Failed to add todo:', error);
        }
      }
    };

    const markAsDone = async (todoId) => {
      try {
        const response = await fetch(`/todos/${todoId}`, {
          method: 'DELETE'
        });
        if (response.ok) {
          console.log('Todo marked as done:', todoId);
          todos.value = todos.value.filter(todo => todo.id !== todoId);
        } else {
          console.error('Failed to mark todo as done:', response.statusText);
        }
      } catch (error) {
        console.error('Failed to mark todo as done:', error);
      }
    };

    const getTodos = async () => {
      axios.get("/todos").then(res => { 
      todos.value = res.data;
      }).catch(error => {
        console.error('Error fetching todos:', error);
      });
    };

    onMounted(getTodos);

    return { newTodoText, todos, addTodo, markAsDone };
  }
};
</script>

<style scoped>
/* Add your component's styles here */
</style>

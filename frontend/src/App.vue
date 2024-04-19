<template>
  <div>
    <div class="header">
      <h1 class="title">Most beautiful of all Java Todo Lists</h1>
    </div>
    <div class="add-todo">
      <h2>Add a Todo</h2>
      <div class="todo-input">
        <input type="text" v-model="newTodo" placeholder="Enter Todo Description">
        <button @click="addTodo">Submit!</button>
      </div>
    </div>
    <div class="todo-list">
      <h2>Todo List</h2>
      <div v-for="(todo, index) in todos" :key="index" class="todo-item">
        <span class="item">{{ todo.description }}</span>
        <button @click="markAsDone(index)">Done!</button>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      newTodo: '',
      todos: []
    };
  },
  methods: {
    addTodo() {
      if (this.newTodo.trim() !== '') {
        this.todos.push({
          description: this.newTodo,
          done: false
        });
        this.newTodo = '';
      }
    },
    markAsDone(index) {
      this.todos.splice(index,1);
    }
  }
};
</script>

<style>
/* CSS-Styling hier */
</style>

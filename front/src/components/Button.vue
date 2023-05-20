
<script setup lang="ts">
import { computed, withDefaults } from 'vue';

interface ButtonProps {
    color: 'primary' | 'secondary' | 'success';
}

const props = withDefaults(defineProps<ButtonProps>(), {
  color: 'primary'
})

const btnClass = computed(() => `btn-${props.color}`);
</script>

<template>
    <button :class="btnClass" @click="$emit('click')">
      <slot></slot>
    </button>
</template>

<style scoped>
button {
  display: inline-block;
  padding: 0.55rem 1.25rem;
  font-size: 1rem;
  line-height: 1.5;
  text-align: center;
  white-space: nowrap;
  vertical-align: middle;
  user-select: none;
  border: 1px solid transparent;
  border-radius: 0.25rem;
  transition: color 0.15s ease-in-out, background-color 0.15s ease-in-out, border-color 0.15s ease-in-out, box-shadow 0.15s ease-in-out;
}
  .btn-primary {
    background-color: #007bff;
    color: white;
  }
  .btn-primary:hover{
    background-color: #0169d7;
  }
  
  .btn-secondary {
    background-color: rgb(162, 162, 162);
    color: white;
  }
  .btn-secondary:hover {
    background-color: rgb(117, 117, 117);
  }
  
  .btn-success {
    background-color: rgb(1, 180, 1);
    color: white;
  }
  .btn-success:hover {
    background-color: rgb(0, 145, 0);
  }
  </style>
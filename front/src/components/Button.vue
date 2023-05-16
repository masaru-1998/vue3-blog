
<script setup lang="ts">
import { computed, defineProps } from 'vue';

interface ButtonProps {
    color: 'primary' | 'secondary' | 'success';
}
const props = defineProps<ButtonProps>()
const btnClass = computed(() => `btn-${props.color}`);
</script>

<template>
    <button :class="btnClass" @click="$emit('click')">
      <slot></slot>
    </button>
</template>

<style scoped>
  .btn-primary {
    background-color: blue;
    color: white;
  }
  
  .btn-secondary {
    background-color: gray;
    color: white;
  }
  
  .btn-success {
    background-color: green;
    color: white;
  }
  </style>
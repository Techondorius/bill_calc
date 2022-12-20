<template>
  <h1>Hi!</h1>
  <input v-model="a" />
  <p>{{ a }}</p>
</template>

<script setup lang="ts">
const a = ref("Hello Word")
</script>

<style></style>

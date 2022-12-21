<template>
  <h1>Hi!</h1>
  <input v-model="a" />
  <p>{{ a }}</p>
  <button @click="req">asdf</button>
  <p v-if="{resp}">{{ resp }}</p>
</template>

<script setup lang="ts">
const a = ref("Hello Word")
interface res {
  message: string
}
const req = async () => {
  const url = "http://localhost:8080"
  const res = await fetch(url, {
    method: "POST",
  })
  const resData: res = await res.json()
  resp.value = resData.message
}

const resp = ref("")
</script>

<style></style>

<template>
  <h1>Hi!</h1>
  <input v-model="a" />
  <p>{{ a }}</p>
  <button @click="updateReq">asdf</button>
  <p v-if="{ resp }">{{ resp.message }}</p>
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
  return resData
}
const updateReq = async () => {
  resp.value = await req()
}

const resp = ref(await req())
</script>

<style></style>

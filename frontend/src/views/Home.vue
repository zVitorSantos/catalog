<script setup>
import Top from '../components/Top.vue'
import Search from '../components/Search.vue'
import List from '../components/List.vue'
import Footer from '../components/Footer.vue'
</script>

<template>
  <main>
    <Top />
    <Search />
    <List class="list" />
    <Footer />
  </main>
</template>

<style scoped>
.list {
  margin-top: 15%;
}
</style>
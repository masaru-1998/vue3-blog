<script setup lang="ts">

</script>

<template>
    <nav>
        <router-link to="/" class="nav-link">ホーム画面</router-link>
        <router-link to="/article" class="nav-link">記事一覧</router-link>
        <router-link to="/post-article" class="nav-link">記事投稿</router-link>
    </nav>
</template>
<style scoped>
.nav-link {
  margin-right: 10px;
  text-decoration: none;
  color: black;
}

.nav-link.router-link-active {
  color: blue;
}
</style>
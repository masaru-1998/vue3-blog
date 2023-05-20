<script setup lang="ts">
// import any necessary components or libraries here
</script>

<template>
    <nav class="navbar">
        <div class="container">
            <div class="navbar-brand">
                <router-link to="/" class="navbar-item">Your Logo</router-link>
            </div>
            
            <div class="navbar-menu">
                <div class="navbar-end">
                    <router-link to="/" class="navbar-item">ホーム画面</router-link>
                    <router-link to="/article" class="navbar-item">記事一覧</router-link>
                    <router-link to="/post-article" class="navbar-item">記事投稿</router-link>
                </div>
            </div>
        </div>
    </nav>
</template>

<style scoped>
.navbar {
    background-color: #f8f9fa;
    box-shadow: 0 2px 5px rgba(0, 0, 0, 0.15);
}

.container {
    max-width: 1200px;
    margin: auto;
}

.navbar-brand {
    font-size: 1.5em;
}

.navbar-item {
    padding: 10px 15px;
    color: black;
    text-decoration: none;
}

.navbar-item:hover {
    color: #007bff;
}

.navbar-item.router-link-active {
    color: #007bff;
}
</style>
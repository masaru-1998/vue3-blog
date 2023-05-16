<script setup lang="ts">
import useArticleStore from '@/store/article';

const store = useArticleStore();
const articles = store.readArticles();

</script>
<template>
    <h1>記事一覧</h1>
    <main>
        <div v-for="article in articles" :key="article.id">
            <h2>{{ article.title }}</h2>
            <p>{{ article.description }}</p>
            <p>{{ article.author }}</p>
        </div>
    </main>
</template>
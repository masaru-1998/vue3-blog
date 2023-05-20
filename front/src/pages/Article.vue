<script setup lang="ts">
import useArticleStore from '@/store/article';

const store = useArticleStore();
const articles = store.readAllArticles();

</script>
<template>
    <div class="article-container">
        <h1>記事一覧</h1>
        <main>
            <router-link 
                v-for="article in articles" 
                :key="article.id"
                :to="`/article-detail/${article.id}`" 
                class="article-card"
            >
                <h2>{{ article.title }}</h2>
                <p>{{ article.description }}</p>
                <p class="author">{{ article.author.name }}</p>
            </router-link>
        </main>
    </div>
</template>

<style scoped>
.article-container {
    background-color: #f8f9fa;
    max-width: 90vw;
    margin: auto;
    margin-top: 20px;
    padding: 20px;
    border-radius: 10px;
    box-shadow: 0 2px 5px rgba(0, 0, 0, 0.15);
}

.article-container h1 {
    font-size: 2em;
    color: #343a40;
    margin-bottom: 20px;
    text-align: center;
}

.article-card {
    display: block;
    background-color: #fff;
    margin-bottom: 20px;
    padding: 20px;
    border-radius: 5px;
    box-shadow: 0 2px 5px rgba(0, 0, 0, 0.15);
    text-decoration: none;
    color: inherit;
}

.article-card h2 {
    font-size: 1.5em;
    color: #343a40;
    margin-bottom: 10px;
}

.article-card p {
    font-size: 1em;
    color: #6c757d;
    line-height: 1.6;
}

.author {
    font-style: italic;
    color: #6c757d;
    text-align: right;
}
</style>
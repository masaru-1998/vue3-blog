import { createRouter, createWebHistory } from 'vue-router';
import PostArticle from '@/pages/PostArticle.vue'
import Article from '@/pages/Article.vue';
import Home from '@/pages/Home.vue';

const routes = createRouter({
    history: createWebHistory(),
    routes: [
        {
            path: '/',
            name: 'Home',
            component: Home
        },
        {
            path: '/article',
            name: 'Article',
            component: Article
        },
        {
            path: '/post-article',
            name: 'PostArticle',
            component: PostArticle
        }
    ]
})

export default routes

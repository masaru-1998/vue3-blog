import { createRouter, createWebHistory } from 'vue-router';
import PostArticle from '@/pages/PostArticle.vue'
import Article from '@/pages/Article.vue';
import Home from '@/pages/Home.vue';
import Detail from '@/pages/Detail.vue'

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
        },
        {
            path: '/article-detail/:id',
            name: 'ArticleDetail',
            component: Detail
        }
    ]
})

export default routes

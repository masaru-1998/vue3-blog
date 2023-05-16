import { reactive } from "vue";
import { defineStore } from "pinia";
import { Article, Comment, State} from '@/types/article'
  
const useArticleStore = defineStore('article', () => {
    const initState: State = {
        articles: [],
        comments: []
    }

    const store = reactive(initState)

    const readArticles = (): Article[] => {
        return store.articles;
    }

    const readComments = (articleId: number): Comment[] => {
        return store.comments.filter(comment => comment.articleId === articleId);
    }

    const createArticle = (newArticle: Article) => {
        store.articles.push(newArticle);
    }

    const addComment = (articleId: number, newComment: Comment) => {
        newComment.articleId = articleId;
        store.comments.push(newComment);
    }

    const updateArticle = (updatedArticle: Article) => {
        const index = store.articles.findIndex(article => article.id === updatedArticle.id);
        if (index !== -1) {
            store.articles[index] = updatedArticle;
        }
    }

    const deleteArticle = (articleId: number) => {
        const index = store.articles.findIndex(article => article.id === articleId);
        if (index !== -1) {
            store.articles.splice(index, 1);
        }
    }

    const deleteAllArticles = () => {
        store.articles = [];
        store.comments = [];
    }

    return {
        readArticles,
        readComments,
        createArticle,
        addComment,
        updateArticle,
        deleteArticle,
        deleteAllArticles
    }
})
export default useArticleStore;
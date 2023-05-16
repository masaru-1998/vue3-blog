<script setup lang="ts">
    import { reactive, computed } from 'vue';
    import useArticleStore from '@/store/article';
    import { Article } from '@/types/article';
    import Button from '@/components/Button.vue';

    const store = useArticleStore();
    const articles = computed(() => store.readArticles());
    const tagState = reactive({
        availableTags: ['仕事', '趣味', '勉強', '恋愛', '旅行'],
        selectedTags: []
    })

    const postArticle = reactive<Article>({
        id: 0,
        author:{
            name: ''
        },
        title: '',
        description: '',
        body: '',
        tagList:tagState.selectedTags
    })

    const selectedTags = computed({
        get: () => tagState.selectedTags,
        set: (value) => tagState.selectedTags = value
    })

    const addPost = () => {
        const lastArticle = articles.value[articles.value.length - 1];
        const newArticleId = lastArticle ? lastArticle.id + 1 : 1;

        const newArticle: Article = {
            id: newArticleId,
            author: {
                name: postArticle.author.name
            },
            title: postArticle.title,
            description: postArticle.description,
            body: postArticle.body,
            tagList: [...tagState.selectedTags]
        }

        store.createArticle(newArticle)
        postArticle.author.name = '';
        postArticle.title = '';
        postArticle.description = '';
        postArticle.body = '';
        tagState.selectedTags = [];
    }
</script>
<template>
    <section class="form-container">
        <div class="form-wrapper">
            <form @submit.prevent="addPost">
                <div>
                    <label for="authrer">作者</label>
                    <input id="auther" type="text" v-model="postArticle.author.name" placeholder="Write auther ..." />
                </div>
                <div>
                    <label for="title">タイトル</label>
                    <input for="title" type="text" v-model="postArticle.title" placeholder="Write your title ..." />
                </div>
                <div>
                    <label for="description">説明</label>
                    <input id="descriptino" type="text" v-model="postArticle.description" placeholder="Write description of thie article ..." />
                </div>
                <div>
                    <label for="body">本文</label>
                    <textarea id="body"  v-model="postArticle.body" placeholder="what your article ..." />
                </div>
                <h2>タグ選択</h2>
                <div class="tag-container">
                    <div v-for="tag in tagState.availableTags" :key="tag">
                        <input type="checkbox" :id="tag" :value="tag" v-model="selectedTags">
                        <label :for="tag">{{ tag }}</label>
                    </div>
                </div>
                <Button color="primary" type="submit">投稿</Button>
            </form>
        </div>
    </section>
</template>

<style scoped>
.form-container {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 50vh;
}
.form-wrapper {
    width: 350px;
    padding: 20px;
    border: 1px solid #ccc;
    border-radius: 4px;
}
.form-wrapper form {
    display: flex;
    flex-direction: column;
}
.form-wrapper form div {
    margin-bottom: 10px;
}
.form-wrapper div {
    display: flex;
}
.form-wrapper input,
.form-wrapper textarea {
  margin-left: auto;
  margin-right: 30px;
}

.tag-container {
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: space-between;
}
.tag-container h2 {
    margin-bottom: 10px;
}
.tag-container input[type="checkbox"] {
    margin-right: 5px;
}
ul {
    list-style-type: none;
}
</style>
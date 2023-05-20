<script setup lang="ts">
    import { reactive, computed } from 'vue';
    import useArticleStore from '@/store/article';
    import { Article } from '@/types/article';
    import Button from '@/components/Button.vue';

    const store = useArticleStore();
    const articles = computed(() => store.readAllArticles());
    const tagState = reactive({
        availableTags: ['仕事', '趣味', '勉強', '恋愛', '旅行'],
        selectedTags: []
    })
    const text = 'テスト';
    console.log(`今週の課題は,${text}`);

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
    <div class="form-container">
      <form class="form" @submit.prevent="addPost">
        <div class="form-group">
          <label for="author">作者</label>
          <input id="author" type="text" v-model="postArticle.author.name" placeholder="Write author's name..." />
        </div>
  
        <div class="form-group">
          <label for="title">タイトル</label>
          <input id="title" type="text" v-model="postArticle.title" placeholder="Write your title..." />
        </div>
  
        <div class="form-group">
          <label for="description">説明</label>
          <textarea id="description" v-model="postArticle.description" placeholder="Write the description..." />
        </div>
  
        <div class="form-group">
          <label for="body">本文</label>
          <textarea id="body" v-model="postArticle.body" placeholder="Write the article content..." />
        </div>
  
        <div class="form-group">
          <label>タグ選択</label>
          <div class="tag-container">
            <div class="tag" v-for="tag in tagState.availableTags" :key="tag">
              <input type="checkbox" :id="tag" :value="tag" v-model="selectedTags">
              <label :for="tag">{{ tag }}</label>
            </div>
          </div>
        </div>
  
        <Button color="primary" type="submit">投稿</Button>
      </form>
    </div>
</template>

<style scoped>
.form-container {
  width: 100%;
  max-width: 800px;
  margin: 0 auto;
  padding: 2rem;
}

.form {
  width: 100%;
}

.form-group {
  margin-bottom: 1rem;
}

.form-group label {
  display: block;
  margin-bottom: 0.5rem;
}

.form-group input[type="text"],
.form-group textarea {
  width: 100%;
  padding: 0.5rem;
  border: 1px solid #ccc;
  border-radius: 4px;
}

.tag-container {
  display: flex;
  flex-wrap: wrap;
  gap: 0.5rem;
}

.tag {
  display: flex;
  align-items: center;
}

.tag input[type="checkbox"] {
  margin-right: 0.5rem;
}
</style>
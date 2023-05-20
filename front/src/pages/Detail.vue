Copy code
<script setup lang="ts">
import { ref, computed } from 'vue';
import { useRoute } from 'vue-router';
import useArticleStore from '@/store/article';

import Button from '@/components/Button.vue';

// Retrieve the path parameter
const route = useRoute()
const rawId = route.params.id
if (Array.isArray(rawId)) {
  throw new Error('ID should not be an array')
}
const id: string = rawId
const numID: number = Number(rawId)

// Fetch the article associated with the id from the store
const store = useArticleStore();
const articleState = store.readArticle(id);
const comments = computed( () => store.readComments(numID));

const comment = ref('');

const handleComment = () => {
    if(comment.value.length > 0){
        store.addComment(numID, comment.value)
    }
    comment.value = '';
}
</script>

<template>
    <div class="container">
      <div class="article-container">
        <main>
          <div class="article-content">
            <h1>{{ articleState?.title }}</h1>
            <p>{{ articleState?.body }}</p>
          </div>
        </main>
      </div>
      <div class="comment-container">
        <section class="comment-section">
          <h2>コメント</h2>
          <form @submit.prevent="handleComment">
            <input v-model="comment" type="text" placeholder="コメントを追加" required />
            <Button color="success" type="submit">コメント追加</Button>
          </form>
          <div v-for="comment in comments" :key="comment.id">
            <p>{{  comment.comment }}</p>
          </div>
        </section>
      </div>
    </div>
  </template>
  
  <style scoped>
  .container {
    display: flex;
    padding: 30px;
    flex-direction: column;
    align-items: center;
    gap: 2rem;
    margin-top: 20px;
  }
  
  .article-container,
  .comment-container {
    background-color: #f8f9fa;
    max-width: 800px;
    width: 100%;
    padding: 2rem;
    border-radius: 10px;
    box-shadow: 0 2px 5px rgba(0, 0, 0, 0.15);
  }
  
  .article-content {
    padding: 1rem;
  }
  
  .article-content h1 {
    font-size: 2rem;
    margin-bottom: 1rem;
    color: #333;
  }
  
  .article-content p {
    line-height: 1.6;
    color: #333;
  }
  
  .comment-section {
    margin-top: 2rem;
    margin-bottom: 2rem;
  }
  
  .comment-section h2 {
    font-size: 1.5rem;
    margin-bottom: 1rem;
  }
  
  .comment-section form {
    display: flex;
    gap: 1rem;
  }
  
  .comment-section input {
    flex-grow: 1;
  }

  .comment-section div {
    margin-top: 15px;
    padding: 10px;
    border-radius: 10px;
    box-shadow: 0 2px 5px rgba(0, 0, 0, 0.15);
    background-color:#F7F6EB;
  }
  </style>
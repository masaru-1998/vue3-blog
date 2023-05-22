<script setup lang="ts">
import { useField, useForm } from 'vee-validate';
import axios from 'axios';

import Button from '@/components/Button.vue';
import schema from './signup.schema';

const { handleSubmit, errors } = useForm({validationSchema: schema});
const { value: firstName } = useField<string>("firstName")
const { value: lastName } = useField<string>("lastName")
const { value: email } = useField<string>("email")
const { value: password} = useField<string>("password")

const URL = import.meta.env.SERVER_ENDPOINT || 'http://localhost:80';
const onSubmit = handleSubmit(() => {
    const formData = new FormData();
    formData.append('firstName', firstName.value);
    formData.append('lastName', lastName.value);
    formData.append('email', email.value);
    formData.append('password', password.value);

    axios.post(`${URL}/signup`, formData)
    .then( res => console.log(res))
    .catch( err => console.error(err))
});

</script>
<template>
    <div class="container">
        <h1>アカウント作成</h1>
        <div class="form-container">
            <form @submit.prevent="onSubmit">
                <div class="name-field">
                    <div class="single-input">
                        <label>氏名</label>
                        <input type="text" v-model="firstName" placeholder="氏名"/>
                        <p class="error-message">{{ errors.firstName}}</p>
                    </div>
                    <div class="single-input">
                        <label>名前</label>
                        <input type="text" v-model="lastName" placeholder="名前" />
                        <p class="error-message">{{ errors.lastName }}</p>
                    </div>
                </div>
                <div class="input-field">
                    <label>メールアドレス</label>
                    <input type="email" v-model="email" placeholder="メールアドレス" />
                    <p class="error-message">{{ errors.email }}</p>
                </div>
                <div class="input-field">
                    <label>パスワード</label>
                    <input type="password" v-model="password" placeholder="パスワード" />
                    <p class="error-message">{{ errors.password }}</p>
                </div>
                <div class="button-field">
                    <Button>サインアップ</Button>
                </div>
            </form>
        </div>
    </div>
</template>

<style scoped>
.container{
    max-width: 80vw;
    margin: auto;
    margin-top: 40px;
    font-family: 'Artial', sans-serif;
    color: #333;
    background-color: #e0f7fa;
    padding: 20px;
    border-radius: 8px;
}
h1 {
    text-align: center;
    padding-bottom: 20px;
}
.form-container {
    width: 80%;
    margin: auto;
}
.input-field {
    display: flex;
    flex-direction: column;
    margin-bottom: 20px;
}
.name-field {
    display: flex;
    justify-content: space-between;
    margin-bottom: 20px;
}
.error-message{
    color: red;
}
.single-input {
    flex: 1;
    margin-right: 10px;
}
input {
    padding: 10px;
    border: 1px solid #b2ebf2;
    border-radius: 5px;
}
input:focus {
    outline: none;
    border: 3px solid #4dd0e1;
}
.button-field{
    text-align: center;
}
label {
    margin-bottom: 5px;
}
</style>
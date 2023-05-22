import * as yup from 'yup';
 
const schema = yup.object({
    firstName: yup.string().required("氏名の入力は必須です"),
    lastName: yup.string().required("名前の入力は必須です"),
    email: yup.string().required("メールアドレスは必須です").email("メールアドレスの形式に問題があります"),
    password: yup
    .string()
    .required("パスワードは必須です")
    .matches(
      /^(?=.*[A-Za-z])(?=.*\d)[A-Za-z\d]{8,}$/,
      "パスワードは少なくとも8文字以上、英数字を使用してください。"
    )
});
export default schema;
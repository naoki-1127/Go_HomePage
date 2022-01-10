var app = new Vue({
    // 「el」プロパティーで、Vueの表示を反映する場所＝HTML要素のセレクター（id）を定義
    el: '#app',
    data() {
        return {
            message: "Hello Vue.js"
        }
    }
})

var app3 = new Vue({
    // 「el」プロパティーで、Vueの表示を反映する場所＝HTML要素のセレクター（id）を定義
    el: '#app3',
    data() {
        return {
            message: "Hello everyone",
            username:"",
            email:"",
            message:"",
            password:""
        }
    },
    methods:{
        submit(){
            const data = new URLSearchParams();
            data.append("name", this.username);
            data.append("email", this.email);
            data.append("password", this.password);
            
            axios.post('registration', data)
            .then(response=> {
                // 送信成功時の処理
                if(response.data.result=="ユーザの登録が既にあります"){
                    this.message = response.data.result
                    console.log("ok");
                }else{
                    this.message = "登録に成功しました"
                }
            })
            .catch(function (error) {
                // 送信失敗時の処理
                console.log(error);
            });
        }
    }
})
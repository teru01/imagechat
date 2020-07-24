<template>
  <div class="hello">
    <div class="auth">
      <div v-if="!is_login">
        Email: <input type="text" name="email" v-model="email" />
        Password: <input type="password" v-model="password" />
        <button class="btn" @click="login" id="login">login</button>
      </div>
      <p v-if="is_login">Hello, {{login_name}}</p>
      <button v-if="is_login" class="btn" @click="logout" id="logout">logout</button>
    </div>

    <h1>{{ msg }}</h1>

    <div v-if="is_login">
      Title: <input type="text" name="name" v-model="your_name" />
      <input class="form__item" type="file" @change="onFileChange" />
      <output class="form__output" v-if="preview">
        <img height="200px" :src="preview" alt />
      </output>
      <button class="btn" @click="send" id="sendname">send!!</button>
    </div>

    <p v-if="result">{{result}}</p>
    <button v-if="current_page > 0" class="btn" @click="prev" id="pref">prev</button>
    <button v-if="!isLast" class="btn" @click="next" id="next">next</button>
    <table id="posts">
      <tr v-for="item in items" :key="item.ID">
        <td>{{item.ID}}</td>
        <td>{{item.name}}</td>
        <td>{{item.user_name}}</td>
        <td><img height="200px" :src="item.image_url"></td>
      </tr>
    </table>
  </div>
</template>

<script>
import api from "../api/index.js";
const ROW_PER_PAGE = 20;
export default {
  name: "HelloWorld",
  props: {
    msg: String
  },
  data() {
    return {
      your_name: "",
      result: "",
      email: "",
      password: "",
      names: [],
      current_page: 0,
      isLast: false,
      preview: null,
      photo: null,
      items: [],
      is_login: false,
      login_name: "",
    };
  },

  methods: {
    async send() {
      const formData = new FormData();
      formData.append("photo", this.photo);
      formData.append("name", this.your_name);
      const response = await api()
        .post("/posts", formData)
        .catch(err => err.response || err);
      if (response.status !== 201) {
        this.result = "ERROR";
      } else {
        this.result = "success!!";
      }
      this.reload();
    },

    async setLoginStatus() {
      const sess_result = await api().get(`session`).catch(err => err.response || err);
      if (sess_result.status === 200) {
        this.is_login = true;
        this.login_name = sess_result.data.Name
      } else {
        this.is_login = false;
      }
    },

    async login() {
      const response = await api().post(`/session`, {
        'email': this.email,
        'password': this.password
      }).catch(err => err.response || err);
      if (response.status !== 200) {
        this.result = "ERROR";
      } else {
        this.result = "success!!";
      }
      await this.setLoginStatus()
    },

    async logout() {
      const response = await api().delete('/session').catch(err => err.response || err);
      if (response.status !== 200) {
        this.result = "logout ERROR";
      } else {
        this.result = "success!!";
        this.is_login = ''
      }
    },

    async prev() {
      if (this.current_page <= 0) {
        return;
      }
      this.isLast = false;
      this.current_page -= 1;
      const response = await api()
        .get(
          `/posts?offset=${ROW_PER_PAGE *
            this.current_page}&limit=${ROW_PER_PAGE}`
        )
        .catch(err => err.response || err);
      if (response.status !== 200) {
        this.result = "ERROR";
      } else {
        this.result = "success!!";
        this.items = response.data;
      }
    },

    async next() {
      this.current_page += 1;
      const response = await api()
        .get(
          `/posts?offset=${ROW_PER_PAGE *
            this.current_page}&limit=${ROW_PER_PAGE}`
        )
        .catch(err => err.response || err);
      if (response.status !== 200) {
        this.result = "ERROR";
      } else {
        this.result = "success!!";
        this.items = response.data;
        if (response.data.length < ROW_PER_PAGE) {
          this.isLast = true;
        }
      }
    },

    async reload() {
      const response = await api()
        .get("/posts")
        .catch(err => err.response || err);
      if (response.status !== 200) {
        this.result = "ERROR";
      } else {
        this.result = "success!!";
        this.items = response.data;
        if (response.data.length < ROW_PER_PAGE) {
          this.isLast = true;
        }
      }
      await this.setLoginStatus()
    },

    onFileChange(event) {
      if (event.target.files.length === 0) {
        return false;
      }
      if (!event.target.files[0].type.match("image.*")) {
        return false;
      }
      const reader = new FileReader();
      reader.onload = e => {
        this.preview = e.target.result;
      };
      reader.readAsDataURL(event.target.files[0]);
      this.photo = event.target.files[0];
    },

    async submit() {}
  },

  mounted: async function() {
    this.reload();
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
#posts {
    margin: auto;
}
</style>

<template>
  <div class="hello">
    <h1>{{ msg }}</h1>
    <h3>name</h3>
    <input type="text" name="name" v-model="your_name" />
    <input class="form__item" type="file" @change="onFileChange" />
    <output class="form__output" v-if="preview">
      <img height="200px" :src="preview" alt />
    </output>
    <button class="btn" @click="send" id="sendname">send!!</button>
    <p v-if="result">{{result}}</p>
    <button v-if="current_page > 0" class="btn" @click="prev" id="pref">prev</button>
    <button v-if="!isLast" class="btn" @click="next" id="next">next</button>
    <table>
      <tr v-for="item in items" :key="item.ID">
        <td>{{item.ID}}</td>
        <td>{{item.Name}}</td>
        <td><img height="200px" :src="item.ImageUrl"></td>
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
      names: [],
      current_page: 0,
      isLast: false,
      preview: null,
      photo: null,
      items: []
    };
  },

  methods: {
    async send() {
      const formData = new FormData();
      formData.append("photo", this.photo);
      formData.append("name", this.your_name);
      const response = await api()
        .post("/hoges", formData)
        .catch(err => err.response || err);
      if (response.status !== 201) {
        this.result = "ERROR";
      } else {
        this.result = "success!!";
      }
      this.reload();
    },

    async prev() {
      if (this.current_page <= 0) {
        return;
      }
      this.isLast = false;
      this.current_page -= 1;
      const response = await api()
        .get(
          `/hoges?offset=${ROW_PER_PAGE *
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
          `/hoges?offset=${ROW_PER_PAGE *
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
        .get("/hoges")
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
</style>

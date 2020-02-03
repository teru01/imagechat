<template>
  <div class="hello">
    <h1>{{ msg }}</h1>
    <h3>name</h3>
    <input type="text" name="name" v-model="your_name" />
    <button class="btn" @click="send" id="sendname">send!!</button>
    <p v-if="result">{{result}}</p>
    <button v-if="current_page > 0" class="btn" @click="prev" id="pref">prev</button>
    <button v-if="!isLast" class="btn" @click="next" id="next">next</button>
    <table>
      <tr v-for="name in names" :key="name.ID">
        <td>{{name.ID}}</td>
        <td>{{name.Name}}</td>
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
      isLast: false
    };
  },

  methods: {
    async send() {
      const response = await api()
        .post("/hoges", { name: this.your_name })
        .catch(err => err.response || err);
      if (response.status !== 201) {
        this.result = "ERROR";
      } else {
        this.result = "success!!";
      }
      this.reload()
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
        this.names = response.data;
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
        this.names = response.data;
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
        this.names = response.data;
      }
    }
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

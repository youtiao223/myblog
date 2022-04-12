<template>
  <div>
    <v-app-bar app color="light-blue darken-2">
      <v-container class="py-0 fill-height justify-center">
        <v-btn
          text
          color="white"
          class="font-weight-bold"
          @click="$router.push('/')"
          >首页</v-btn
        >
        <v-btn
          v-for="cate in catelist"
          :key="cate.id"
          text
          color="white"
          @click="gotoCate(cate.id)"
        >
          {{ cate.name }}</v-btn
        >
      </v-container>

      <v-spacer></v-spacer>

      <v-responsive max-width="260" color="white">
        <v-text-field
          dense
          flat
          hide-details
          solo-inverted
          rounded
          dark
          placeholder="请输入关键字查找"
          append-icon="mdi-text-search"
          v-on:keyup.enter="searchArt"
          v-model="keywords"
        ></v-text-field>
      </v-responsive>
    </v-app-bar>
  </div>
</template>
<script>
export default {
  data() {
    return {
      catelist: [],
      keywords: "",
    };
  },
  created() {
    this.getCatelist();
  },
  methods: {
    // 获取所有分类
    async getCatelist() {
      const { data: res } = await this.$axios.get("/categories");
      this.catelist = res.data;
    },
    // 转到对应分类
    gotoCate(cid) {
      this.$router.push(`/category/${cid}`).catch((err) => err);
    },
    // 搜索文章
    searchArt() {
      this.$router.push(`/search/${this.keywords}`);
    },
  },
};
</script>
<style>
</style>
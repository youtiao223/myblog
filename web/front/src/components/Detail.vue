<template>
  <v-card min-height="100vh">
    <!-- 文章标题 -->
    <div class="pa-3 text-center text-h5 font-weight-bold blue--text text--lighten-2">
      {{artInfo.title}}
    </div>
    <!-- <v-divider class="ma-3 pa-2"></v-divider> -->
    <div class="ml-3 mr-3 mt-4 pl-2 pr-3">
      <v-icon small>{{'mdi-account'}}</v-icon>
      <span class="text-caption">作者：</span>
      <span class="text-caption">油虾条</span>

      <v-icon small class="ml-5">{{'mdi-calendar-month'}}</v-icon>
      <span class="text-caption">创建于</span>
      <span class="text-caption">{{artInfo.CreatedAt | dateformat('YYYY-MM-DD HH:SS')}}</span>

      <v-icon small class="ml-5">{{'mdi-calendar-month'}}</v-icon>
      <span class="text-caption">更新于</span>
      <span class="text-caption">{{artInfo.UpdatedAt | dateformat('YYYY-MM-DD HH:SS')}}</span>
    </div>
    <div class="d-flex mr-10 justify-center">
      <v-icon class="mr-1" color="green" small>{{ 'mdi-eye' }}</v-icon>
      <span>{{ artInfo.read_count }}</span>
    </div>
    <v-divider class="ma-3 pa-2"></v-divider>

    <!-- 概述 -->
    <v-alert border="left" elevation="2" color="blue" dark outlined class="ma-4 pa-3">
      {{artInfo.desc}}
    </v-alert>
    <!-- 文章内容 -->
    <div class="content ma-4 pa-3 text-justify" v-html="this.content"></div>
  </v-card>

</template>
<script>
import Prism from 'prismjs';
import {marked} from 'marked';


export default {
  props: ['id'],
  data () {
    return {
      artInfo: {},
      content: "",
    }
  },
  created () {
    this.getArtInfo();
  },
  methods: {
    async getArtInfo () {
      const { data: res } = await this.$axios.get(`article/${this.id}`);
      this.artInfo = res.data;
      
      this.content = marked(res.data.content||'',{
          sanitize:true
      })
      window.sessionStorage.setItem('title', this.artInfo.title)
    }
  },
}
</script>
<style scoped>
.content >>> img,
span,
p {
  max-width: 96%;
}
</style>
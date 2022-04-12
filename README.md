在https://gitee.com/itmxs/gin-vue-blog 基础上修改

## 后端技术栈
- 配置文件yaml解析
  - spf13/viper
- 日志工具
  - logrus
  - file-rotatelogs
  - lfshook
- 用户密码存储
  - **md5加盐加密** (目前使用的)
  - bcrypt
  - scrypt（专家方案）
- 用户身份认证
  - jwt-go
- 表单数据验证
  - go-playground/validator (gin的默认验证器)
- 文件上传
  - qiniu 
- 跨域配置
  - gin-contrib/cors

## 配置文件 
application.yaml
``` yaml
# 服务器配置
server:
  mode: debug
  httpPort: 3000
  jwtKey: smo3(*kkc
# 数据库配置
db:
  db: mysql
  host: localhost
  port: 33060
  name: myblog
  user: root
  password:
# 七牛云配置
qiniu:
  accessKey: DJnlV6bI1eK1GASiHKfYzr9FRWlVjAg2smwFt8-d
  secretKey: od9SvbBsOqqRiN85UorLNP91CT5zQVTHo0rjlMYw
  bucket: my-gin-blog
  url: http://r9nc6zhlu.hn-bkt.clouddn.com/

```

## ToDoList （咕咕咕🕊️）

- 💡 博客页面文章排序
- 💡 浏览量记录
- 💡 点赞功能
- 💡 评论功能
- 💡 markdown解析后的美化


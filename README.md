# 个人博客

目录结构
---

```
├── LICENSE // MIT协议
├── README.md
├── api
│   └── v1 // 控制层
│       ├── article.go
│       ├── category.go
│       ├── login.go
│       ├── upload.go
│       └── user.go
├── config
│   └── config.ini // 在这里配置一些网站的参数
├── go.mod
├── go.sum
├── main.go // 入口函数
├── middleware // 中间件
│   ├── cors.go // 跨域
│   ├── jwt.go // jwt认证
│   └── log.go // 日志处理
├── model // 数据库模型
│   ├── article.go
│   ├── category.go
│   ├── db.go
│   └── user.go
├── routers
│   └── router.go // 路由
├── server
│   └── upload.go // 七牛云对象存储
├── static
│   └── admin // 后台管理静态页面
└── utils
    ├── errormsg
    │   └── error_message.go // 错误处理模块
    ├── salt_code.go // 在这里配置加密盐值
    ├── setting.go // 一些默认值
    └── validator
        └── validator.go // 验证器
```

## 使用的技术

Golang, Gin框架, Gorm

## 运行之前要做的

改`salt`值，改`config.ini`，改`setting.go`里面的内容(如果你没有配置`config.ini`的话就会以我的默认值来运行)。

## 运行方式

1. 进入该目录

2. 安装依赖`go mod tidy`

3. 配置`config.ini`

   ```ini
   # 配置博客的一些基本参数
   [server]
   # debug:开发模式 release:生产模式
   AppMode = debug
   HttpPort = :3000
   JwtKey = 自行定义
   
   # 数据库的一些参数
   [database]
   DbHost = 你的地址
   DbPort = 3306
   DbUser = 你的数据库用户名
   DbPassword = 你的数据库密码
   DbName = my-blog
   
   # 七牛云存储
   [qiniu]
   AccessKey = 自
   SecretKey = 己
   Bucket = 填
   QiNiuServer = 写
   ```

4. 你需要配置好你的数据库信息，因为使用了Gorm提供的迁移功能，它会自动生成对应的表。

5. 推荐使用`GoLand`运行。或者你可以使用`go run main.go`来执行

6. 不出意外你可以在`localhost:3000/admin`下看到后台页面，你也可以根据接口写出你自己的页面

## 一些TODO

- [ ] 部署docker
- [ ] 前台展示页面
- [ ] 验证使用中间件来实现
- [ ] 接口文档
- [ ] 评论功能

# 个人博客

目录结构
---

```
├── Dockerfile
├── LICENSE 
├── README.md
├── api // 控制层
│   └── v1
│       ├── article.go
│       ├── category.go
│       ├── login.go
│       ├── profile.go
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
│   ├── profile.go
│   └── user.go
├── routers
│   └── router.go // 路由
├── server
│   └── upload.go // 七牛云对象存储上传配置
├── static
│   ├── admin // 后台管理静态页面
│   └── front // 前台展示静态页面
├── upload
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

**‼️‼️改`salt`值，改`config.ini`，改`setting.go`里面的内容**(如果你没有配置`config.ini`的话就会以我的默认值来运行)。

也许你想访问一下后台管理页面，但是对于管理员没有开放注册，所以你可以先使用接口测试工具，在`user/add`接口下注册一个管理员账户，如下
一定记住`role`为`1`

```json
{
    "username": "admintest",
    "password": "123456",
    "role": 1
}
```

## 运行方式

### 普通方法

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

6. 不出意外你可以在`localhost:3000`下看到前台展示页面，`localhost:3000/admin`下看到后台页面，你也可以根据接口写出你自己的页面

### Docker方式

1. 先讲项目`clone`到本地，并且`cd`到项目目录下(也就是有Dokcerfile的目录)
1. 如`普通方法`第3步，先配置好`config.ini`
2. 执行`docker build -t my-blog .`
3. 第2步完成之后，执行`docker run -d -p 3000:3000 --name my-blog my-blog`

## 一些TODO

- [x] 部署docker
- [x] 前台展示页面
- [ ] 验证使用中间件来实现
- [ ] 接口文档
- [ ] 评论功能
- [ ] 展示页面登录功能

# The Gin Blog
English | [简体中文](./README.md)

Project Structure
---

```
├── Dockerfile
├── LICENSE 
├── README.md
├── README_EN.md
├── api
├── config
│   └── config.ini // blog 
├── go.mod
├── go.sum
├── main.go
├── middlewar
│   ├── cors.go
│   ├── jwt.go 
│   └── log.go 
├── model // database module
├── routers
│   └── router.go
├── server
│   └── upload.go // qiniu Kodo upload config
├── static
│   ├── admin // static admin page
│   └── front // static front page
├── upload
└── utils
    ├── errormsg
    │   └── error_message.go // error message module
    ├── salt_code.go // configure your salt code
    ├── setting.go // some default settings
    └── validator
        └── validator.go
```



## Tech Stack

Golang, Gin Framework, Gorm

## Before running

**‼️‼️Edit the `salt` code, `config.ini` and `setting.go`**(If not, `config.ini` will run with default settings)。

Maybe you want to visit the admin page, however, the Admin can not be registered by the `user/add`. You can use some tools like `Postman` to register an Admin Account.
REMEMBER to set the`role`as`1`

```json
{
    "username": "admintest",
    "password": "123456",
    "role": 1
}
```

## Usage

### Use in Default

1. Enter THE directory

2. Run `go mod tidy`

3. Configure the `config.ini`

   ```ini
   # Configure some basic parameters of the blog
   [server]
   # debug: debug mode, elease: release mode
   AppMode = debug
   HttpPort = :3000
   JwtKey = make your own JwtKey
   
   # Parameters of the Database
   [database]
   DbHost = Your MySQL address
   DbPort = 3306
   DbUser = Your MySQL's username
   DbPassword = MySQL's password
   DbName = my-blog
   
   # Qiniu kodo
   [qiniu]
   AccessKey = 
   SecretKey = 
   Bucket = 
   QiNiuServer = 
   
   [ChatGPT]
   ApiKey = Your ChatGPT ApiKey
   ```

4. You need to configure your database parameters well. Because of the Gorm `Auto Migrations`feature, it will automatically generate the table.

5. `GoLand`is recommended. Or you can just run the `go run main.go` command.

6. Visit`localhost:3000`to see the front page. `localhost:3000/admin`to browse the admin page. You can also make your own!

### Run in Docker

1. `clone`the project, `cd` to the directory which with the `Dockerfile`
1. Do the same thing as `Use in Default`step 3, configure the `config.ini`
2. run the command : `docker build -t my-blog .`
3. And the `docker run -d -p 3000:3000 --name my-blog my-blog`as well.

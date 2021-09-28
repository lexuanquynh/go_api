# golang_api

Create the sample APIs, using clean architecture.

Create by: [codetoanbug.com](https://codetoanbug.com)

## Require:
- install go: https://golang.org/doc/install
- install mysql: https://dev.mysql.com/downloads/mysql/
- install workbench: https://dev.mysql.com/downloads/workbench/
- You need start mysql server in your machine.
## Step to run this project:
- Create user name for mysql: root & create password for root user.
- Create database mysql by command line:

```sql
 export PATH=${PATH}:/usr/local/mysql/bin
 mysql -u root -p
 create database golang_api;
```

- Create .env file with contents below, you need change YOUR_PASSWORD_FOR_ROOT:

```env
DB_USER=root
DB_PASS=YOUR_PASSWORD_FOR_ROOT
DB_HOST=localhost
DB_NAME=golang_api
JWT_SECRET=yourkey_keysecret
```
- Run this project with command line:
```go
go get github.com/mashingan/smapping
go get github.com/golang/crypto
go get github.com/golang-jwt/jwt
go get github.com/joho/godotenv
go get github.com/gin-gonic/gin
go get -u gorm.io/gorm
go get gorm.io/driver/mysql
go run server.go
```
- Try to use postman and use your APIs POST method:
```PHP
http://localhost:8080/api/auth/login
http://localhost:8080/api/auth/register
```

- Header:

```PHP
key:Content-Type
value: application/json
```
- Body:
```PHP
{
    "password":"password",
    "email":"codetoanbug@gmail.com"
}
```

- User APIs GET method:
```PHP
http://localhost:8080/api/user/profile
```
- Header:

```PHP
key:Content-Type
value: application/json
key: Authorization
value: <token from login api>
```

- User APIs PUT method:
```PHP
http://localhost:8080/api/user/profile
```
- Header:

```json
key:Content-Type
value: application/json
key: Authorization
value: <token from login api>
```
- Body:
```json
{
    "name":"Code toan Bug",
    "email":"codetoanbug@gmail.com"
}
```
- If you get error:
```json
Error: connect ECONNREFUSED 127.0.0.1:8080
```
Please change DB_PASS=YOUR_PASSWORD in .env file! 

- If you want to deploy to vps server, you need change code in database-config.go, sample:
```go
     errEnv := godotenv.Load("/usr/local/src/golang_api/.env")
```
## APIs for books:

- Get books:
```go
[GET] http://localhost:8080/api/books
```
- Insert book:
```go
[POST] http://localhost:8080/api/books
```
```json
{
    "title":"sach hay",
    "description":"sach hay ve van hoc"
}
```

- Update book:
```
[PUT] http://localhost:8080/api/books/1
```
```json
{
    "id":1,
    "title":"sach hay",
    "description":"sach hay ve van hoc nha"
}
```

- Delete book:
```go
[DELETE] http://localhost:8080/api/books/2
```

## Third party libraries:

- https://github.com/gin-gonic/gin
- https://github.com/joho/godotenv
- https://github.com/golang-jwt/jwt
- https://github.com/golang/crypto
- https://github.com/mashingan/smapping

## License 
This repository is released under an MIT license.  See [License.md](https://github.com/lexuanquynh/go_api/blob/main/LICENSE) for more information.

# golang_api

Create the sample APIs, using clean architecture.

Create by: codetoanbug.com

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
- Try to use postman and use your APIs:
```PHP
http://localhost:8080/api/auth/login
http://localhost:8080/api/auth/register
```
- If you get error:
```
Error: connect ECONNREFUSED 127.0.0.1:8080
```
Please change DB_PASS=YOUR_PASSWORD in .env file! 
## Third party libraries:

- https://github.com/gin-gonic/gin
- https://github.com/joho/godotenv
- https://github.com/golang-jwt/jwt
- https://github.com/golang/crypto
- https://github.com/mashingan/smapping

## License 
This repository is released under an MIT license.  See [License.md](https://github.com/lexuanquynh/go_api/blob/main/LICENSE) for more information.

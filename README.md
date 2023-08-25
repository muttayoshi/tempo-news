# Tempo Media Web App - Golang 

tempo-news adalah aplikasi web sederhana yang menggunakan bahasa [Go](https://go.dev/),  dibuat dengan web framework [Gin](https://github.com/gin-gonic/gin).

**tempo-news menggunakan:**

- [Gin Web Framework](https://gin-gonic.com/).
- [GORM](https://gorm.io/)
- [Load .env](https://github.com/joho/godotenv)
- [Go Cryptography](https://pkg.go.dev/golang.org/x/crypto#section-readme)
- [JWT](https://github.com/golang-jwt/jwt)
- [Postgres Database](https://www.postgresql.org/)
- [Validator](https://pkg.go.dev/github.com/go-playground/validator/v10Z)

## Getting started

### Running Postgres

jalankan postgres dengan docker compose

```
docker-compose up
```

### Running tempo-news

jalankan aplikasi web server

```
go run main.go
```

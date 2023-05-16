# Booking API

Originla repo: https://github.com/gandra/bookstore_users-api switched to GORM

## Stack used: 
- https://github.com/gin-gonic/gin
- https://gorm.io/
- https://github.com/uber-go/zap
- Postgresql


## Run Mysql in docker

Mariadb:
```
docker run --name bookinggormdb -e POSTGRES_PASSWORD=mysecret -e POSTGRES_DB=bookings_users -p 5432:5432 -d postgres
docker container start bookinggormdb
```

```
docker rm bookinggormdb
```

https://www.udemy.com/course/golang-how-to-design-and-build-rest-microservices-in-go/learn/lecture/16459582#overview


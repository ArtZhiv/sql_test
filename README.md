# sql_test

```
```
go run cmd\main.go

iwr -useb get.scoop.sh | iex
scoop install migrate

migrate -path ./schema -database 'mysql://login:password@tcp(host:port)/database' down
migrate -path ./schema -database 'mysql://login:password@tcp(host:port)/database' up
```
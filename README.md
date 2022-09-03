migration
```
migrate.exe create -ext sql -dir db\migration -seq init_schema

migrate -path db/migration -database "postgresql://root:admin@127.0.0.1:5432/simple_bank?sslmode=disable" -verbose up

```

sqlc
```
sqlc init
docker run --rm -v D:\git\bank:/src -w /src kjconroy/sqlc generate
```

mock
```
mockgen -destination db/mock/store.go github.com/LiShangAn/bank/db/sqlc IStore
```
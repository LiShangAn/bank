migration
```
migrate.exe create -ext sql -dir db\migration -seq init_schema

migrate -path db/migration -database "postgresql://postgres:admin@127.0.0.1:5432/postgres?sslmode=disable" -verbose u p
```

sqlc
```
sqlc init
```
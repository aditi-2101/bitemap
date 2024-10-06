migrateup: 
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/bitemap?sslmode=disable" -verbose up

migrateup1: 
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/bitemap?sslmode=disable" -verbose up 1

migratedown: 
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/bitemap?sslmode=disable" -verbose down

migratedown1: 
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/bitemap?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -race -cover ./... 

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go bitemap/db/sqlc Store

.PHONY:	postgres createdb dropdb migrateup migratedown sqlc server mock migrateup1 migratedown1
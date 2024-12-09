postgres:
	 docker run --name postgres_learning -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=professor -d postgres
createdb:
	docker exec -it postgres_learning createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres_learning dropdb simple_bank

migrateup:
	 migrate -path db/migration -database "postgresql://root:professor@localhost:5432/simple_bank?sslmode=disable" -verbose up
	 
migratedown:
	 migrate -path db/migration -database "postgresql://root:professor@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc:
	 sqlc generate
	 
test:
	go test -v -cover ./...

server:
	go run main.go
mock:
	mockgen --package mockDB --destination db/mock/store.go github.com/umeh-promise/simple_bank/db/sqlc Store

.PHONY: createdb dropdb migrateup migratedown sqlc test server
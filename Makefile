postgres:
	docker run --name postgres12  -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14-alpine

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres12 dropdb simple_bank

migrateup:
	migrate -path db/migrations -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up  

migrateup1:
	migrate -path db/migrations -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up 1


migratedown:
	echo "y"|migrate -path db/migrations -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" down -all

migratedown1:
	migrate -path db/migrations -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mockstore:
	 mockgen -build_flags=--mod=mod -package mockdb -destination db/mock/store.go  github.com/marioTiara/api.simplebank.go/db/sqlc Store 


.PHONY: postgres createdb dropdb migrateup migratedown sqlc server mockstore migratedown1 migrateup1

postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine
createdb:
	docker exec -it postgres12 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres12 dropdb  simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose  up
migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose  down

sqlc:
	docker run --rm -v C:\Users\amitw\Documents\simplebank:/src -w /src kjconroy/sqlc generate
test:
	go test -v -cover ./...


.server:
	go run main.go

.PHONY:postgres createdb dropdb migrateup migratedown sqlc test server


# $ curl -L https://github.com/golang-migrate/migrate/releases/download/v4.16.2/migrate.linux-amd64.tar.gz | tar xvz
# migrate.linux-amd64.tar.gz 	

#sudo docker exec -it postgres12 psql -U root simple_bank


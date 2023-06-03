postgres:
	docker run --name postgres15 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:15-alpine

createdb:
	docker exec -it postgres15 createdb --username=root --owner=root bankita

dropdb:
	docker exec -it postgres15 dropdb bankita

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/bankita?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/bankita?sslmode=disable" -verbose down

.PHONY: postgres createdb dropdb migrateup migratedown
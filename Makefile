DB_URL=postgresql://postgres:admin123@localhost:2022/todo?sslmode=disable
download:
	go mod download

dep:
	go mod tidy

run:
	go run main.go

testing:
	go test test/tasks_test.go -v

build:
	go build -o bin/moonlay ./main.go

docker-image:
	docker build -t todo:latest .

docker-run:
	docker run --name todo -p 8084:8080 todo:latest

migrate:
	migrate create -ext sql -dir db/migrations create_table_user

postgres:
	docker run --detach --name postgres -d -p 2022:5432 -e POSTGRES_PASSWORD=admin123  postgres

createdb:
	docker exec -it postgres createdb --username=postgres  todo

dropdb:
	docker exec -it postgres dropdb new-app

migrate_up:
	migrate -path db/migrations -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down



.PHONY: createdb dropdb
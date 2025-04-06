include Makefile.local

dev:
	go run main.go

build:
	go build

db-migrate:
	goose up

db-delete:
	rm -f database/main.sqlite

db-seed:
	./scripts/seed.sh

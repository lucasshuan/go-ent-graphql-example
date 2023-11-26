server:
	go run cmd/main.go

docker-up:
	docker-compose up -d

docker-down:
	docker-compose down -v

.PHONY: entity

entity:
	@read -p "Enter the name of the new entity: " name; \
	go run -mod=mod entgo.io/ent/cmd/ent new $$name

deps:
	go mod tidy
export

run:
	go run ./cmd/main.go
.PHONY: run
compose-up:
	docker-compose up --build -d && docker-compose logs -f test-app
.PHONY: compose-up
compose-up-all-logs:
	docker-compose up --build -d && docker-compose logs -f
.PHONY: compose-up-all-logs
compose-down:
	docker-compose down --remove-orphans
.PHONY: compose-down
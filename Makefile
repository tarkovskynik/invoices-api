build:
	docker-compose build invoices-api

run:
	docker-compose up invoices-api

migrate:
	migrate -path ./schema -database 'postgres://postgres:qwerty@0.0.0.0:5432/postgres?sslmode=disable' up
up:
	docker-compose up -d
	docker exec -it postgres createdb --username=postgres --owner=postgres url-services
	migrate -path ./pkg/postgres/migrations/schema -database "postgresql://postgres:qwer1234@localhost:5430/url-services?sslmode=disable" -verbose down
	migrate -path ./pkg/postgres/migrations/schema -database "postgresql://postgres:qwer1234@localhost:5430/url-services?sslmode=disable" -verbose up
start:
	docker-compose up -d
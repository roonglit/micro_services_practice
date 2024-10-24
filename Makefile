
migrate :
	migrate create -ext sql -dir db/migrations -seq create_users_table

migrateup:
	migrate -path db/migrations -database "postgres://root:root@localhost:5432/auth_development?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migrations -database "postgres://root:root@localhost:5432/auth_development?sslmode=disable" -verbose down 1

migrategooseup:
	goose postgres "user=root password=root dbname=auth_development sslmode=disable host=localhost" -dir "./db/migrations/goose" up

migrategoosedown:
	goose postgres "user=root password=root dbname=auth_development sslmode=disable host=localhost" -dir "./db/migrations/goose" down

dropdb: 
	docker exec -it micro_services_practice-postgres-1 dropdb auth_development

createdb: 
	docker exec -it micro_services_practice-postgres-1 createdb --username=root --owner=root auth_development

sqlc: 
	sqlc generate

.PHONY: migrateup migratedown dropdb createdb sqlc migrategooseup migrategoosedown

# Dirty database version 1. Fix and force version:
# 	migrate -database "postgres://root:root@localhost:5432/auth_development?sslmode=disable" -path db/migrations force 15  
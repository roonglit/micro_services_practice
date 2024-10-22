
migrate :
	migrate create -ext sql -dir db/migrations -seq create_users_table

migrateup:
	migrate -path db/migrations -database "postgres://root:root@localhost:5432/auth_development?sslmode=disable" up 1

migratedown:
	migrate -path db/migrations -database "postgres://root:root@localhost:5432/auth_development?sslmode=disable" down 1

dropdb: 
	docker exec -it micro_services_practice-postgres-1 dropdb auth_development

createdb: 
	docker exec -it micro_services_practice-postgres-1 createdb --username=root --owner=root auth_development

sqlc: 
	sqlc generate

.PHONY: migrateup migratedown dropdb createdb sqlc

# Dirty database version 1. Fix and force version:
# 	migrate -database "postgres://root:root@localhost:5432/auth_development?sslmode=disable" -path db/migrations force 15  
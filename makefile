
migration-up:
	migrate -path ./migrations/postgres -database 'postgres://admin:admin@localhost:5432/postgres?sslmode=disable' up
	
migration-down:
	migrate -path ./migrations/postgres -database 'postgres://admin:admin@localhost:5432/postgres?sslmode=disable' down
	
migration-force-1v:
	migrate -path ./migrations/postgres -database 'postgres://admin:admin@localhost:5432/postgres?sslmode=disable' force 1
	
	
migrate -path ./migrations/postgres -database 'postgres://admin:admin@localhost:5432/postgres?sslmode=disable' up 1 01_add_column.up
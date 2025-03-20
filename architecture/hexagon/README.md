To be continued

### Tech Summary
- Web Framework: [Gin](https://github.com/gin-gonic/gin)
- Database: Postgres
    - Generate Go code from SQL: [sqlc](https://github.com/sqlc-dev/sqlc)

### How to run

Run the following command to start the server for development:
```bash
./scripts/run-dev.sh
````

Migrate new changes to the database:
```bash
./scripts/run-migrate.sh
```

Generate Go code from SQL Queries:
```bash
./scripts/generate-query.sh
```

Notes: Folder sql in `internal/adapters/framework/right/database/sql` contains schema migration and sql queries to generate

package main

import (
	"coffee/internal/adapters/framework/left/http"
	"coffee/internal/adapters/framework/right/database"
	rightPort "coffee/internal/ports/right/database"
	"database/sql"
	"log"
)

type DB struct {
	core   rightPort.CoreDbPort
	engine *sql.DB
}

func main() {
	env := setupEnvironment()

	db := setupDB(env)

	defer db.core.CloseConnection()

	di := NewDI(db.engine, env)

	http.NewAdapter(di.BuildApplicationPorts()).Run(env.PORT)
}

func setupDB(env ENV) DB {
	dbAdapter, err := database.NewAdapter(env.DB_HOST, env.DB_PORT, env.DB_USER, env.DB_PASS, env.DB_DATABASE)

	if err != nil {
		log.Fatalf("failed to initiate dbase connection: %v", err)
	}

	return DB{
		core:   dbAdapter,
		engine: dbAdapter.Db,
	}
}

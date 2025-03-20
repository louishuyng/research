package database

import (
	"database/sql"
	"embed"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

//go:embed sql/schema/*.sql
var embedMigrations embed.FS

type Adapter struct {
	Db *sql.DB
}

func NewAdapter(host, port, user, password, dbname string) (*Adapter, error) {
	databaseURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port, dbname)

	db, err := sql.Open("postgres", databaseURL)

	if err != nil {
		log.Fatalf("db connection failure: %v", err)
	}

	// test db connection
	err = db.Ping()
	if err != nil {
		log.Fatalf("db ping failure: %v", err)
	}

	log.Println("db connection successful")

	goose.SetBaseFS(embedMigrations)
	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatalf("dialect failure: %v", err)
	}

	if err := goose.Up(db, "sql/schema"); err != nil {
		log.Fatalf("migrations failure: %v", err)
	}

	return &Adapter{
		Db: db,
	}, nil
}

func (a *Adapter) CloseConnection() {
	err := a.Db.Close()

	if err != nil {
		log.Fatalf("db close failure: %v", err)
	}
}

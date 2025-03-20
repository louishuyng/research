package main

import (
	"coffee/internal/adapters/framework/left/http"
	"coffee/internal/adapters/framework/right/database"
	"coffee/internal/adapters/framework/right/database/repository"
	"coffee/internal/application/api/auth"
	"coffee/internal/application/core/user"
	"database/sql"
)

type DI struct {
	engine *sql.DB
	env    ENV
}

func NewDI(engnine *sql.DB, env ENV) *DI {
	return &DI{
		engine: engnine,
		env:    env,
	}
}

func (di *DI) BuildApplicationPorts() http.ApiPorts {

	return http.ApiPorts{
		AuthPort: di.setupAuthDI(),
	}
}

func (di *DI) setupAuthDI() *auth.AuthApplication {
	q := repository.New(di.engine)

	userDb := database.NewUserDbAdapter(q)
	userCore := user.New()

	return auth.NewAuthApplication(userDb, userCore, auth.AuthENV{
		JWT_SECRET: di.env.JWT_SECRET,
	})
}

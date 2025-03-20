package main

import (
	"os"

	"github.com/joho/godotenv"
)

const DEFAULT_PORT = "9000"
const DEFAULT_JWT_SECRET = "secret"

type ENV struct {
	DB_HOST     string
	DB_PORT     string
	DB_USER     string
	DB_PASS     string
	DB_DATABASE string
	PORT        string
	JWT_SECRET  string
}

func setupEnvironment() ENV {
	godotenv.Load()

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbDatabase := os.Getenv("DB_DATABASE")

	jwtSecret := os.Getenv("JWT_SECRET")
	port := os.Getenv("PORT")

	if port == "" {
		port = DEFAULT_PORT
	}

	if jwtSecret == "" {
		jwtSecret = DEFAULT_JWT_SECRET
	}

	return ENV{
		DB_HOST:     dbHost,
		DB_PORT:     dbPort,
		DB_USER:     dbUser,
		DB_PASS:     dbPass,
		DB_DATABASE: dbDatabase,
		PORT:        port,
		JWT_SECRET:  jwtSecret,
	}
}

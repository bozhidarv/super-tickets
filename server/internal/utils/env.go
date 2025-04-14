package utils

import (
	"os"
	"supertickets/internal/models"
)

var EnvVars *models.ReadOnlyEnvVars

func LoadEnvVars() {
	envVars := &models.EnvVars{}

	dbUrl := os.Getenv("DATABASE_URL")
	if dbUrl == "" {
		panic("DATABASE_URL environment variable not found")
	}
	envVars.SetDbUrl(dbUrl)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	envVars.SetPort(port)

	jwtKey := os.Getenv("JWT_KEY")
	if jwtKey == "" {
		panic("JWT_KEY environment variable not found")
	}
	envVars.SetJwtKey(jwtKey)

	EnvVars = &models.ReadOnlyEnvVars{EnvVars: *envVars}
}

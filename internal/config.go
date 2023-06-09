package internal

import (
	"fmt"
	"github.com/joho/godotenv"
	"time"

	"github.com/caarlos0/env/v6"
	"github.com/sirupsen/logrus"
)

var Config config

func init() {
	if err := godotenv.Load(); err != nil {
		logrus.Warnf("Fail loading .env: %v", err)
	}

	if err := env.Parse(&Config); err != nil {
		logrus.Fatalf("Error initializing: %s", err.Error())
	}
}

type config struct {
	Environment        string `env:"APP_ENV"`
	JWTSecret          string `env:"APP_JWT_SECRET"`
	RandomStringAPIKey string `env:"APP_RANDOM_STRING_API_KEY"`
	DatabaseConnection
}

type DatabaseConnection struct {
	DbUser     string        `env:"DB_USER" envDefault:""`
	DbPassword string        `env:"DB_PASSWORD" envDefault:""`
	DbHost     string        `env:"DB_HOST" envDefault:""`
	DbName     string        `env:"DB_NAME" envDefault:""`
	Timeout    time.Duration `env:"DB_TIMEOUT" envDefault:"1s"`
}

func (b *DatabaseConnection) DBConnectionString() string {
	return fmt.Sprintf(
		"mongodb://%s:%s@%s:27017",
		b.DbUser,
		b.DbPassword,
		b.DbHost,
	)
}

func (b *DatabaseConnection) DBName() string {
	return b.DbName
}

func (b *DatabaseConnection) DBTimeout() time.Duration {
	return b.Timeout
}

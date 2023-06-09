package main

import (
	"github.com/jesusEstaba/calculator/internal"
	"github.com/jesusEstaba/calculator/internal/database"
	"github.com/jesusEstaba/calculator/internal/database/seeds"
	"github.com/sirupsen/logrus"
	"log"
)

func main() {
	db, err := database.NewDatabase(&internal.Config)
	if err != nil {
		logrus.Fatal(err)
	}

	err = seeds.SeedOperations(db)
	if err != nil {
		log.Fatal(err)
	}
}

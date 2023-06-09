package seeds

import (
	"github.com/jesusEstaba/calculator/internal/database"
	"github.com/jesusEstaba/calculator/pkg/domain"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

func SeedOperations(db *mongo.Database) error {
	repository, err := database.NewRepository[domain.Operation](db, "operations")
	if err != nil {
		logrus.Fatal(err)
	}

	operations := []domain.Operation{
		{Type: "addition", Cost: 800},
		{Type: "subtraction", Cost: 800},
		{Type: "multiplication", Cost: 1100},
		{Type: "division", Cost: 1100},
		{Type: "square_root", Cost: 1500},
		{Type: "random_string", Cost: 2000},
	}

	ops := make([]*domain.Operation, len(operations))
	for i := range operations {
		ops[i] = &operations[i]
	}
	_, err = repository.InsertMany(ops)
	if err != nil {
		return err
	}

	logrus.Infof("Operations seeded successfully")

	return nil
}

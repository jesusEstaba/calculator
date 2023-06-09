package persistence

import (
	"errors"
	"github.com/jesusEstaba/calculator/internal/database"
	"github.com/jesusEstaba/calculator/pkg/domain"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strings"
)

type OperationRepositoryImplementation struct {
	recordRepository    database.CRUDRepository[domain.Record, *domain.Record]
	operationRepository database.CRUDRepository[domain.Operation, *domain.Operation]
}

func NewOperationRepositoryImplementation(db *mongo.Database) domain.OperationRepository {
	recordRepository, err := database.NewRepository[domain.Record](db, "records")
	if err != nil {
		logrus.Fatal(err)
	}

	operationRepository, err := database.NewRepository[domain.Operation](db, "operations")
	if err != nil {
		logrus.Fatal(err)
	}

	return &OperationRepositoryImplementation{
		recordRepository,
		operationRepository,
	}
}

func (r *OperationRepositoryImplementation) GetOperation(operation string) (*domain.Operation, error) {
	ope, err := r.operationRepository.FindByKey("type", operation)
	if err != nil && strings.Contains(err.Error(), "no documents in result") {
		return nil, errors.New("operation not found")
	}

	return ope, nil
}

func (r *OperationRepositoryImplementation) RecordOperation(record domain.Record) error {
	return r.recordRepository.Insert(&record)
}

func (r *OperationRepositoryImplementation) GetRecordsByUserAndSearchTermPaginated(search domain.RecordSearch) ([]*domain.Record, error) {
	opts := &options.FindOptions{}

	if search.Page > 0 {
		skip := (search.Page - 1) * search.PerPage
		opts.SetSkip(int64(skip))
	}

	opts.SetLimit(int64(10))
	if search.PerPage > 0 {
		opts.SetLimit(int64(search.PerPage))
	}

	if search.Sort == "desc" {
		opts.SetSort(bson.D{{Key: "_id", Value: -1}})
	}

	filter := bson.M{
		"user_id": search.UserID,
	}

	if search.SearchTerm != "" {
		filter["$or"] = []bson.M{
			{"operation_id": bson.M{"$regex": search.SearchTerm, "$options": "i"}},
			{"amount": bson.M{"$regex": search.SearchTerm, "$options": "i"}},
		}
	}

	return r.recordRepository.GetByFilter(
		filter,
		opts,
	)
}

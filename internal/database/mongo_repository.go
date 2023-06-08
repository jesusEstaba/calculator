package database

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"reflect"
	"regexp"
	"strings"
)

type Model[T any] interface {
	GetID() *primitive.ObjectID
	SetID(*primitive.ObjectID)
	UpsertKey() *bson.M
	*T
}

type CRUDRepository[V any, T Model[V]] interface {
	Insert(T) error
	FindByID(id string) (T, error)
	FindByKey(key string, value any) (T, error)
	GetByFilter(filter interface{}, opts ...*options.FindOptions) ([]T, error)
}

type mongoErr struct {
	Message string `json:"message"`
	Type    string `json:"type"`
}

func (e *mongoErr) Error() string {
	return fmt.Sprintf("[%s]: %s", e.Type, e.Message)
}

func mongoError(err error) error {
	if err == nil {
		return err
	}

	return &mongoErr{
		Message: err.Error(),
		Type:    reflect.TypeOf(err).Name(),
	}
}

func IDValidation(matchID string) error {
	r, err := regexp.Compile(`^[a-f\d]{24}$`)
	if err != nil {
		return err
	}
	if !r.Match([]byte(matchID)) {
		return errors.New("invalid primitive ID")
	}
	return nil
}

func NewRepository[V any, T Model[V]](db *mongo.Database, coll string) (CRUDRepository[V, T], error) {
	ctx, cancel := nctx()
	defer cancel()

	if db == nil {
		return nil, errors.New("nil database reference")
	}

	if err := db.CreateCollection(ctx, coll); err != nil {
		if !strings.Contains(err.Error(), "already exists") {
			return nil, err
		}
	}

	collection := db.Collection(coll)

	var subj V
	index, isIndexable := any(subj).(indexableModel)
	if isIndexable {
		indices := index.GetIndexes()
		if len(indices) > 0 {
			_, err := collection.Indexes().CreateMany(ctx, indices)
			if err != nil {
				return nil, err
			}
		}
	}

	return &crudRepository[V, T]{collection: collection}, nil
}

type indexableModel interface {
	GetIndexes() []mongo.IndexModel
}

type crudRepository[V any, T Model[V]] struct {
	collection *mongo.Collection
}

func (r *crudRepository[V, T]) Insert(v T) error {
	var (
		id           primitive.ObjectID
		err          error
		upsertFilter = v.UpsertKey()
	)
	ctx, cancel := nctx()
	defer cancel()

	logrus.Infof("Inserting: %v", v)

	if upsertFilter != nil {
		var res *mongo.UpdateResult
		res, err = r.collection.ReplaceOne(ctx, upsertFilter, v, options.Replace().SetUpsert(true))
		if err != nil {
			logrus.Infof("Failed Inserting: %v", v)
			return err
		}
		logrus.Info(res.MatchedCount, res.ModifiedCount, res.UpsertedCount, res.UpsertedID)
		if res.UpsertedCount > 0 {
			id = res.UpsertedID.(primitive.ObjectID)
		}
	} else {
		var res *mongo.InsertOneResult
		res, err = r.collection.InsertOne(ctx, v)
		if err != nil {
			logrus.Infof("Failed Inserting: %v", v)
			return err
		}
		id = res.InsertedID.(primitive.ObjectID)
	}

	v.SetID(&id)

	logrus.Infof("Inserted: %v", v)

	return nil
}

func (r *crudRepository[V, T]) FindByID(id string) (T, error) {
	err := IDValidation(id)
	if err != nil {
		return nil, err
	}
	pid, errConvert := primitive.ObjectIDFromHex(id)
	if errConvert != nil {
		return nil, errConvert
	}

	return r.FindByKey("_id", pid)
}

func (r *crudRepository[V, T]) FindByKey(key string, value any) (T, error) {
	var result T

	ctx, cancel := nctx()
	defer cancel()

	res := r.collection.FindOne(
		ctx,
		bson.D{{Key: key, Value: value}})

	if err := res.Decode(&result); err != nil {
		return nil, mongoError(err)
	}

	return result, nil
}

func (r *crudRepository[V, T]) GetByFilter(filter interface{}, opts ...*options.FindOptions) ([]T, error) {
	result := []T{}
	ctx, cancel := nctx()
	defer cancel()

	res, err := r.collection.Find(ctx, filter, opts...)
	if err != nil {
		return nil, mongoError(err)
	}

	defer res.Close(ctx)

	for res.Next(ctx) {
		var u T

		if err := res.Decode(&u); err != nil {
			return nil, mongoError(err)
		}

		result = append(result, u)
	}

	return result, mongoError(res.Err())
}

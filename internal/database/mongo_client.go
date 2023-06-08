package database

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var mongoTimeout = 60 * time.Second

type MongoConfig interface {
	DBConnectionString() string
	DBName() string
	DBTimeout() time.Duration
}

func nctx() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), mongoTimeout)
}

func NewDatabase(c MongoConfig) (*mongo.Database, error) {
	mongoTimeout = c.DBTimeout()
	conn, err := NewClient(c.DBConnectionString())
	if err != nil {
		return nil, err
	}

	db := conn.Database(c.DBName())

	return db, nil
}

func NewClient(uri string) (*mongo.Client, error) {
	logrus.Infof("db: Connecting to %s", uri)
	conctx, cancel := nctx()
	defer cancel()

	client, err := mongo.Connect(conctx, options.Client().ApplyURI(uri).SetRetryWrites(false))
	if err != nil {
		return nil, err
	}

	pingctx, cancelping := nctx()
	defer cancelping()
	if err := client.Ping(pingctx, readpref.Primary()); err != nil {
		return nil, err
	}

	logrus.Debug("db: Successfully connected and pinged.")
	return client, nil
}

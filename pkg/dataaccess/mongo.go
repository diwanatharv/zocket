package dataaccess

import (
	"awesomeProject6/pkg/config/dbconfig"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongodbs struct {
	db *mongo.Collection
}

func MongoManager(collectioname string) *mongodbs {
	return &mongodbs{
		db: dbconfig.Makemongoserver().Collection(collectioname),
	}
}

type mongomethodsinterface interface {
	Insert(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error)
	Totalcount(ctx context.Context, opts ...*options.EstimatedDocumentCountOptions) (int64, error)
	Findone(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) *mongo.SingleResult
	Updateone(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error)
}

func (m *mongodbs) Insert(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	return m.db.InsertOne(ctx, document, opts...)
}

func (m *mongodbs) Totalcount(ctx context.Context, opts ...*options.EstimatedDocumentCountOptions) (int64, error) {
	return m.db.EstimatedDocumentCount(ctx, opts...)
}
func (m *mongodbs) Findone(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) *mongo.SingleResult {
	return m.db.FindOne(ctx, filter, opts...)
}
func (m *mongodbs) Updateone(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	return m.db.UpdateOne(ctx, filter, update, opts...)
}

package demo

import (
	"context"
	"runtime/debug"
	"sync"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type BaseEntity interface {
	GetId() string
	SetId(id string)
}

type PageFilter struct {
	SortBy     string
	SortMode   int8
	Limit      *int64
	Skip       *int64
	Filter     map[string]interface{}
	RegexFiler map[string]string
}

type MongoClient struct {
	Client *mongo.Client
	Ctx    context.Context
}

var Mongo *MongoClient

var Conf struct {
	MongoUri string
	MongoDB  string
}

func init() {

	var once sync.Once
	once.Do(func() {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel() // bug may happen
		if client, err := mongo.Connect(ctx, options.Client().ApplyURI(Conf.MongoUri)); err == nil {
			Mongo = &MongoClient{}
			Mongo.Ctx = ctx
			Mongo.Client = client
		}
	})
}

func (m *MongoClient) Create(collection string, e BaseEntity) (error, string) {
	var err error
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				debug.PrintStack()
			}
		}
	}()
	collections := m.Client.Database(Conf.MongoDB).Collection(collection)
	// e.SetId(UUID())
	if cid, err := collections.InsertOne(m.Ctx, e); err == nil {
		return nil, cid.InsertedID.(primitive.ObjectID).Hex()
	}
	return err, ""
}

func (m *MongoClient) Get(collection, id string) (err error, e BaseEntity) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				debug.PrintStack()
			}
		}
	}()
	collections := m.Client.Database(Conf.MongoDB).Collection(collection)
	objID, _ := primitive.ObjectIDFromHex(id)
	result := collections.FindOne(m.Ctx, bson.M{"_id": objID})
	result.Decode(&e)
	return
}

func (m *MongoClient) GetOne(collection, id string) (err error, e interface{}) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				debug.PrintStack()
			}
		}
	}()
	collections := m.Client.Database(Conf.MongoDB).Collection(collection)
	result := collections.FindOne(m.Ctx, bson.M{"Id": id})
	result.Decode(&e)
	return
}

func (m *MongoClient) Count(collection string, filter PageFilter) (err error, c int64) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				debug.PrintStack()
			}
		}
	}()
	if filter.RegexFiler != nil {
		for k, v := range filter.RegexFiler {
			filter.Filter[k] = primitive.Regex{Pattern: v, Options: ""}
		}
	}
	collections := m.Client.Database(Conf.MongoDB).Collection(collection)
	collections.CountDocuments(m.Ctx, filter.Filter, &options.CountOptions{Skip: filter.Skip, Limit: filter.Limit})
	return
}

func (m *MongoClient) List(collection string, filter PageFilter) (err error, e []interface{}) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				debug.PrintStack()
			}
		}
	}()
	if filter.RegexFiler != nil {
		for k, v := range filter.RegexFiler {
			filter.Filter[k] = primitive.Regex{Pattern: v, Options: ""}
		}
	}
	collections := m.Client.Database(Conf.MongoDB).Collection(collection)
	cur, err := collections.Find(m.Ctx, filter, &options.FindOptions{Limit: filter.Limit, Skip: filter.Skip, Sort: bson.M{filter.SortBy: filter.SortMode}})
	defer cur.Close(m.Ctx)
	if err == nil {
		for cur.Next(m.Ctx) {
			var e interface{}
			cur.Decode(&e)
		}
	}
	return
}

func (m *MongoClient) Delete(collection, id string) (error, bool) {
	var err error
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				debug.PrintStack()
			}
		}
	}()

	collections := m.Client.Database(Conf.MongoDB).Collection(collection)
	objID, _ := primitive.ObjectIDFromHex(id)
	result, err := collections.DeleteOne(m.Ctx, bson.M{"_id": objID})
	return err, result.DeletedCount == 1
	// result, err := collections.DeleteMany(ctx, bson.M{"phone": primitive.Regex{Pattern: "456", Options: ""}})
}

func (m *MongoClient) Modify(collection string, e BaseEntity) (error, bool) {
	var err error
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				debug.PrintStack()
			}
		}
	}()
	collections := m.Client.Database(Conf.MongoDB).Collection(collection)
	// collections.UpdateOne
	// collections.UpdateMany
	objID, _ := primitive.ObjectIDFromHex(e.GetId())
	result, err := collections.ReplaceOne(m.Ctx, bson.M{"_id": objID}, e)
	return err, result.ModifiedCount == 1
}

type Test struct {
	Id       string
	Name     string
	Creator  string
	CreateAt int64
}

func (e Test) GetId() string {
	return e.Id
}

func (e *Test) SetId(id string) {
	e.Id = id
}

func TestMongo(ts *testing.T) {
	var t Test
	Mongo.Create("LicenseReview", &t)
}

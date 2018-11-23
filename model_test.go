package model

import (
	"fmt"
	"testing"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"

	"github.com/golib/assert"
	"log"
)

var (
	newMockConfig = func() *Config {
		cfgStr := (`{
      "host": "localhost:27017",
      "user": "root",
      "password": "",
      "database": "testing_model",
      "mode": "Strong",
      "pool": 5,
      "timeout": 5
}`)

		modelConfig, _ := NewConfig([]byte(cfgStr))

		return modelConfig
	}

	defaultLogger = &mockLogger{}

	mockModel = func() *Model {
		return NewModel(newMockConfig(), defaultLogger)
	}()

	mockTestModelIndexes = []mgo.Index{
		{
			Key:    []string{"name"},
			Unique: false,
		},
	}
)

type (
	mockTestModel struct {
		Id   bson.ObjectId `bson:"_id" json:"id"`
		Name string        `bson:"name" json:"name"`
	}

	mockLogger struct{}
)

func (_ *mockLogger) Errorf(format string, v ...interface{}) {
	log.Printf(format, v...)
}

func (_ *mockLogger) Panic(v ...interface{}) {
	log.Panic(v...)
}

func Test_NewModel(t *testing.T) {
	assertion := assert.New(t)

	model := NewModel(newMockConfig(), defaultLogger)
	assertion.NotNil(model)
	assertion.NotNil(model.session)
	assertion.Nil(model.collection)
	assertion.NotNil(model.config)
	assertion.NotNil(model.logger)
	assertion.Condition(func() bool {
		return fmt.Sprintf("%p", defaultLogger) == fmt.Sprintf("%p", model.logger)
	})
	assertion.Empty(model.indexes)
}

func Test_ModelUse(t *testing.T) {
	assertion := assert.New(t)

	model := NewModel(newMockConfig(), defaultLogger)
	assertion.Equal("testing_model", model.Database())

	model.Use("testing_database")
	assertion.Equal("testing_database", model.Database())
}

func Test_ModelCopy(t *testing.T) {
	assertion := assert.New(t)
	model := NewModel(newMockConfig(), defaultLogger)

	copiedModel := model.Copy()
	assertion.Condition(func() bool {
		return fmt.Sprintf("%p", model) != fmt.Sprintf("%p", copiedModel)
	})
	assertion.Condition(func() bool {
		return fmt.Sprintf("%p", model.session) != fmt.Sprintf("%p", copiedModel.session)
	})
	assertion.Condition(func() bool {
		return fmt.Sprintf("%p", model.config) != fmt.Sprintf("%p", copiedModel.config)
	})
	assertion.Condition(func() bool {
		return fmt.Sprintf("%p", model.logger) == fmt.Sprintf("%p", copiedModel.logger)
	})

	copiedModel.Use("testing_database")
	assertion.NotEqual(model.Database(), copiedModel.Database())
}

func Test_ModelC(t *testing.T) {
	assertion := assert.New(t)
	model := NewModel(newMockConfig(), defaultLogger)

	db := model.C("testing_collection")
	assertion.NotNil(db.collection)
	assertion.Equal(model.Database(), db.Database())
}

func Test_ModelQuery(t *testing.T) {
	assertion := assert.New(t)
	model := NewModel(newMockConfig(), defaultLogger)
	test := &mockTestModel{bson.NewObjectId(), "testing"}

	model.Query("testing_collection", mockTestModelIndexes, func(c *mgo.Collection) {
		err := c.Insert(test)
		assertion.Nil(err)
	})
}

func Benchmark_ModelQuery(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testModel := &mockTestModel{bson.NewObjectId(), "testing"}

		mockModel.Query("testing_collection", mockTestModelIndexes, func(c *mgo.Collection) {
			c.Insert(testModel)
		})
	}
}

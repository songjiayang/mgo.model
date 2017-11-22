package models

import (
	"github.com/songjiayang/mgo.model"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type _Simple struct{}

var (
	Simple *_Simple

	simpleCollection = "simples"
	simpleIndexes    = []mgo.Index{}
)

type SimpleModel struct {
	ID bson.ObjectId `bson:"_id"`

	Data  string `bson:"data1"`
	Data2 string `bson:"data2"`
}

func NewSimpleModel() *SimpleModel {
	return &SimpleModel{
		ID: bson.NewObjectId(),
	}
}

func (_ *_Simple) Find(id string) (sm *SimpleModel, err error) {
	if !bson.IsObjectIdHex(id) {
		return nil, model.ErrInvalidId
	}

	Simple.Query(func(c *mgo.Collection) {
		err = c.FindId(bson.ObjectIdHex(id)).One(&sm)
	})

	return
}

func (_ *_Simple) Insert(sm *SimpleModel) (err error) {
	Simple.Query(func(c *mgo.Collection) {
		err = c.Insert(sm)
	})
	return
}

func (_ *_Simple) Query(query func(c *mgo.Collection)) {
	Model().Query(simpleCollection, simpleIndexes, query)
}

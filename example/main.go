package main

import (
	"fmt"

	"github.com/dolab/gogo"
	"github.com/songjiayang/mgo.model"
	"github.com/songjiayang/mgo.model/example/models"
)

var mgo *model.Model

func main() {
	cfg := (`{
    "host": "localhost:27017",
      "user": "",
      "password": "",
      "database": "mgomodel",
      "mode": "Strong",
      "pool": 5,
      "timeout": 5
  }`)

	mgoConfig, _ := model.NewConfig([]byte(cfg))
	logger := gogo.NewAppLogger("stdout", "")

	models.NewWithConfig(mgoConfig, logger)

	sm := models.NewSimpleModel()
	sm.Data = "data1"
	sm.Data2 = "data2"
	err := models.Simple.Insert(sm)
	fmt.Println(err) // should be nil

	foundModel, _ := models.Simple.Find(sm.ID.Hex())

	fmt.Println(foundModel.Data)  // should be data1
	fmt.Println(foundModel.Data2) // should be data2
}

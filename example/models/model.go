package models

import (
	"github.com/dolab/gogo"
	"github.com/songjiayang/mgo.model"
)

var mongo *model.Model

func NewWithConfig(cfg *model.Config, logger gogo.Logger) {
	mongo = model.NewModel(cfg, logger)
}

func Model() *model.Model {
	return mongo
}

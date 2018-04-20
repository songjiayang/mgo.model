package model

import (
	"errors"

	"github.com/globalsign/mgo"
)

var (
	ErrNotFound     = mgo.ErrNotFound
	ErrNotPersisted = errors.New("Record has not persisted!")
	ErrInvalidId    = errors.New("Invalid BSON object id!")
	ErrInvalidArgs  = errors.New("Invalid arguments of the query method!")
)

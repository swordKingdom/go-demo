package api

import "github.com/globalsign/mgo"

type MgoClientPool struct {
	mgo.Database
}

func NewMongoDBPoolWithConf(conf interface{})*MgoClientPool{

}
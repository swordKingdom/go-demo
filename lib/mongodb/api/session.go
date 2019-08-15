package api

import (
	"github.com/globalsign/mgo"
	"this_is_a_explame/lib"
)

type MongoSession struct {
	*mgo.Session
	conf *lib.MgoConf
}

func (s *MongoSession) Find() {
	s.DB(s.conf.Database)
}

func NewMongoDBSessionWithConf(conf *lib.MgoConf) (*MongoSession, error) {
	cloneConf := conf.Cype()
	dialInfo := &mgo.DialInfo{
		Username:      cloneConf.Username,
		Password:      cloneConf.Password,
		Addrs:         cloneConf.Addrs,
		Database:      cloneConf.Database,
		Timeout:       cloneConf.Timeout,
		MaxIdleTimeMS: cloneConf.MaxIdleTimeMS,
		PoolLimit:     cloneConf.PoolLimit,
	}
	session, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		return nil, err
	}
	return &MongoSession{session, cloneConf}, nil
}

package mongodb

import (
	"fmt"

	"this_is_a_explame/lib/configloadder/conf"
	"this_is_a_explame/lib/db/mongodb/api"
)

var mgoConnMap = make(map[string]*api.MongoSession)

func Mongo(key string) (*api.MongoSession, error) {
	if conn, ok := mgoConnMap[key]; ok {
		return conn, nil
	} else {
		return nil, fmt.Errorf("the db %v is not exist ", key)
	}
}

func init() {
	userName := conf.GlobalConf.GetString("userName", "hhh")
	mgoConf := &api.MgoConf{
		Username: userName,
	}
	session, err := api.NewMongoDBSessionWithConf(mgoConf)
	if err == nil {
		mgoConnMap[""] = session
	}
}

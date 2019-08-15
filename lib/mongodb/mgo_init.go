package mongodb

import (
	"fmt"
	"this_is_a_explame/lib"
	"this_is_a_explame/lib/configloader"
)

var mgoConnMap = make(map[string]*lib.MongoSession)

func Mongo(key string) (*lib.MongoSession, error) {
	if conn, ok := mgoConnMap[key]; ok {
		return conn, nil
	} else {
		return nil, fmt.Errorf("the db %v is not exist ", key)
	}
}

func init() {
	userName := configloader.GlobalConf.GetString("userName", "hhh")
	mgoConf := &lib.MgoConf{
		Username: userName,
	}
	session, err := lib.NewMongoDBSessionWithConf(mgoConf)
	if err == nil {
		mgoConnMap[""] = session
	}
}

package mongodb

import (
	"fmt"

	"this_is_a_explame/api"
)

var mgoClient = make(map[string]*api.MgoClientPool,0)

func GetMgoClient(key string)(*api.MgoClientPool,error){
	if cli,ok :=mgoClient[key];ok{
		return cli,nil
	}else{
		return nil,fmt.Errorf("the dbs %v is not't exist",key)
	}
}

func init(){
	pool := api.NewMongoDBPoolWithConf("")
	mgoClient[""] = pool
}
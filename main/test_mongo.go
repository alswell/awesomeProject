package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func main() {
	s, e := mgo.Dial("172.17.0.2:27017")
	fmt.Printf("%#v, %v\n", s, e)
	//change := mgo.Change{
	//	Update: bson.M{"$inc": bson.M{"n": 1}},
	//	ReturnNew: true,
	//}
	change := mgo.Change{
		Update:    bson.M{"name": "zzz"},
		Upsert:    true,
		ReturnNew: true,
	}
	info, err := s.DB("local").C("mytest").Find(bson.M{"age": 27}).Apply(change, nil)
	fmt.Printf("%#v, %v\n", info, err)
}

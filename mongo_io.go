package main

import (
		"gopkg.in/mgo.v2"
		"gopkg.in/mgo.v2/bson"
	"strconv"
)

func mongo_i(session_id string, name string, value string){
        
        session, err := mgo.Dial("vpn.rebirtharmitage.com:21701")
        if err != nil {
                panic(err)
        }
        defer session.Close()

        // Optional. Switch the session to a monotonic behavior.
        session.SetMode(mgo.Monotonic, true)
        
        c := session.DB("intatl").C(session_id)

        c.Insert(bson.M{"key": name, "value": value})
}

func mongo_j(session_id string, id string, value geocode){
	    session, err := mgo.Dial("vpn.rebirtharmitage.com:21701")
        if err != nil {
                panic(err)
        }
        defer session.Close()

        // Optional. Switch the session to a monotonic behavior.
        session.SetMode(mgo.Monotonic, true)
        
        c := session.DB("intatl").C(session_id)
		a := strconv.FormatFloat(value.lat, 'f', -1, 64) 
		b := strconv.FormatFloat(value.lng, 'f', -1, 64)
	
        c.Insert(bson.M{"key": id, "value": []string{a, b}})
}
package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"strconv"
	"log"
)

type ServiceProvider struct{
	ProviderName, ProviderURL string
	Service []float32
}

func mongo_o(session_id string)([]ServiceProvider){
	//var sp []ServiceProvider
	
	session, err := mgo.Dial("vpn.rebirtharmitage.com:21701")
	if err != nil {
			panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("intatl").C(session_id)

	
	result := []ServiceProvider{}
	err = c.Find(bson.M{}).All(&result)
	if err != nil {
			log.Fatal(err)
	}

	return result
}

func mongo_i(session_id string, w Wireless){
        
	session, err := mgo.Dial("vpn.rebirtharmitage.com:21701")
	if err != nil {
			panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("intatl").C(session_id)
	var sp ServiceProvider
	for i := range w.Results.WirelineServices {
		sp.ProviderName = w.Results.WirelineServices[i].ProviderName
		sp.ProviderURL = w.Results.WirelineServices[i].ProviderURL
		for j := range w.Results.WirelineServices[i].Technologies{
			sp.Service = append(sp.Service, w.Results.WirelineServices[i].Technologies[j].TechnologyCode)
			sp.Service = append(sp.Service, w.Results.WirelineServices[i].Technologies[j].DownloadQuality)
			sp.Service = append(sp.Service, w.Results.WirelineServices[i].Technologies[j].MaximumAdvertisedDownloadSpeed)
			sp.Service = append(sp.Service, w.Results.WirelineServices[i].Technologies[j].MaximumAdvertisedUploadSpeed)
		}
		c.Insert(sp)
		sp.Service = sp.Service[:0]
	}
	
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
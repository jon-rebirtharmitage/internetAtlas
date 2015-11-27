 package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"strconv"
	"fmt"
	"log"
)

type ServiceProvider struct{
	ProviderName, ProviderURL string
	Service []float32
}

type ServiceList struct {
	Name string
	Upper float32
	Sp []ServiceProvider
}

func mongo_o(session_id string)([]ServiceList){
	//var sp []ServiceProvider
	
	session, err := mgo.Dial("vpn.rebirtharmitage.com:21701")
	if err != nil {
			panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	//session.SetMode(mgo.Monotonic, true)

	c := session.DB("intatl").C(session_id)

	result := []ServiceList{}
	iter := c.Find(nil).Limit(100).Iter()
	err = iter.All(&result)
	if err != nil {
			log.Fatal(err)
	}
	
	fmt.Println(result)
	return result
	
}

func mongo_i(session_id string, sig Signal){
        
	session, err := mgo.Dial("vpn.rebirtharmitage.com:21701")
	if err != nil {
			panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("intatl").C(session_id)
	
	var top float32
	top = 0.0
	
	for h := range sig.W {
			var spl []ServiceProvider
			var sp ServiceProvider

			for i := range sig.W[h].Results.WirelineServices {
				sp.ProviderName = sig.W[h].Results.WirelineServices[i].ProviderName
				sp.ProviderURL = sig.W[h].Results.WirelineServices[i].ProviderURL
				for j := range sig.W[h].Results.WirelineServices[i].Technologies{
					sp.Service = append(sp.Service, sig.W[h].Results.WirelineServices[i].Technologies[j].TechnologyCode)
					sp.Service = append(sp.Service, sig.W[h].Results.WirelineServices[i].Technologies[j].DownloadQuality)
					sp.Service = append(sp.Service, sig.W[h].Results.WirelineServices[i].Technologies[j].MaximumAdvertisedDownloadSpeed)
					sp.Service = append(sp.Service, sig.W[h].Results.WirelineServices[i].Technologies[j].MaximumAdvertisedUploadSpeed)
					if (sig.W[h].Results.WirelineServices[i].Technologies[j].MaximumAdvertisedDownloadSpeed > top){
						top = sig.W[h].Results.WirelineServices[i].Technologies[j].MaximumAdvertisedDownloadSpeed
					}
				}
				spl = append(spl, sp)
				sp.Service = sp.Service[:0]
			}

		c.Insert(bson.M{"name":sig.id, "upper":top , "sp": spl})
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
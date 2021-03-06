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
	TechnologyCode, Download, Upload, Rating float32
	TechnologyCode_S, Download_S, Upload_S, Rating_S string
}

type ServiceList struct {
	Name string
	Upper float32
	Sp []ServiceProvider
}

func mongo_o(session_id string)([]ServiceList){
	//var sp []ServiceProvider
	
	session, err := mgo.Dial("vps.rebirtharmitage.com:21701")
	if err != nil {
			panic(err)
	}
	defer session.Close()

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
        
	session, err := mgo.Dial("vps.rebirtharmitage.com:21701")
	if err != nil {
			panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	c := session.DB("intatl").C(session_id)
	
	var top float32
	top = 0.0
	var CurrentTechCode float32
	CurrentTechCode = 0.0
	
	for h := range sig.W {
			//var spl []ServiceProvider
			var sp ServiceProvider
			for i := range sig.W[h].Results.WirelineServices {
				sp.ProviderName = sig.W[h].Results.WirelineServices[i].ProviderName
				sp.ProviderURL = sig.W[h].Results.WirelineServices[i].ProviderURL
				for j := range sig.W[h].Results.WirelineServices[i].Technologies{
					if (sig.W[h].Results.WirelineServices[i].Technologies[j].TechnologyCode > CurrentTechCode){
						sp.TechnologyCode = sig.W[h].Results.WirelineServices[i].Technologies[j].TechnologyCode
						sp.TechnologyCode_S = TechCode(sig.W[h].Results.WirelineServices[i].Technologies[j].TechnologyCode)
						sp.Download = sig.W[h].Results.WirelineServices[i].Technologies[j].MaximumAdvertisedDownloadSpeed
						sp.Download_S = DownCode(sig.W[h].Results.WirelineServices[i].Technologies[j].MaximumAdvertisedDownloadSpeed)
						sp.Upload = sig.W[h].Results.WirelineServices[i].Technologies[j].MaximumAdvertisedUploadSpeed
						sp.Upload_S = DownCode(sig.W[h].Results.WirelineServices[i].Technologies[j].MaximumAdvertisedUploadSpeed)
						sp.Rating = 0
						sp.Rating_S = "Placeholder"
						if (sig.W[h].Results.WirelineServices[i].Technologies[j].MaximumAdvertisedDownloadSpeed > top){
							top = sig.W[h].Results.WirelineServices[i].Technologies[j].MaximumAdvertisedDownloadSpeed
						}
						CurrentTechCode = sig.W[h].Results.WirelineServices[i].Technologies[j].TechnologyCode
					}
				}
				Extend(sp)
				CurrentTechCode = 0
			}
		c.Insert(bson.M{"name":sig.id, "upper":top , "sp": SPL})
		top = 0
	}
	SPL = SPL[:0]
}

var SPL []ServiceProvider

func Extend(sp ServiceProvider) {
	fmt.Println(sp)
	SPL = append(SPL, sp)
}

func TechCode(i float32) (string){
	if i == 30 {
		return "DSL"
	}else if i == 40 {
		return "Cable"
	}else if i == 41 {
		return "Fiber"
	}else if i == 50 {
		return "Cable"
	}else if i == 10 {
		return "DSL"
	}
	return "Unknown"
}

func DownCode(i float32) (string){
	if i == 1{
		return "Up to  200 kbps"
	}else if i == 2{
		return "Up to 768 kbps"
	}else if i == 3{
		return "Up to 1.5 mbps"
	}else if i == 4{
		return "Up to 3 mbps"
	}else if i == 5{
		return "Up to 6 mbps"
	}else if i == 6{
		return "Up to 10 mbps"
	}else if i == 7{
		return "Up to 25 mbps"
	}else if i == 8{
		return "Up to 50 mbps"
	}else if i == 9{
		return "Up to 100 mbps"
	}else if i == 10{
		return "Up to 1 gbps"
	}else if i == 11{
		return "Over 1 gbps"
	}
	return "Unknown"
}

func mongo_j(session_id string, id string, value geocode){
	    session, err := mgo.Dial("vps.rebirtharmitage.com:21701")
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
package main

import (
	"net/http"
	"io/ioutil"
	"io"
	"encoding/json"
	"strings"
	"strconv"
	"log"
	"fmt"
)

type geocode struct{
	lat, lng float64
}

type Wireless struct {
	Results struct {
		BroadbandSource struct {
			Organization    string `json:"organization"`
			OrganizationURL string `json:"organizationURL"`
			StateFips       string `json:"stateFips"`
		} `json:"broadbandSource"`
		WirelineServices []struct {
			BlockFIPS            string `json:"blockFIPS"`
			DoingBusinessAs      string `json:"doingBusinessAs"`
			Frn                  string `json:"frn"`
			HoldingCompanyName   string `json:"holdingCompanyName"`
			HoldingCompanyNumber string `json:"holdingCompanyNumber"`
			ProviderName         string `json:"providerName"`
			ProviderURL          string `json:"providerURL"`
			Technologies         []struct {
				DownloadQuality                float32 `json:"downloadQuality"`
				MaximumAdvertisedDownloadSpeed float32 `json:"maximumAdvertisedDownloadSpeed"`
				MaximumAdvertisedUploadSpeed   float32 `json:"maximumAdvertisedUploadSpeed"`
				MaximumDownloadScore           float32 `json:"maximumDownloadScore"`
				MaximumProviderScore           float32 `json:"maximumProviderScore"`
				MaximumSpeedScore              float32 `json:"maximumSpeedScore"`
				MaximumTechnologyScore         float32 `json:"maximumTechnologyScore"`
				MaximumUploadScore             float32 `json:"maximumUploadScore"`
				OverallQuality                 float32 `json:"overallQuality"`
				ProviderQuality                float32 `json:"providerQuality"`
				SpeedQuality                   float32 `json:"speedQuality"`
				TechnologyCode                 float32 `json:"technologyCode"`
				TechnologyQuality              float32 `json:"technologyQuality"`
				TypicalDownloadSpeed           float32 `json:"typicalDownloadSpeed"`
				TypicalUploadSpeed             float32 `json:"typicalUploadSpeed"`
				UploadQuality                  float32 `json:"uploadQuality"`
			} `json:"technologies"`
		} `json:"wirelineServices"`
	} `json:"Results"`
	Message      []interface{} `json:"message"`
	ResponseTime int           `json:"responseTime"`
	Status       string        `json:"status"`
}


func geocoding(session_id string, input address) (geocode){
	url := "https://maps.googleapis.com/maps/api/geocode/json?address=" + input.street + "," + input.city + "," + input.state + "&key=AIzaSyAd4WHqblWQ2ac4JMf0yOZfBsIkvOlKRQo"

	req, _ := http.NewRequest("GET", url, nil)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	
	type GeoData struct {
		Results []struct {
			AddressComponents []struct {
				LongName  string   `json:"long_name"`
				ShortName string   `json:"short_name"`
				Types     []string `json:"types"`
			} `json:"address_components"`
			FormattedAddress string `json:"formatted_address"`
			Geometry         struct {
				Location struct {
					Lat float64 `json:"lat"`
					Lng float64 `json:"lng"`
				} `json:"location"`
				LocationType string `json:"location_type"`
				Viewport     struct {
					Northeast struct {
						Lat float64 `json:"lat"`
						Lng float64 `json:"lng"`
					} `json:"northeast"`
					Southwest struct {
						Lat float64 `json:"lat"`
						Lng float64 `json:"lng"`
					} `json:"southwest"`
				} `json:"viewport"`
			} `json:"geometry"`
			PlaceID string   `json:"place_id"`
			Types   []string `json:"types"`
		} `json:"results"`
		Status string `json:"status"`
	}
	n := string(body[:])
	
	dec := json.NewDecoder(strings.NewReader(n))
	
	var geo geocode
	
	for {
		var a GeoData
		if err := dec.Decode(&a); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		
		for k := range a.Results{
			fmt.Println(a.Results[k].Geometry.Location.Lat)
			fmt.Println(a.Results[k].Geometry.Location.Lng)
			geo.lat = a.Results[k].Geometry.Location.Lat
			geo.lng = a.Results[k].Geometry.Location.Lng
			//mongo_j()
			return geo
		}
	}
	return geo
}

func wireServiceCall(input geocode, session_id string){

	a := strconv.FormatFloat(input.lat, 'f', -1, 64)
	b := strconv.FormatFloat(input.lng, 'f', -1, 64)
	url := "http://www.broadbandmap.gov/broadbandmap/broadband/jun2014/wireline?latitude=" + a + "&longitude=" + b + "&format=json"
	fmt.Println(url)
	req, _ := http.NewRequest("GET", url, nil)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	
	type Message struct {
		Status string
		Results map[string]interface{}
	}

	n := string(body[:])
	
	dec := json.NewDecoder(strings.NewReader(n))
	
	//var W []Wireless
	
	for {
		var w Wireless
		if err := dec.Decode(&w); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		mongo_i(session_id, w)
		//W = append(W, w)
	}
}
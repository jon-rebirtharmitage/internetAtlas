package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func main() {

	url := "http://www.broadbandmap.gov/broadbandmap/broadband/jun2014/wireless?latitude=40.7195898&longitude=-73.9998334&format=json"

	req, _ := http.NewRequest("GET", url, nil)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}
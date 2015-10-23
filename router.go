package main

import (
    "fmt"
    "github.com/gorilla/mux"
	"html/template"
    "net/http"
)

type Page struct {
	Title string
	Body  string
}

func (p *Page) GOFUNC(){
	fmt.Println("HOLY FREAKING CRAP")
}

func Index(w http.ResponseWriter, r *http.Request) {
    //a := wirelessServiceCall()
	//b := geocoding()
    //mongo_i("Test", "Holy shit", "Did this work.")
	p, _ := loadPage_Index("Awesomesauce")
    t, _ := template.ParseFiles("index.html")
    t.Execute(w, p)
    //fmt.Fprint(w, "Welcome!\n")
}


func Hello(w http.ResponseWriter, r *http.Request) {
/* 	m := parseAddresses(ps.ByName("name"))
	for k := range m{
		b := geocoding(m[k])
		c := wirelessServiceCall(b)
		fmt.Println(c)
	} */
    //fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func loadPage_Index(title string) (*Page, error){
        return &Page{Title: title, Body: "blank..."}, nil
}

var router = mux.NewRouter()

func main() {
    router.HandleFunc("/", Index)
    router.HandleFunc("/hello", Hello).Methods("GET")
	
    router.PathPrefix("/").Handler(http.FileServer(http.Dir("../internetatlas/")))
	
	http.Handle("/", router)
	http.ListenAndServe(":8081", nil)
}
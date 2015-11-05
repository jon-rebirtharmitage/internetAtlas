package main

import (
	"fmt"
	//"github.com/julienschmidt/httprouter"
	"github.com/gorilla/mux"
	"html/template"
    "net/http"
	"time"
	"math/rand"
	//"strconv"
)

type Page struct {
	Title string
	Body  string
	R []ServiceProvider
}

func Index(w http.ResponseWriter, r *http.Request) {
	p, _ := loadPage_Index("Awesome")
    renderTemplate(w, "./html/index", p)
}

func CreateSessionID() (string){
	source := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKMNOPQRSTUVWXYZ"
	rand.Seed(time.Now().UnixNano())
	s := ""
	for i := 0; i < 24; i++{
		s = s + string(source[rand.Intn(50)])
	}
	return s
}


func Process(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	s := CreateSessionID()
	vars := mux.Vars(r)
	fmt.Println(s)
	m := parseAddresses(vars["Value"])
	for k := range m{
		b := geocoding(s, m[k])
		wireServiceCall(b, s)
	}
	fmt.Fprint(w, s)
}

func CreateResults(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
	vars := mux.Vars(r)
	s := vars["Session"]
	result := mongo_o(s)
	p, _ := loadPage_Results("OMFG", result)
    renderTemplate(w, "./html/results", p)
} 

func loadPage_Results(session_id string, results []ServiceProvider) (*Page, error){
	return &Page{Title: session_id, R: results}, nil
}

func loadPage_Index(title string) (*Page, error){
    return &Page{Title: title, Body: "blank..."}, nil
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
    t, err := template.ParseFiles(tmpl + ".html")
	if err != nil{
		fmt.Println("THIS DONE FUCKED UP!", err)
	}
    t.Execute(w, p)
}

/*
	Start of ROUTER Section
*/
var router = mux.NewRouter()

func main() {
	router.HandleFunc("/", Index)
	router.HandleFunc("/Process/{Value}", Process)
	router.HandleFunc("/Results/{Session}", CreateResults)

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("../internetatlas/")))
	
	http.Handle("/", router)
	http.ListenAndServe(":8081", nil)
}
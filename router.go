package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
  "net/http"
	"time"
	"math/rand"
)

type Page struct {
	Title string
	Body  string
	R []ServiceList
}

func Index(w http.ResponseWriter, r *http.Request) {
	p, _ := loadPage_Index("Internet Atlas IO")
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

func email(response http.ResponseWriter, request *http.Request) {
	name := request.FormValue("name")
	if (name == ""){ name = "NO NAME"}
	email := request.FormValue("email")
	if (email == ""){ email = "NO EMAIL"}
	phone := request.FormValue("phone")
	if (phone == ""){ phone = "NO PHONE"}
	body := request.FormValue("body")
	if (body == ""){ body = "NO MESSAGE"}
	a := sendMail([]string{"jon@rebirtharmitage.com"}, email + " : " + phone, body)
	if (a == nil){
		redirectTarget := "/"
		http.Redirect(response, request, redirectTarget, 302)
	}else{
		redirectTarget := "/"
		http.Redirect(response, request, redirectTarget, 302)
	}
}

func CreateResults(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
	vars := mux.Vars(r)
	s := vars["Session"]
	result := mongo_o(s)
	p, _ := loadPage_Results("OMFG", result)
  renderTemplate(w, "./html/results", p)
} 

func DisplayDetails(w http.ResponseWriter, r *http.Request){
		p, _ := loadPage_Details("Awesome")
    renderTemplate(w, "./html/details", p)
}

func loadPage_Details(title string) (*Page, error){
	return &Page{Title: title, Body: "blank..."}, nil
}

func loadPage_Results(session_id string, results []ServiceList) (*Page, error){
	return &Page{Title: session_id, R: results}, nil
}

func loadPage_Index(title string) (*Page, error){
    return &Page{Title: title, Body: "blank..."}, nil
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
  t, err := template.ParseFiles(tmpl + ".html")
	if err != nil{
		fmt.Println("Error in loading template file.", err)
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
	router.HandleFunc("/Details", DisplayDetails)
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("../internetatlas/")))
	http.Handle("/", router)
	http.ListenAndServe(":8080", nil)
}
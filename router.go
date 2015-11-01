package main

import (
    "fmt"
	"github.com/julienschmidt/httprouter"
	"html/template"
    "net/http"
	"time"
	"math/rand"
)

type Page struct {
	Title string
	Body  string
}

type rPage struct {
	Title string
	r []string
}

func Index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	p, _ := loadPage_Index("Awesomesauce")
    t, _ := template.ParseFiles("./html/index.html")
    t.Execute(w, p)
}

func CreateSessionID() (string){
	source := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKMNOPQRSTUVWXYZ"
	rand.Seed(time.Now().UnixNano())
	s := ""
	for i := 0; i < 24; i++{
		s = s + string(source[rand.Intn(52)])
	}
	return s
}


func Process(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	s := CreateSessionID()
	fmt.Println(s)
	m := parseAddresses(ps.ByName("Value"))
	for k := range m{
		b := geocoding(s, m[k])
		wireServiceCall(b, s)
	} 
	CreateResults(s)
}

func CreateResults(s string) {
	r := mongo_o(s)
	for i := range r{
		fmt.Println(r[i].Service)
	}
} 

func loadPage_Index(title string) (*Page, error){
    return &Page{Title: title, Body: "blank..."}, nil
}

func main() {
	router := httprouter.New()
    router.GET("/", Index)
	router.GET("/Process/:Value", Process)
	
    router.ServeFiles("/js/*filepath", http.Dir("/home/jcronin/internetatlas/js"))
	router.ServeFiles("/css/*filepath", http.Dir("/home/jcronin/internetatlas/css"))
	router.ServeFiles("/fonts/*filepath", http.Dir("/home/jcronin/internetatlas/fonts"))
	router.ServeFiles("/img/*filepath", http.Dir("/home/jcronin/internetatlas/img"))
	
	http.Handle("/", router)
	http.ListenAndServe(":8081", nil)
}
package main

import (
    "fmt"
    "github.com/julienschmidt/httprouter"
	"html/template"
    "net/http"
    "log"
)

type Page struct {
	Title string
	Body  string
}

func (p *Page) GOFUNC(){
	fmt.Println("HOLY FREAKING CRAP")
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    //a := wirelessServiceCall()
	//b := geocoding()
    //mongo_i("Test", "Holy shit", "Did this work.")
	p, _ := loadPage_Index("Awesomesauce")
    t, _ := template.ParseFiles("index.html")
    t.Execute(w, p)
    //fmt.Fprint(w, "Welcome!\n")
}


func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	m := parseAddresses(ps.ByName("name"))
	for k := range m{
		b := geocoding(m[k])
		c := wirelessServiceCall(b)
		fmt.Println(c)
	}
    //fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func loadPage_Index(title string) (*Page, error){
        return &Page{Title: title, Body: "blank..."}, nil
}

func main() {
    router := httprouter.New()
    router.GET("/", Index)
    router.GET("/hello/:name", Hello)
	
    log.Fatal(http.ListenAndServe(":8081", router))
}
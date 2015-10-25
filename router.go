package main

import (
    "fmt"
	"github.com/julienschmidt/httprouter"
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

func Index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    //a := wirelessServiceCall()
	//b := geocoding()
    //mongo_i("Test", "Holy shit", "Did this work.")
	p, _ := loadPage_Index("Awesomesauce")
    t, _ := template.ParseFiles("index.html")
    t.Execute(w, p)
    //fmt.Fprint(w, "Welcome!\n")
}


func Process(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	m := parseAddresses(ps.ByName("Value"))
	for k := range m{
		b := geocoding(m[k])
		c := wireServiceCall(b)
		fmt.Println(c)
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
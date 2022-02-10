package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/sessions"
)

var (
	key   = []byte("sssshhhhh-dont-leak-me")
	store = sessions.NewCookieStore(key)
)

func main() {
	http.HandleFunc("/", HelloWorld)
	http.ListenAndServe(":8080", nil)
}

type PageData struct {
	XList []int
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	error := r.ParseForm()
	if error != nil {
		fmt.Errorf("Ahhhhh")
	}

	session, _ := store.Get(r, "my-session")
	xObj := session.Values["x"]

	var x int
	if xObj == nil {
		x = 0
	} else {
		x = xObj.(int)
	}

	addX := r.Form.Get("addX")
	minusX := r.Form.Get("subX")

	if addX != "" {
		session.Values["x"] = x + 1
		x = x + 1
		session.Save(r, w)
	}

	if minusX != "" && x > 0 {
		session.Values["x"] = x - 1
		x = x - 1
		session.Save(r, w)
	}

	xList := make([]int, x)
	for i := 0; i < x; i++ {
		xList[i] = i
	}

	fmt.Println(x)

	data := PageData{
		XList: xList,
	}

	tmpl := template.Must(template.ParseFiles("index.html"))
	tmpl.Execute(w, data)
}

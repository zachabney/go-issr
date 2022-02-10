package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/kataras/go-sessions/v3"
)

var session = sessions.New(sessions.Config{Cookie: "go-session"})

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
	s := session.Start(w, r)

	x, error := s.GetInt("x")
	if error != nil {
		x = 0
	}
	if r.Form.Get("addX") != "" {
		x = x + 1
	}
	if r.Form.Get("subX") != "" && x > 0 {
		x = x - 1
	}
	s.Set("x", x)

	xList := make([]int, x)
	for i := 0; i < x; i++ {
		xList[i] = i
	}

	data := PageData {
		XList: xList,
	}

	tmpl := template.Must(template.ParseFiles("index.html"))
	tmpl.Execute(w, data)
}

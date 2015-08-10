package main

import (
	"fmt"
	"io"
	"net/http"
)

type String string

type Struct struct {
	Greeting string
	Punct    string
	Who      string
}

func (s String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, string(s))
}

func (s *Struct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, s.Greeting+s.Punct+s.Who)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	url := r.URL
	params := url.Query()
	for v := range params {
		fmt.Println(v)
	}
	io.WriteString(w, params.Get("foo"))
}

func main() {
	http.Handle("/string", String("Im a "))
	http.Handle("/struct", &Struct{"Hello", ":", "Gophers!"})
	http.HandleFunc("/foo", HelloServer)
	//	http.Handle("/string", String("I'm a frayed knot"))
	http.ListenAndServe("localhost:4000", nil)
}

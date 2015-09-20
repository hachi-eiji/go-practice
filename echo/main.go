package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
)

func allowGet(r *http.Request) bool {
	return GET.String() == r.Method
}

func pong(w http.ResponseWriter, r *http.Request) {
	if allowGet(r) {
		w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, "pong")
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func increment(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()

	var result = make(map[string]int)
	for k, v := range params {
		i, err := strconv.Atoi(v[0])
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		result[k] = i + 1
	}
	json, err := json.Marshal(result)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, string(json))
	}
}

func main() {
	http.HandleFunc("/ping", pong)
	http.HandleFunc("/add", increment)
	if err := http.ListenAndServe("localhost:8000", nil); err != nil {
		log.Fatalf("cannot start server. ", err)
	}
}

package main

import (
	"fmt"
	"net/http"

	"github.com/convox/kernel/Godeps/_workspace/src/github.com/codegangsta/negroni"
	"github.com/convox/kernel/Godeps/_workspace/src/github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("ok"))
	})

	router.HandleFunc("/check", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("ok"))
	})

	n := negroni.New()

	n.UseHandler(router)

	n.Run(fmt.Sprintf(":5000"))
}

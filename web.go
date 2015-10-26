package main

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("ok"))
	})

	router.HandleFunc("/check", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("ok"))
	})

	router.HandleFunc("/slack/release", func(rw http.ResponseWriter, r *http.Request) {
		r.ParseMultipartForm(2048)

		if r.Form.Get("token") != os.Getenv("SLACK_INCOMING_WEBHOOK_TOKEN") {
			http.Error(rw, "invalid token", 401)
			return
		}

		args := strings.Split(r.Form.Get("text"), " ")

		if len(args) < 1 {
			http.Error(rw, "please specify a command (create, publish)", 403)
			return
		}

		switch args[0] {
		case "create":
			branch := "master"

			if len(args) > 1 {
				branch = args[1]
			}

			cmd := exec.Command("bin/create", branch)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			go cmd.Run()

			rw.Write([]byte(fmt.Sprintf("creating release from `%s`", branch)))
		case "publish":
			http.Error(rw, "not yet supported", 403)
		case "":
			http.Error(rw, "please specify a command (create, publish)", 403)
		default:
			http.Error(rw, fmt.Sprintf("invalid command: %s", args[0]), 403)
		}
	})

	n := negroni.New()

	n.UseHandler(router)

	n.Run(fmt.Sprintf(":5000"))
}

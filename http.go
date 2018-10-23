package main

import (
	"fmt"
	"net/http"

	"encoding/json"

	"github.com/urfave/cli"
)

// DymServer starts the http server for corrections.
func DymServer(c *cli.Context) error {
	port := c.String("port")
	http.HandleFunc("/corrections", func(w http.ResponseWriter, r *http.Request) {
		name := r.FormValue("name")
		if len(name) == 0 {
			w.WriteHeader(400)
			fmt.Fprintf(w, "must provide 'name' parameter")
			return
		}
		corrections := Corrections(name)
		js, err := json.Marshal(corrections)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "error generating corrections %v", err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	})

	return http.ListenAndServe(":"+port, nil)
}

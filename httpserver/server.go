package httpserver

import (
	"fmt"
	"net/http"

	"github.com/faisal097/employee-manager/config"
)

func Serve(c *config.HttpServerConfig) error {
	http.HandleFunc("/healthz", getHealth)
	http.HandleFunc("/get/{name}", get)
	http.HandleFunc("/set", set)
	http.HandleFunc("/delete", delete)
	http.HandleFunc("/update/{name}", update)

	addr := fmt.Sprintf(":%s", c.GetPort())
	return http.ListenAndServe(addr, nil)
}

func getHealth(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Fprintf(w, "OK")
		return
	}
	fmt.Fprintf(w, "Bad method")
}

func get(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Fprintf(w, "Not implemented")
		return
	}
	fmt.Fprintf(w, "Bad method")
}

func set(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		fmt.Fprintf(w, "Not implemented")
		return
	}
	fmt.Fprintf(w, "Bad method")
}

func update(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		fmt.Fprintf(w, "Not implemented")
		return
	}
	fmt.Fprintf(w, "Bad method")
}

func delete(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		fmt.Fprintf(w, "Not implemented")
		return
	}
	fmt.Fprintf(w, "Bad method")
}

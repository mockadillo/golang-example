package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	if r.Method != "POST" {
		fmt.Fprintf(w, `{"success": false, "error": "Unsupported method"}`)
		return
	}

	if r.Header.Get("x-custom-header") == "" {
		fmt.Fprintf(w, `{"success": false, "error": "You haven't provided the right header value"}`)
		return
	}

	decoder := json.NewDecoder(r.Body)
	var result map[string]string
	err := decoder.Decode(&result)
	if err != nil {
		fmt.Fprintf(w, `{"success": false, "error": "%s"}`, err.Error())
		return
	}

	if result["bodyParam"] == "" {
		fmt.Fprintf(w, `{"success": false, "error": "You haven't provided the right body value"}`)
		return
	}

	fmt.Fprintf(w, `{"success": true, "result": "You have reached the end"}`)
}

func main() {
	http.HandleFunc("/route", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

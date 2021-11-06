package main

import (
	"log"
	"net/http"
	"path/filepath"
)

var path = filepath.FromSlash("")

func printUserData(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Println(err.Error())
		return
	}

	log.Printf(
		"first-name: %s\n last-name: %s\n email: %s\n password: %s\n",
		req.Form.Get("first-name"),
		req.Form.Get("last-name"),
		req.Form.Get("email"),
		req.Form.Get("password"),
	)
}

func main() {

	fs := http.FileServer(http.Dir(path))

	http.Handle("/", fs)
	http.HandleFunc("/main_page", printUserData)

	log.Println("Server is starting on port 9000...")
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Fatal(err)
	}
}

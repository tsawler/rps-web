package main

import (
	"encoding/json"
	"html/template"
	"log"
	"myapp/rps"
	"net/http"
	"strconv"
)

// homePage is a handler to display the web page.
func homePage(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index.html")
}

// playRound is a handler that decides the outcome of a round,
// and returns JSON indicating round result.
func playRound(w http.ResponseWriter, r *http.Request) {
	playerChoice, _ := strconv.Atoi(r.URL.Query().Get("c"))
	result := rps.PlayRound(playerChoice)

	out, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

// main is the main function, and entrypoint to the application.
func main() {
	// define our handlers
	http.HandleFunc("/play", playRound)
	http.HandleFunc("/", homePage)

	// start a web server
	log.Println("Starting web server on port 8080")
	http.ListenAndServe(":8080", nil)
}

// renderTemplate opens a template, parses it, and executes it.
// By rights, this should be cached, and not executed every time
// a request is made, but this is deliberately a very simple web app.
func renderTemplate(w http.ResponseWriter, page string) {
	t, err := template.ParseFiles(page)
	if err != nil {
		log.Println(err)
		return
	}

	err = t.Execute(w, nil)
	if err != nil {
		log.Println(err)
		return
	}
}

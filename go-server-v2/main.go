package main // Needed with any go file

import (
	// Go will automatically delete unused imports!!
	"fmt"
	"go-server-v2/logging"
	"log"
	"net/http"
	"time"
)

// Checks the form for any errors and also does the logging stuff.
func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Printf("ParseForm() err: %v", err)
		return
	}
	var inputPass string = r.FormValue("password")
	var auth bool
	if inputPass == "OLDNAVY" {
		auth = true
	} else {
		auth = false
	}
	logging.LogInformation(time.Now(), w, r, auth)

}

// Archaic helloHandler that needs to be re-vamped.
func helloHandler(w http.ResponseWriter, r *http.Request) {

	// Right now this is only executed if
	// the path is /hello, but we still add this.
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found. Nothing else.", http.StatusNotFound)
		return

	}
	if r.Method != "GET" {
		http.Error(w, "Method is not supported... this will be reported.", http.StatusNotFound)
		return
	}
}

// Where most of the operations will be done
func main() {
	// Tells it where to look for the index.html file, it already knows
	// it needs to find index.html
	fileserver := http.FileServer(http.Dir("./static"))

	http.Handle("/", fileserver)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	// Have to use Printf instead of Println because it doesn't add a new line
	fmt.Printf("Server is listening on port 8080\n")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

package main // Needed with any go file

import (
	// Go will automatically delete unused imports!!
	"fmt"
	"log"
	"math/rand" // Extra package for random numbers
	"net/http"
	"time" // Extra package for time
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Printf("ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful:   ")
	city := r.FormValue("city")
	cityTemperature := rand.Intn(100)
	state := r.FormValue("state")
	stateAvgTemp := rand.Intn(100)
	fmt.Fprintf(w, "City: %s is at %d", city, cityTemperature)
	fmt.Fprintf(w, "State: %s is at %d average temp.", state, stateAvgTemp)

}

// * is a pointer here.
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

	currentTime := time.Now()
	fmt.Fprintf(w, "Hello World! The time is %s", currentTime.String())
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

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

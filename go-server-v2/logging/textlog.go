// Includes logging functions that help identify who submitted the form.
// Currently is able to get a user's IP and write to a text file.
package logging

// Not all these imports may be needed?
import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

// Reuses the w and r from main.go, and gives us the IP of the client.
// May be better in the main.go file?
func GetIpAddress(w http.ResponseWriter, r *http.Request) (string, string) {
	ip := r.RemoteAddr
	xforward := r.Header.Get("X-Forwarded-For")
	return ip, xforward
}

// Main function that takes in a time.Now(), w and r from main.go.
// It checks for any errors and also does the writing to a pre-existing text file.
func LogInformation(timeIn time.Time, w http.ResponseWriter, r *http.Request, auth bool) {

	// Open the file for reading
	file, err := os.OpenFile("logs/log_info.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)

	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Get the IP's
	ip, xforward := GetIpAddress(w, r)

	// Check if the XF is empty
	if len(strings.TrimSpace(xforward)) == 0 {
		xforward = "No XF"
	}

	// Get time to PT
	loc, err := time.LoadLocation("America/Los_Angeles")

	if err != nil {
		fmt.Println(err)
		return
	}

	// Do the actual PT conversion
	timeIn = timeIn.In(loc)

	s := timeIn.Format("Monday, January 2, 2006 at 3:04 PM")

	var incident string
	if auth {
		incident = fmt.Sprintf("GRANT: %v by %v / %v!\n", s, ip, xforward)
	} else {
		incident = fmt.Sprintf("DENY: %v by %v / %v!\n", s, ip, xforward)
	}
	_, err = file.WriteString(incident)
	if err != nil {
		fmt.Println(err)
		return
	}
}

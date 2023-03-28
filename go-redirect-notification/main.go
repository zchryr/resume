package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// User request info struct.
type request struct {
	RemoteAddress string
	RemotePort string
	UserAgent string
}

// Sends request info a Tines webhook.
func tines(r request) {
	// Convert struct to JSON.
	json, _ := json.Marshal(r)

	request, _ := http.NewRequest("POST", "https://autumn-silence-4783.tines.com/webhook/220d2ef276eddb56edf8af42301cc72b/747a971570b79acdf6cd47829113b7dc", bytes.NewBuffer(json))
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	// HTTP client.
	client := &http.Client{}
	response, err	:= client.Do(request)

	if err != nil {
		log.Fatal(err)
		return
	}

	defer response.Body.Close()

	// Simple logging.
	log.Print("response status:", response.Status)
	log.Print("response headers:", response.Header)
	body, _ := ioutil.ReadAll(response.Body)
	log.Print("response body:", string(body))
}

// Captures request info and redirects to resume link.
func redirect(w http.ResponseWriter, r *http.Request) {
	remoteaddr := strings.Split(r.RemoteAddr, ":")

	request := request{
		RemoteAddress:remoteaddr[0],
		RemotePort: remoteaddr[1],
		UserAgent: r.UserAgent(),
	}

	tines(request)

	http.Redirect(w, r, "https://cdn.rohrbach.tech/Zachary-Rohrbach-Resume.pdf", http.StatusTemporaryRedirect)
}

// Main.
func main() {
	http.HandleFunc("/", redirect)

	err := http.ListenAndServe(":1337", nil)

	if err != nil {
		log.Fatal(err)
	}
}
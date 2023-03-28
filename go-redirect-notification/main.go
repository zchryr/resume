package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

// User request info struct.
type request struct {
	RemoteAddress string
	RemotePort string
	UserAgent string
}

// Sends request info a webhook.
func webhook(r request) {
	// Convert struct to JSON.
	json, _ := json.Marshal(r)

	request, _ := http.NewRequest("POST", os.Getenv("WEBHOOK"), bytes.NewBuffer(json))
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

	webhook(request)

	http.Redirect(w, r, "https://cdn.rohrbach.tech/Zachary-Rohrbach-Resume.pdf", http.StatusTemporaryRedirect)
}

// Main.
func main() {
	http.HandleFunc("/", redirect)

	err := http.ListenAndServe(":80", nil)

	if err != nil {
		log.Fatal(err)
	}
}
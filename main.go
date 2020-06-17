package main

import (
	"fmt"
	"net/http"
	log "github.com/sirupsen/logrus"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {

	page := r.URL.Path[1:]
	data, ok := r.URL.Query()["data"]
	if !ok {
		data = []string{"no data received"}
	}

	log.WithFields(log.Fields{
		"page": page,
		"data": data,
	}).Info("new request received")

}

func main() {

	f, err := os.OpenFile("/var/log/xss.json", os.O_APPEND | os.O_CREATE | os.O_RDWR, 0666)
	if err != nil {
		fmt.Printf("error opening log file: %v", err)
	}
	defer f.Close()

	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(f)
	log.SetLevel(log.DebugLevel)

	http.HandleFunc("/", handler)
	fmt.Print("application started")
	log.Info("application started")
	log.Fatal(http.ListenAndServe(":80", nil))
}

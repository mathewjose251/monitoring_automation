package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Greet ... Creating a json data type struct
type Greet struct {
	Phrase string `json:"Phrase"`
}

// Greeting ... Creating a global variable for use by main()
// of the type Greet (struct)
var Greeting []Greet

//
var (
	visitCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "hello_world_http_counter",
		Help: "The total number of requests to the api",
	})
)

// main function to setup the listener and
// also has a small faux database stucture
// for storing the json objects
func main() {
	Greeting = []Greet{
		Greet{Phrase: "Hello World"},
	}

	startListening()
}

// Returns the json object to the apiRequest calls
func apiHandler(w http.ResponseWriter, r *http.Request) {
	// Count the visits to the api with prometheus
	visitCounter.Inc()
	// Json reply
	json.NewEncoder(w).Encode(Greeting)
}

func startListening() {
	// Creates server mux to route to api entry point
	httpListen := http.NewServeMux()
	httpListen.HandleFunc("/", apiHandler)

	// Creates server mux to route to prometheus endpoint
	metricListen := http.NewServeMux()
	metricListen.Handle("/metrics", promhttp.Handler())

	go func() {
		// takes incoming requests over port 80 and sends
		// them to the apiHandler to return the json object
		log.Print("Beginning to serve on port :80 / ")
		log.Fatal(http.ListenAndServe("localhost:80", httpListen))
	}()
	// Prometheus metrics endpoint
	log.Print("Beginning to serve on port :9100 /metrics")
	log.Fatal(http.ListenAndServe("localhost:9100", metricListen))
}

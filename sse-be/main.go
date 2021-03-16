package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

// Example SSE server in Golang.
//     $ go run sse.go

type Broker struct {

	// Events are pushed to this channel by the main events-gathering routine
	Notifier chan []byte

	// New client connections
	newClients chan chan []byte

	// Closed client connections
	closingClients chan chan []byte

	// Client connections registry
	clients map[chan []byte]bool
}

func NewServer() (broker *Broker) {
	// Instantiate a broker
	broker = &Broker{
		Notifier:       make(chan []byte, 1),
		newClients:     make(chan chan []byte),
		closingClients: make(chan chan []byte),
		clients:        make(map[chan []byte]bool),
	}

	// Set it running - listening and broadcasting events
	go broker.listen()

	return
}

func (broker *Broker) SentEventHandler(rw http.ResponseWriter, req *http.Request) {

	// Make sure that the writer supports flushing.
	//
	flusher, ok := rw.(http.Flusher)

	if !ok {
		http.Error(rw, "Streaming unsupported!", http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "text/event-stream")
	rw.Header().Set("Cache-Control", "no-cache")
	rw.Header().Set("Connection", "keep-alive")
	rw.Header().Set("Access-Control-Allow-Origin", "*")

	// Each connection registers its own message channel with the Broker's connections registry
	messageChan := make(chan []byte)

	// Signal the broker that we have a new connection
	broker.newClients <- messageChan

	// Remove this client from the map of connected clients
	// when this handler exits.
	defer func() {
		broker.closingClients <- messageChan
	}()

	// Listen to connection close and un-register messageChan
	// notify := rw.(http.CloseNotifier).CloseNotify()
	notify := req.Context().Done()

	go func() {
		<-notify
		broker.closingClients <- messageChan
	}()

	for {

		// Write to the ResponseWriter
		// Server Sent Events compatible
		// fmt.Printf("%s", <-messageChan)
		fmt.Fprintf(rw, "%s", <-messageChan)

		// Flush the data immediatly instead of buffering it for later.
		flusher.Flush()
	}

}

func (broker *Broker) listen() {
	for {
		select {
		case s := <-broker.newClients:

			// A new client has connected.
			// Register their message channel
			broker.clients[s] = true
			log.Printf("Client added. %d registered clients", len(broker.clients))
		case s := <-broker.closingClients:

			// A client has dettached and we want to
			// stop sending them messages.
			delete(broker.clients, s)
			log.Printf("Removed client. %d registered clients", len(broker.clients))
		case event := <-broker.Notifier:

			// We got a new event from the outside!
			// Send event to all connected clients
			for clientMessageChan, _ := range broker.clients {
				clientMessageChan <- event
			}
		}
	}

}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprint(w, "Hello, World!")
}

func main() {

	broker := NewServer()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	go func() {
		for {
			time.Sleep(time.Second * 2)
			resp := ListenEventResponse()
			eventString, err := resp.OutputEventMessage()
			if err != nil {
				log.Println("[ERROR] Receiving event")
			}
			// log.Println("Receiving event: ", eventString)
			broker.Notifier <- []byte(eventString)
		}
	}()

	http.HandleFunc("/sse", broker.SentEventHandler)
	http.HandleFunc("/", indexHandler)

	log.Fatal("HTTP server error: ", http.ListenAndServe("0.0.0.0:"+port, nil))

}

type Response struct {
	Message     string `json:"message"`
	MagicNumber int    `json:"magic_number"`
	Timestamp   string `json:"timestamp"`
}

func ListenEventResponse() *Response {

	dicts := []string{"Hello World", "Lorem Ipsum", "Jakarta", "Surabaya", "Tidal", "Nova", "Alfa", "Mike", "Foxtrot", "Audios"}
	dictRand := dicts[rand.Intn(len(dicts))]

	return &Response{
		Message:     dictRand,
		MagicNumber: rand.Intn(1000),
		Timestamp:   time.Now().String(),
	}
}

func (r *Response) OutputEventMessage() (string, error) {

	b, err := json.Marshal(r)
	if err != nil {
		return "", err
	}

	output := fmt.Sprintf("event: message\ndata: %s\n\n", string(b))

	return output, nil
}

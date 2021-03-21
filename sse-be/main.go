package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/suryatresna/sse-sample/sse-be/internal/broker"
	"github.com/suryatresna/sse-sample/sse-be/internal/event"
	"github.com/suryatresna/sse-sample/sse-be/internal/event/like"
	"github.com/suryatresna/sse-sample/sse-be/internal/event/post"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprint(w, "Hello, World!")
}

func main() {

	listMods := map[string]event.EventMessageInterface{
		"post": post.NewEvent(time.Second * 5),
		"like": like.NewEvent(time.Second * 2),
	}

	brokerSrv := broker.NewServer()
	eventSrv := event.EventListener(brokerSrv, listMods)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	go eventSrv.Listen()

	http.HandleFunc("/sse", brokerSrv.SentEventHandler)
	http.HandleFunc("/", indexHandler)

	log.Fatal("HTTP server error: ", http.ListenAndServe("0.0.0.0:"+port, nil))

}

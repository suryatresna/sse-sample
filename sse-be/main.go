package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/suryatresna/sse-sample/sse-be/internal/broker"
	"github.com/suryatresna/sse-sample/sse-be/internal/event"
	"github.com/suryatresna/sse-sample/sse-be/internal/event/channel"
	"github.com/suryatresna/sse-sample/sse-be/internal/event/views"
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
		// "post":    post.NewEvent(time.Second * 5),
		// "like":    like.NewEvent(time.Second * 2),
		// "comment": comment.NewEvent(time.Second * 2),
		"channel": channel.NewEvent(time.Second * 5),
		"views":   views.NewEvent(time.Second * 3),
	}

	// sequenceMods := []string{"post", "like", "like", "like", "comment"}
	sequenceMods := []string{"channel", "views", "views", "views"}

	brokerSrv := broker.NewServer()
	eventSrv := event.EventListener(brokerSrv, listMods)

	eventSrv.SetModuleSequences(sequenceMods)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	go eventSrv.Listen()

	http.HandleFunc("/feeds-sse", brokerSrv.SentEventHandler)
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/health", indexHandler)

	log.Fatal("HTTP server error: ", http.ListenAndServe("0.0.0.0:"+port, nil))

}

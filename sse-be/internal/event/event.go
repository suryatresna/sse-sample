package event

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
)

type brokerInterface interface {
	Broadcast(messageByte []byte)
}

type EventMessageInterface interface {
	GetResponse(postID int) (interface{}, error)
	GetNameEvent() string
}

type Event struct {
	broker brokerInterface
	events map[string]EventMessageInterface
}

func EventListener(broker brokerInterface, events map[string]EventMessageInterface) *Event {
	return &Event{
		broker: broker,
		events: events,
	}
}

func (e *Event) Listen() {
	var (
		i         int = 1
		idEvent   int = 0
		postIDRnd int = 0
		// module sequence
		modulesSeq []string = []string{"post", "like", "like", "like"}
	)
	for {
		if mod, ok := e.events[modulesSeq[i]]; ok {
			// do iteration
			i++
			// if iteration already limit, reset to zero
			if i == len(modulesSeq) {
				i = 0
			}

			// if mod is posts, generate random ID
			if modulesSeq[i] == "post" {
				postIDRnd = rand.Intn(1090)
			}
			// get response from module event
			eventResp, err := mod.GetResponse(postIDRnd)
			if err != nil {
				log.Println("get response module event error ", err)
				continue
			}
			// convert response event to format event with name of event
			eventMsg, err := e.eventJSONMessage(idEvent, mod.GetNameEvent(), eventResp)
			if err != nil {
				log.Println("get eventJSONMessage event error ", err)
				continue
			}
			idEvent++
			// broadcast event to clients
			e.broker.Broadcast([]byte(eventMsg))
			// log.Println(eventMsg)
		}
	}
}

func (e *Event) eventJSONMessage(idEvent int, name string, rsp interface{}) (string, error) {

	b, err := json.Marshal(rsp)
	if err != nil {
		return "", err
	}

	output := fmt.Sprintf("id: %d\nevent: %s\ndata: %s\n\n", idEvent, name, string(b))
	return output, nil
}

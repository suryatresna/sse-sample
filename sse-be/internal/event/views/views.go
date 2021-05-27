package views

import (
	"math/rand"
	"strconv"
	"time"
)

var nameEvent = "viewsupdate"

type Event struct {
	duration time.Duration
}

type response struct {
	ID             int    `json:"id"`
	ChannelID      int    `json:"channel_id"`
	TotalView      int    `json:"total_view"`
	TotalViewHuman string `json:"total_view_human"`
	Timestamp      string `json:"timestamp"`
}

func NewEvent(duration time.Duration) *Event {
	return &Event{duration}
}

func (e *Event) GetResponse(channelID int) (interface{}, error) {
	return e.getDummyData(channelID), nil
}

func (e *Event) GetNameEvent() string {
	return nameEvent
}

func (e *Event) getDummyData(channelID int) response {
	time.Sleep(e.duration)

	viewNumRnd := rand.Intn(1090)
	return response{
		ID:             rand.Intn(100000),
		ChannelID:      channelID,
		TotalView:      viewNumRnd,
		TotalViewHuman: e.convertLikeToHuman(viewNumRnd),
		Timestamp:      time.Now().String(),
	}

}

func (e *Event) convertLikeToHuman(num int) string {
	if num > 1000 {
		return "1k"
	} else {
		return strconv.Itoa(num)
	}

}

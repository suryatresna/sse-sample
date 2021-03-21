package post

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/brianvoe/gofakeit/v6"
)

var nameEvent = "postupdate"

type Event struct {
	duration time.Duration
}

type response struct {
	ID             int    `json:"id"`
	Title          string `json:"title"`
	Description    string `json:"description"`
	LikeCountNum   int    `json:"like_count_num"`
	LikeCountHuman string `json:"like_count_human"`
	Timestamp      string `json:"timestamp"`
}

func NewEvent(duration time.Duration) *Event {
	return &Event{duration}
}

func (e *Event) GetResponse(postID int) (interface{}, error) {
	return e.getDummyData(postID), nil
}

func (e *Event) GetNameEvent() string {
	return nameEvent
}

func (e *Event) getDummyData(postID int) response {
	time.Sleep(e.duration)

	likeNumRnd := rand.Intn(1090)

	return response{
		ID:             postID,
		Title:          gofakeit.Sentence(4),
		Description:    gofakeit.Sentence(40),
		LikeCountNum:   likeNumRnd,
		LikeCountHuman: e.convertLikeToHuman(likeNumRnd),
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

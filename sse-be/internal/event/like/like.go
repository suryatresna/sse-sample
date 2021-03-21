package like

import (
	"math/rand"
	"strconv"
	"time"
)

var nameEvent = "likeupdate"

type Event struct {
	duration time.Duration
}

type response struct {
	ID             int    `json:"id"`
	PostID         int    `json:"post_id"`
	TotalLike      int    `json:"total_like"`
	TotalLikeHuman string `json:"total_like_human"`
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
		ID:             rand.Intn(100000),
		PostID:         postID,
		TotalLike:      likeNumRnd,
		TotalLikeHuman: e.convertLikeToHuman(likeNumRnd),
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

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
	ID        int             `json:"id"`
	Timestamp string          `json:"timestamp"`
	Type      string          `json:"type"`
	Activity  string          `json:"activity"`
	Content   responseKolpost `json:"content"`
}

type responseKolpost struct {
	Header responseKolpostHeader `json:"header"`
	Body   responseKolpostBody   `json:"body"`
	Footer responseKolpostFooter `json:"footer"`
}
type responseKolpostHeader struct {
	Avatar      string `json:"avatar"`
	AvatarTitle string `json:"avatarTitle"`
}
type responseKolpostMedia struct {
	Type      string `json:"type"`
	Thumbnail string `json:"thumbnail"`
}
type responseKolpostBody struct {
	Media   []responseKolpostMedia `json:"media"`
	Caption responseKolpostCaption `json:"caption"`
}
type responseKolpostCaption struct {
	Text string `json:"text"`
}
type reponseKolpostFooterVal struct {
	Fmt   string `json:"fmt"`
	Value int64  `json:"value"`
}
type responseKolpostFooter struct {
	Like    reponseKolpostFooterVal `json:"like"`
	Comment reponseKolpostFooterVal `json:"comment"`
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

func (e *Event) getDummyImageAvatar() string {
	images := []string{
		"https://images.tokopedia.net/img/cache/215-square/shops-1/2018/7/2/3537902/3537902_a0930577-a86b-435b-9f01-1aa9cece6868.jpg",
		"https://images.tokopedia.net/img/cache/215-square/shops-1/2021/1/19/10655605/10655605_7c915ea0-32da-428f-8fdd-869d5d72ab61.jpg",
		"https://images.tokopedia.net/img/cache/215-square/shops-1/2020/8/28/8968559/8968559_42ec088a-eb9b-4fda-af0e-52cd55ed8189.jpg",
		"https://images.tokopedia.net/img/cache/215-square/shops-1/2017/11/26/2741686/2741686_4da82c26-237d-459e-8a3a-8ac3e6add60d.png",
	}

	return gofakeit.RandomString(images)
}

func (e *Event) getDummyImagePhoto() string {
	images := []string{
		"https://ecs7.tokopedia.net/img/cache/700/attachment/2021/3/23/inv/inv_36a2494c-b922-4dac-8077-53ef07b9eb8a.jpg",
		"https://ecs7.tokopedia.net/img/cache/700/attachment/2021/3/23/inv/inv_8087eb79-037c-49b4-a4ce-d1811139986c.jpg",
		"https://ecs7.tokopedia.net/img/cache/700/attachment/2021/3/23/inv/inv_bd8452c9-e166-43db-9a51-8938282e6787.jpg",
		"https://ecs7.tokopedia.net/img/cache/700/attachment/2021/3/23/inv/inv_2cc2e4d9-9e16-4f8e-a0ee-e5563c237022.jpg",
		"https://ecs7.tokopedia.net/img/cache/700/attachment/2020/2/10/158130705318668/158130705318668_43f9d2ed-78ea-43c2-8465-aa2f46fdc542.png",
		"https://ecs7.tokopedia.net/img/cache/700/attachment/2021/3/22/68657055/68657055_1ae388a3-e718-4047-b5bf-f0b40a8c46c9.png",
		"https://ecs7.tokopedia.net/img/cache/700/attachment/2021/3/23/inv/inv_c1478dff-8561-4364-bd7b-95d127eef3ba.jpg",
	}

	return gofakeit.RandomString(images)
}

func (e *Event) getDummyData(postID int) response {
	time.Sleep(e.duration)

	likeNumRnd := rand.Int63n(1090)
	commentNumRnd := rand.Int63n(1090)

	return response{
		ID:        postID,
		Timestamp: time.Now().String(),
		Type:      "cardpost",
		Activity:  "kolpost",
		Content: responseKolpost{
			Header: responseKolpostHeader{
				Avatar:      e.getDummyImageAvatar(),
				AvatarTitle: gofakeit.Name(),
			},
			Body: responseKolpostBody{
				Media: []responseKolpostMedia{
					{
						Type:      "image",
						Thumbnail: e.getDummyImagePhoto(),
					},
				},
				Caption: responseKolpostCaption{
					Text: gofakeit.Sentence(20),
				},
			},
			Footer: responseKolpostFooter{
				Like: reponseKolpostFooterVal{
					Fmt:   strconv.FormatInt(likeNumRnd, 10),
					Value: likeNumRnd,
				},
				Comment: reponseKolpostFooterVal{
					Fmt:   strconv.FormatInt(commentNumRnd, 10),
					Value: commentNumRnd,
				},
			},
		},
	}

}

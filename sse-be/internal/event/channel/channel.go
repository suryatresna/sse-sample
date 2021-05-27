package channel

import (
	"strconv"
	"time"

	"github.com/brianvoe/gofakeit/v6"
)

var nameEvent = "channelupdate"

type Event struct {
	duration time.Duration
}

type response struct {
	ID        int    `json:"id"`
	Timestamp string `json:"timestamp"`
	Type      string `json:"type"`
	Items     []item `json:"items"`
}

type item struct {
	ID            int64       `json:"id"`
	Title         string      `json:"title"`
	StartTime     string      `json:"start_time"`
	IsLive        bool        `json:"is_live"`
	AirTime       string      `json:"air_time"`
	CoverURL      string      `json:"cover_url"`
	Partner       itemPartner `json:"partner"`
	Video         itemVideo   `json:"video"`
	Stats         itemStats   `json:"stats"`
	Configuration struct {
		HasPromo bool `json:"has_promo"`
		Reminder struct {
			IsSet bool `json:"is_set"`
		} `json:"reminder"`
	} `json:"configuration"`
	Applink string `json:"applink"`
	Weblink string `json:"weblink"`
}

type itemPartner struct {
	ID   int64  `json:"id"`
	Type string `json:"type"`
	Name string `json:"name"`
}

type itemVideo struct {
	ID           int64  `json:"id"`
	Type         string `json:"type"`
	StreamSource string `json:"stream_source"`
	Autoplay     bool   `json:"autoplay"`
}

type itemStats struct {
	View struct {
		Formatted string `json:"formatted"`
	} `json:"view"`
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

func (e *Event) getDummyVideoURL() string {
	videos := []string{
		"https://vod.tokopedia.com/view/adaptive.m3u8?id=d02794a8eb2e4768a3f397d2bbaa91aa",
		"https://vod.tokopedia.com/view/adaptive.m3u8?id=5be77d6136da4c71aaff81e6f30341f0",
		"https://vod.tokopedia.com/view/adaptive.m3u8?id=a12594e7c60e40dcae12700948afe2ee",
		"https://vod.tokopedia.com/view/adaptive.m3u8?id=481295e5af664f3089cba9d6dafaf479",
		"https://vod.tokopedia.com/view/adaptive.m3u8?id=ba3cb88885bd4cf9a114177915f4f4f7",
		"https://vod.tokopedia.com/view/adaptive.m3u8?id=83e4f4f942844053a949dc64894c06ac",
		"https://vod.tokopedia.com/view/adaptive.m3u8?id=e7931bd91aff4fa9bddeac2ecae1994a",
		"https://vod.tokopedia.com/view/adaptive.m3u8?id=246edbe41b1848158bc07c5933283890",
	}

	return gofakeit.RandomString(videos)
}

func (e *Event) getDummyData(channelID int) response {
	time.Sleep(e.duration)
	rsp := response{
		ID:        gofakeit.Number(1, 10000),
		Timestamp: time.Now().Format(time.RFC3339),
		Type:      "channel",
		Items: []item{
			{
				ID:        int64(channelID),
				Title:     gofakeit.Name(),
				StartTime: time.Now().Format(time.RFC3339),
				IsLive:    false,
				AirTime:   "WATCH_AGAIN",
				CoverURL:  e.getDummyImagePhoto(),
				Applink:   "",
				Weblink:   "",
				Video: itemVideo{
					ID:           int64(gofakeit.Number(1, 10000)),
					Type:         "vod",
					StreamSource: e.getDummyVideoURL(),
					Autoplay:     false,
				},
				Partner: itemPartner{
					ID:   int64(gofakeit.Number(1, 10000)),
					Type: "shop",
					Name: gofakeit.Name(),
				},
				Stats: itemStats{
					View: struct {
						Formatted string "json:\"formatted\""
					}{
						Formatted: strconv.FormatInt(int64(gofakeit.Number(1, 1000)), 10),
					},
				},
			},
		},
	}

	return rsp

}

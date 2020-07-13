package vk

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/LarsFox/crawler26/go/entities"
)

const (
	vkURL = "https://api.vk.com/method/%s?%s"
)

// Extractor выдергивает комментарии из ВК.
type Extractor struct {
	token string
}

// NewExtractor возвращает новый Extractor.
func NewExtractor(token string) *Extractor {
	return &Extractor{token: token}
}

// GetComments возвращает комментарии.
// TODO: post_id param.
// TODO: format name properly if needed.
// TODO: empty comment.
func (v *Extractor) GetComments(count int64) ([]*entities.Comment, error) {
	wallGetVals := url.Values{}
	wallGetVals.Add("domain", "mudakoff")

	wallGetAnswer := &vkBody{}
	if err := v.getVK("wall.get", wallGetVals, wallGetAnswer); err != nil {
		return nil, err
	}

	if len(wallGetAnswer.Response.Items) == 0 {
		return nil, entities.ErrNotFound
	}

	comments := make([]*entities.Comment, 0, count)
	for i := 0; i < int(count)/100; i++ {
		wallGetCommentsVals := url.Values{}
		wallGetCommentsVals.Add("owner_id", strconv.FormatInt(wallGetAnswer.Response.Items[0].OwnerID, 10))
		wallGetCommentsVals.Add("post_id", "36857907")
		wallGetCommentsVals.Add("offset", strconv.FormatInt(int64(i)*100, 10))
		wallGetCommentsVals.Add("count", "100")
		wallGetCommentsVals.Add("thread_items_count", "10")

		wallGetCommentsAnswer := &vkBody{}
		if err := v.getVK("wall.getComments", wallGetCommentsVals, wallGetCommentsAnswer); err != nil {
			return nil, err
		}

		if len(wallGetCommentsAnswer.Response.Items) == 0 {
			return comments, nil
		}

		for _, item := range wallGetCommentsAnswer.Response.Items {
			comment := &entities.Comment{
				Text:    item.Text,
				Author:  strconv.FormatInt(item.FromID, 10),
				Replies: make([]*entities.Comment, 0, 10),
			}

			if item.Thread == nil {
				continue
			}
			for _, jtem := range item.Thread.Items {
				comment.Replies = append(comment.Replies, &entities.Comment{
					Text:   jtem.Text,
					Author: strconv.FormatInt(jtem.FromID, 10),
				})
			}

			comments = append(comments, comment)
		}

		time.Sleep(time.Millisecond * 500)
	}
	return comments, nil
}

func (v *Extractor) getVK(method string, values url.Values, to interface{}) error {
	values.Add("access_token", v.token)
	values.Add("v", "5.91")
	uri := fmt.Sprintf(vkURL, method, values.Encode())
	resp, err := http.DefaultClient.Get(uri)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(to); err != nil {
		return err
	}
	return nil
}

package entities

import "errors"

var (
	// ErrNotFound возвращается, когда ничего не нашлось.
	ErrNotFound = errors.New("no comments found")
)

// Comment ТУДУ
type Comment struct {
	Author  string     `json:"author"`
	Text    string     `json:"text"`
	Replies []*Comment `json:"replies,omitempty"`
}

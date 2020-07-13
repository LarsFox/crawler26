package vk

type vkBody struct {
	Response *vkArray `json:"response"`
}
type vkArray struct {
	Items []*vkItem `json:"items"`
	Count int64     `json:"count"`
}
type vkItem struct {
	ID      int64  `json:"id"`
	OwnerID int64  `json:"owner_id"`
	Text    string `json:"text"`

	FromID int64    `json:"from_id"`
	Thread *vkArray `json:"thread"`
}

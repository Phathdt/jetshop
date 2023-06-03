package response

import "time"

type Message struct {
	Id             int       `json:"id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	MessageId      string    `json:"message_id"`
	ThreadId       string    `json:"thread_id"`
	Content        string    `json:"content"`
	FromType       string    `json:"from_type"`
	SendTime       int64     `json:"send_time"`
	MessageType    string    `json:"message_type"`
	Status         string    `json:"status"`
	AutoReply      bool      `json:"auto_reply"`
	Platform       string    `json:"platform"`
	ChannelCode    string    `json:"channel_code"`
	OrganizationId int       `json:"organization_id"`
}

type ListMessage struct {
	Code int       `json:"code"`
	Data []Message `json:"data"`
}

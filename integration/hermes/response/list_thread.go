package response

import "time"

type Thread struct {
	Id                int        `json:"id"`
	CreatedAt         *time.Time `json:"created_at"`
	UpdatedAt         time.Time  `json:"updated_at"`
	ThreadId          string     `json:"thread_id"`
	Content           string     `json:"content"`
	ContentType       string     `json:"content_type"`
	CustomerId        string     `json:"customer_id"`
	CustomerName      string     `json:"customer_name"`
	CustomerAvatarUrl string     `json:"customer_avatar_url"`
	UnreadCount       int        `json:"unread_count"`
	Platform          string     `json:"platform"`
	LastMessageId     string     `json:"last_message_id"`
	LastMessageTime   int64      `json:"last_message_time"`
	ChannelCode       string     `json:"channel_code"`
	OrganizationId    int        `json:"organization_id"`
	SelfPosition      int        `json:"self_position"`
	ToPosition        int        `json:"to_position"`
	FromType          string     `json:"from_type"`
	LastReadMessageId string     `json:"last_read_message_id"`
}
type ListThread struct {
	Code int      `json:"code"`
	Data []Thread `json:"data"`
}

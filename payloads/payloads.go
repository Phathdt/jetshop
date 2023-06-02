package payloads

type PullDetailThreadParams struct {
	ChannelCode      string `json:"channel_code"`
	PlatformThreadId string `json:"platform_thread_id"`
}

type SyncMessageParams struct {
	ChannelCode      string `json:"channel_code"`
	PlatformThreadId string `json:"platform_thread_id"`
}

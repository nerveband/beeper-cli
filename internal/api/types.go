package api

// Chat represents a Beeper chat/conversation
// Using a simplified structure that works with JSON unmarshaling
type Chat struct {
	ID          string                 `json:"id"`
	Title       string                 `json:"title"`
	Type        string                 `json:"type"`
	Network     string                 `json:"network"`
	UnreadCount int                    `json:"unreadCount"`
	IsMuted     bool                   `json:"isMuted"`
	IsArchived  bool                   `json:"isArchived"`
	IsPinned    bool                   `json:"isPinned"`
	// Store participants as raw JSON since it's a complex nested object
	Participants map[string]interface{} `json:"participants"`
}

// Message represents a Beeper message
type Message struct {
	ID        string `json:"id"`
	ChatID    string `json:"chatID"`
	Sender    string `json:"senderName"`
	Text      string `json:"text"`
	Timestamp string `json:"timestamp"` // ISO 8601 timestamp string
	IsSender  bool   `json:"isSender"`
}

// SendMessageRequest represents a message send request
type SendMessageRequest struct {
	Text string `json:"text"`
}

// SendMessageResponse represents the API response after sending a message
type SendMessageResponse struct {
	ID string `json:"id"`
}

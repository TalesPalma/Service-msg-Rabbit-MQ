package models

import ()

type Message struct {
	Content string `json:"content"`
	Channel string `json:"channel"`
}

func NewMessage(content, channel string) *Message {
	return &Message{
		Content: content,
		Channel: channel,
	}
}

func (m *Message) NotValid() bool {
	return m.Content == "" || m.Channel == ""
}

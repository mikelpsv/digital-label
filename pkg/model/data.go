package model

import "github.com/mikelpsv/digital-label/pkg/repositories/dbo"

type LinkData struct {
	KeyLink   string `json:"key_link"`
	KeyData   string `json:"key_data"`
	Type      int    `json:"type"`
	Payload   string `json:"payload"`
	Action    string `json:"_"`
	CreatedAt string `json:"-"`
}

func (l *LinkData) FromDbo(data *dbo.LinkData) *LinkData {
	l.KeyLink = data.KeyLink
	l.KeyData = data.KeyData
	l.Type = data.Type
	l.Payload = data.Payload
	l.Action = data.Action
	l.CreatedAt = data.CreatedAt
	return l
}

func (l *LinkData) ToDbo() *dbo.LinkData {
	return &dbo.LinkData{
		KeyLink:   l.KeyLink,
		KeyData:   l.KeyData,
		Type:      l.Type,
		Payload:   l.Payload,
		Action:    l.Action,
		CreatedAt: l.CreatedAt,
	}
}

package model

import (
	"github.com/mikelpsv/digital-label/internal/repositories/dbo"
)

type ViewData struct {
	OwnerBox     string `json:"owner_box"`      // принадлежность/владелец короба
	Title        string `json:"title"`          // заголовок этикетки
	OrderNum     string `json:"order_num"`      // номер заказа
	Client       string `json:"client"`         // клиент
	Address      string `json:"address"`        // адрес доставки
	BoxLabel     string `json:"box_label"`      // идентификатор короба
	BoxNumber    int    `json:"box_number"`     // порядковый номер короба
	BoxOneOf     int    `json:"box_one_of"`     // количество коробов в посылке
	CustomField1 string `json:"custom_field_1"` // произвольное поле 1
	CustomField2 string `json:"custom_field_2"` // произвольное поле 2
	CustomField3 string `json:"custom_field_3"` // произвольное поле 3
}

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

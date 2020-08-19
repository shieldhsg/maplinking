package model

import (
	"maplinking/common"
)

type Event struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Creator     int    `json:"creator"`
	Type        string `json:"password"`
	Age         int    `json:"age"`
	CreateDt    string `json:"create_dt"`
}

func (m *Event) GetEventInfo(id int) bool {
	return common.Db().First(&m, id).RecordNotFound()
}

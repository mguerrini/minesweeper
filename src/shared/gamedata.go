package shared

import "time"

type GameData struct {
	Id         string         `json:"id"`
	StartTime  time.Time      `json:"start_time"`
	FinishTime time.Time      `json:"finish_time"`
	Status     GameStatusType `json:"status"`

	Board BoardData `json:"board"`
}



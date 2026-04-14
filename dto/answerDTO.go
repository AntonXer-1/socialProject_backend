package dto

import "time"

type AnswerDTO struct {
	Massage string    `json:"massage"`
	Status  bool      `json:"status"`
	Tm      time.Time `json:"time"`
}

func CreateAnswer(status bool, massage string) *AnswerDTO {
	return &AnswerDTO{
		Status:  status,
		Massage: massage,
		Tm:      time.Now(),
	}
}

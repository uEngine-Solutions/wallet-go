package point

import (
	"time"
)

type PointsDeducted struct{
	EventType string	`json:"eventType" type:"string"`
	TimeStamp string 	`json:"timeStamp" type:"string"`
	Id int `json:"id" type:"int"` 
	Holder string `json:"holder" type:"string"` 
	Amount int `json:"amount" type:"int"` 
	
}

func NewPointsDeducted() *PointsDeducted{
	event := &PointsDeducted{EventType:"PointsDeducted", TimeStamp:time.Now().String()}

	return event
}

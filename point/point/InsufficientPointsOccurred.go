package point

import (
	"time"
)

type InsufficientPointsOccurred struct{
	EventType string	`json:"eventType" type:"string"`
	TimeStamp string 	`json:"timeStamp" type:"string"`
	Id int `json:"id" type:"int"` 
	
}

func NewInsufficientPointsOccurred() *InsufficientPointsOccurred{
	event := &InsufficientPointsOccurred{EventType:"InsufficientPointsOccurred", TimeStamp:time.Now().String()}

	return event
}

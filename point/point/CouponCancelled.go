package point

import (
	"time"
)

type CouponCancelled struct{
	EventType string	`json:"eventType" type:"string"`
	TimeStamp string 	`json:"timeStamp" type:"string"`
	Id int `json:"id" type:"int"` 
	Buyer string `json:"buyer" type:"string"` 
	Price int `json:"price" type:"int"` 
	Name string `json:"name" type:"string"` 
	
}

func NewCouponCancelled() *CouponCancelled{
	event := &CouponCancelled{EventType:"CouponCancelled", TimeStamp:time.Now().String()}

	return event
}

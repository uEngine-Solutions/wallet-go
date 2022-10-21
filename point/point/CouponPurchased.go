package point

import (
	"time"
)

type CouponPurchased struct{
	EventType string	`json:"eventType" type:"string"`
	TimeStamp string 	`json:"timeStamp" type:"string"`
	Id int `json:"id" type:"int"` 
	Buyer string `json:"buyer" type:"string"` 
	Price int `json:"price" type:"int"` 
	Name string `json:"name" type:"string"` 
	
}

func NewCouponPurchased() *CouponPurchased{
	event := &CouponPurchased{EventType:"CouponPurchased", TimeStamp:time.Now().String()}

	return event
}

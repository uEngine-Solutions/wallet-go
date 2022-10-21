package point

import (
	"github.com/mitchellh/mapstructure"
)

func wheneverCouponPurchased_UsePoint(data map[string]interface{}) {

	event := NewCouponPurchased()
	mapstructure.Decode(data, &event)
	point := &Point{}
	// TODO Set value from event ( ex: delivery.OrderId = event.Id )
	// TODO Change according to the event (save, delete, put..)
	//	pointrepository.save(point)

	// Sample Logic //
	point.UsePoint(event) //TODO: cannot use event (variable of type *CouponPurchased) as type CouponPurchased in argument to point.UsePoint
}

func wheneverCouponCancelled_CompensatePoint(data map[string]interface{}) {

	event := NewCouponCancelled()
	mapstructure.Decode(data, &event)
	point := &Point{}
	// TODO Set value from event ( ex: delivery.OrderId = event.Id )
	// TODO Change according to the event (save, delete, put..)
	//pointrepository.save(point)

	// Sample Logic //
	point.CompensatePoint(event)
}

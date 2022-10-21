package point

import (
	"gopkg.in/jeevatkm/go-model.v1"
	
	"gorm.io/gorm"
	"fmt"
	"point/external"
)

type Point struct {
	gorm.Model
	Id int `gorm:"primaryKey" json:"id" type:"int"`
	Holder string `json:"holder"`
	Amount int `json:"amount"`

}

func (self *Point) onPostPersist() (err error){
	pointsDeducted := NewPointsDeducted()
	model.Copy(pointsDeducted, self)

	Publish(pointsDeducted)
	insufficientPointsOccurred := NewInsufficientPointsOccurred()
	model.Copy(insufficientPointsOccurred, self)

	Publish(insufficientPointsOccurred)

	return nil
}
func (self *Point) onPrePersist() (err error){}
func (self *Point) onPreUpdate() (err error){}
func (self *Point) onPostUpdate() (err error){}
func (self *Point) onPreRemove() (err error){}
func (self *Point) onPostRemove() (err error){}

func (self *Point) Use(useCommand UseCommand){
}

func (self *UsePoint) UsePoint(couponPurchased CouponPurchased){
	/** Example 1:  new item
	point := &Point{}
	pointrepository.save(point)

	pointsDeducted := NewPointsDeducted()
	model.Copy(pointsDeducted, self)
	Publish(pointsDeducted)
	insufficientPointsOccurred := NewInsufficientPointsOccurred()
	model.Copy(insufficientPointsOccurred, self)
	Publish(insufficientPointsOccurred)
	*/

	/** Example 2:  finding and process

	pointrepository.findById(couponPurchased.get???()).ifPresent(point->{

		point // do something
		pointrepository.save(point);

		pointsDeducted := NewPointsDeducted()
		model.Copy(pointsDeducted, self)
		Publish(pointsDeducted)
		insufficientPointsOccurred := NewInsufficientPointsOccurred()
		model.Copy(insufficientPointsOccurred, self)
		Publish(insufficientPointsOccurred)

	 });
	*/
}
func (self *CompensatePoint) CompensatePoint(couponCancelled CouponCancelled){
	/** Example 1:  new item
	point := &Point{}
	pointrepository.save(point)

	*/

	/** Example 2:  finding and process

	pointrepository.findById(couponCancelled.get???()).ifPresent(point->{

		point // do something
		pointrepository.save(point);


	 });
	*/
}

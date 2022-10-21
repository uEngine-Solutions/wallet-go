package point

import (
	"log"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type PointDB struct {
	db *gorm.DB
}

var pointrepository *PointDB

func PointDBInit() {
	var err error
	pointrepository = &PointDB{}
	pointrepository.db, err = gorm.Open(sqlite.Open("Point_table.db"), &gorm.Config{})

	if err != nil {
		panic("DB Connection Error")
	}
	pointrepository.db.AutoMigrate(&Point{})

}

func PointRepository() *PointDB { //TODO: 여기인가?
	return pointrepository
}

func (self *PointDB) save(entity interface{}) error {

	tx := self.db.Create(entity)

	if tx.Error != nil {
		log.Print(tx.Error)
		return tx.Error
	}
	return nil
}

func (self *PointDB) GetList() []Point {

	entities := []Point{}
	self.db.Find(&entities)

	return entities
}

func (self *PointDB) GetID(id int) *Point { //TODO: FindById
	entity := &Point{}
	self.db.Where("id = ?", id).First(entity)

	return entity
}

func (self *PointDB) Delete(entity *Point) error {
	err2 := self.db.Delete(&entity).Error
	return err2
}

func (self *PointDB) Update(id int, params map[string]string) error {
	entity := &Point{}
	err1 := self.db.Where("id = ?", id).First(entity).Error
	if err1 != nil {
		return err1
	} else {
		update := &Point{}
		ObjectMapping(update, params)

		err2 := self.db.Model(&entity).Updates(update).Error
		return err2
	}

}

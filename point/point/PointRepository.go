package point

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func (self *Point) Get(c echo.Context) error { //TODO: how come
	repository := PointRepository()
	entities := repository.GetList()
	return c.JSON(http.StatusOK, entities)
}

func (self *Point) GetbyId(c echo.Context) error {
	repository := PointRepository() //TODO: 선언이 어디있는?

	id, _ := strconv.Atoi(c.Param("id"))
	self = repository.GetID(id)

	return c.JSON(http.StatusOK, self)
}

func (self *Point) Persist(c echo.Context) error {
	repository := PointRepository()
	params := make(map[string]string)

	c.Bind(&params)
	ObjectMapping(self, params)
	self.onPrePersist()
	err := repository.save(self)
	self.onPostPersist()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	} else {
		return c.JSON(http.StatusOK, self)
	}
}

func (self *Point) Put(c echo.Context) error {
	repository := PointRepository()
	id, _ := strconv.Atoi(c.Param("id"))
	params := make(map[string]string)

	c.Bind(&params)
	self.onPreUpdate()
	err := repository.Update(id, params)
	self.onPostUpdate()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	} else {
		entity := repository.GetID(id)
		return c.JSON(http.StatusOK, entity)
	}
}

func (self *Point) Remove(c echo.Context) error {
	repository := PointRepository()
	id, _ := strconv.Atoi(c.Param("id"))
	self = repository.GetID(id)
	self.onPreRemove()
	err := repository.Delete(self)
	self.onPostRemove()
	return c.JSON(http.StatusOK, err)
}

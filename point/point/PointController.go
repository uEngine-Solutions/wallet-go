package point

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func Use(c echo.Context) error { //TODO: 해당 Point 를 찾아서 메서드 호출해주기
	repository := PointRepository()
	id, _ := strconv.Atoi(c.Param("id"))
	params := make(map[string]string)

	c.Bind(&params)

	//    err := repository.Update(id, params)

	point := repository.GetID(id)

	command := UseCommand{}
	command.Amount = 100
	point.Use(command) 

	return c.JSON(http.StatusOK, point)

	// if err != nil {
	// 	return c.JSON(http.StatusInternalServerError, err)
	// } else {
	// 	entity := repository.GetID(id)
	// 	return c.JSON(http.StatusOK, entity)
	// }
}

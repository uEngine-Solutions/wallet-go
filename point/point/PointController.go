package point
import (
"net/http"
"strconv"
"github.com/labstack/echo"
)

func Use(c echo.Context) error  {
    repository := PointRepository()
    id, _ := strconv.Atoi(c.Param("id"))
    params := make(map[string] string)

    c.Bind(&params)

    err := repository.Update(id, params)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, err)
    } else {
        entity := repository.GetID(id)
        return c.JSON(http.StatusOK, entity)
    }
}


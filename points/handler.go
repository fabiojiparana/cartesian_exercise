package points

import (
	"net/http"

	"github.com/labstack/echo"
)

func FindDistanceHandler(c echo.Context) error {

	x := c.QueryParam("x")
	y := c.QueryParam("y")

	if x == "" || y == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Please enter the x and y coordinates.")
	}

	return c.JSON(http.StatusNoContent, nil)
}

package points

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func FindDistanceHandler(c echo.Context) error {

	x := c.QueryParam("x")
	y := c.QueryParam("y")

	if x == "" || y == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Please enter the x and y coordinates.")
	}

	x1, err := strconv.ParseFloat(x, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid format coordinates x.")
	}

	y1, err := strconv.ParseFloat(y, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid format coordinates y.")
	}

	points, err := FindDistance(x1, y1)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusAccepted, points)
}

package points

import "github.com/labstack/echo"

func AddRoutes(e *echo.Echo) {

	public := e.Group("/api/points")
	public.GET("", FindDistanceHandler)
}

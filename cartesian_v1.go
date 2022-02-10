package main

import (
	"cartesian_exercise/points"
	"flag"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var (
	port  *string
	debug *bool
)

func init() {
	port = flag.String("port", "9000", "port for the service HTTP")
	debug = flag.Bool("debug", false, "mod of the debug")
}

func main() {
	e := echo.New()

	e.HideBanner = true

	if *debug {
		e.Debug = true
		e.Use(middleware.Logger())
	}

	e.Use(middleware.CORS())
	e.Use(middleware.Recover())

	points.AddRoutes(e)

	e.Logger.Fatal(e.Start(":" + *port))
}

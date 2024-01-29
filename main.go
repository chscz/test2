package main

import (
	"ab180/internal/config"
	"ab180/internal/http"
	"ab180/internal/influxdb"
	"ab180/internal/mariadb"
	"github.com/labstack/echo"
)

func main() {
	// set config
	cfg, err := config.LoadEnvironment()
	if err != nil {
		panic(err)
	}

	// conn mariaDB
	repo, err := mariadb.InitMariadb(cfg.MariaDB)
	if err != nil {
		panic(err)
	}

	// conn influxDB
	wAPI, err := influxdb.NewInfluxDBClient(cfg.InfluxDB)
	if err != nil {
		panic(err)
	}

	// set handler
	h := http.NewShortLinkHandler(
		&mariadb.ShortLinkRepo{
			DB: repo,
		},
		wAPI,
		cfg.InfluxDB.Measurement,
	)

	// Echo instance
	e := echo.New()
	e.Renderer = http.RegistTemplate()
	e = http.SetRoute(e, h)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

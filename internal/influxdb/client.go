package influxdb

import (
	"ab180/internal/config"
	"context"
	"fmt"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
)

func NewInfluxDBClient(cfg config.InfluxDB) (api.WriteAPI, error) {
	client := influxdb2.NewClient(fmt.Sprintf("http://%s:%s", cfg.Host, cfg.Port), cfg.Token)

	// validate client connection health
	_, err := client.Health(context.Background())
	if err != nil {
		return nil, err
	}

	wAPI := client.WriteAPI(cfg.Org, cfg.Bucket)
	return wAPI, err
}

package config

type Config struct {
	MariaDB  MariaDB
	InfluxDB InfluxDB
}

func LoadEnvironment() (Config, error) {
	cfg := Config{
		MariaDB{
			UserName: "root",
			Password: "1111",
			Host:     "localhost",
			Port:     "3306",
			Schema:   "ab180",
		},
		InfluxDB{
			Token:       "Al9dmQGkmnaeQRTO_RFF4_YcnIUlZ8JEW4SIQXkDIggL-w0F2dK_8VHN6jaZb8BuWv6evnH9IeeAS6VQkhRSKw==",
			Host:        "localhost",
			Port:        "8086",
			Org:         "ab180",
			Bucket:      "ab180",
			Measurement: "data",
		},
	}
	return cfg, nil
}

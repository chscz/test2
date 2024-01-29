package config

type InfluxDB struct {
	Token       string
	Host        string
	Port        string
	Org         string
	Bucket      string
	Measurement string
}

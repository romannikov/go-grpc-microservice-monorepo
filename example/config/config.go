package config

import "flag"

type ExampleConfig struct {
	DbHost string
	DbPort uint64
}

func GetConfig() (*ExampleConfig, error) {
	dbHost := flag.String("db_host", "localhost", "db host")
	dbPort := flag.Uint64("db_port", 8000, "db port")
	flag.Parse()

	return &ExampleConfig{
		DbHost: *dbHost,
		DbPort: *dbPort,
	}, nil
}

package testing

import (
	"order-gokomodo/configs"
	"order-gokomodo/pkg/db/mysql_db"
)

// NewConfig returns a new decoded Config struct
func NewConfigTest() (*configs.Config, error) {
	// Create config structure
	config := &configs.Config{
		Mysqldb: mysql_db.DatabaseConfig{
			Username:     "order_user",
			Password:     "mysql_db",
			Port:         "3308",
			DatabaseName: "order-s1",
			Host:         "127.0.0.1",
		},
		SecretKey: "abcdef123",
	}
	return config, nil
}

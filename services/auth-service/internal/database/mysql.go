package database

import (
	"fmt"

	"donation-platform/auth-service/internal/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect(cfg *config.Config) (*gorm.DB, error) {

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.MySQLUsername,
		cfg.MySQLPassword,
		cfg.MySQLHost,
		cfg.MySQLPort,
		cfg.MySQLDatabase,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, nil
}
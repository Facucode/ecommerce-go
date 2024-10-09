package gorm_repo

import (
	"ecommerce-go/internal/core/domain"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetGormConnection(config domain.Environment) (*gorm.DB, error) {
	dsn := fmt.Sprintf("postgresql://%s/%s?user=%s&password=%s&sslmode=%s", config.DBDomain, config.DBName, config.DBUser, config.DBPass, config.DBSsl)
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

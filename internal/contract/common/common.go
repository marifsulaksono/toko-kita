package common

import (
	"context"
	"log"

	"github.com/marifsulaksono/go-echo-boilerplate/internal/config"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/constants"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/model"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Contract struct {
	DB    *gorm.DB
	Redis *redis.Client
}

func NewCommon(ctx context.Context) (*Contract, error) {
	db, err := config.Config.Database.ConnectDatabase(ctx, constants.DB_POSTGRESQL)
	if err != nil {
		return nil, err
	}

	rediscli, err := config.Config.Redis.InitRedisClient()
	if err != nil {
		return nil, err
	}

	return &Contract{
		DB:    db,
		Redis: rediscli,
	}, nil
}

func (c *Contract) AutoMigrate() {
	if err := c.DB.AutoMigrate(
		&model.User{},
		&model.TokenAuth{},
		&model.Role{},
		&model.Supplier{},
		&model.Item{},
		&model.StockBatchItem{},
		&model.Sale{},
		&model.SalesDetail{},
		&model.SalesBatchItemDetail{},
	); err != nil {
		log.Fatalf("Error on migration database: %v", err)
	}
	log.Println("Migration successfully.....")
}

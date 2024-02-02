package database

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gormTest/config"
	"gormTest/model"
)

type Database struct {
	PostgreConnection *gorm.DB
}

func ConnectToPostgreSQL() *Database {
	cfg := config.LoadENV("config/.env")
	cfg.ParseENV()
	connStr := "user=" + cfg.UserDBName + " password=" + cfg.UserDBPassword + " dbname=" + cfg.DBName + " sslmode=disable"
	fmt.Println(connStr)

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Warn().Err(err).Msg(" can`t connect to DB")
		return nil
	}
	log.Info().Msg("successfully connected to DB")

	err = db.AutoMigrate(&model.Image{}, &model.Tag{}, &model.Universe{}, &model.Product{}, &model.TagProduct{}, &model.ImageProduct{})
	//	, &model.ImageProduct{}, &model.TagProduct{})
	if err != nil {
		log.Info().Err(err).Msg("cannot auto-migrate models")
		return nil
	}

	return &Database{PostgreConnection: db}
}

package service

import (
	"gormTest/model"
)

type ServiceInterface interface {
	InputValidation(image model.Image) error
	//CreateImage(filename, path string) (uint, error, int)
	//SortImages(options SortOptions) (*[]model.Image, error, int)
	//GetImageByID(userID int64) (*model.Image, error, int)
	//EditImage(ID int64, filename string) (uint, error, int)

}

//type Service struct {
//	DbUser   database.Database
//	validate *validator.Validate
//	//RedisClient redisDB.ClientRedisInterface
//}
//
//func NewService(DbUser database.Database, validate *validator.Validate) *Service {
//	return &Service{DbUser: DbUser, validate: validate}
//}

//func (us *Service) InputValidation(image model.Image) error {
//
//	if validationErr := us.validate.Struct(&image); validationErr != nil {
//		return validationErr
//	}
//	return nil
//}

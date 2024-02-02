package universe

import (
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
	"gormTest/database/universe"
	"gormTest/model"
	"net/http"
	"time"
)

type UniverseServiceInterface interface {
	CreateUniverse(title, ID string) (string, error, int)
	SortUniverse(options model.SortOptions) (*[]model.Universe, error, int)
	GetUniverseByID(userID string) (*model.Universe, error, int)
	EditUniverse(ID, newTagName string) (string, error, int)
}

type UniverseService struct {
	//	GeneralServiceInterface ServiceInterface
	UniverseDB universe.UniverseDatabase
	//RedisClient redisDB.ClientRedisInterface
}

//func NewImageService(serviceInterface ServiceInterface, service Service) *ImageService {
//	return &ImageService{GeneralServiceInterface: serviceInterface, GeneralService: service}
//}

func NewUniverseService(universeDB *universe.UniverseDatabase) *UniverseService {
	return &UniverseService{UniverseDB: *universeDB}
}

func (us *UniverseService) CreateUniverse(title, ID string) (string, error, int) {
	createdUniverse := model.Universe{
		Title:     title,
		PublicID:  ID,
		CreatedAt: time.Now(),
	}

	//if err := is.GeneralService.InputValidation(image); err != nil {
	//	return 0, err, http.StatusBadRequest
	//}

	publicID, err := us.UniverseDB.InsertUniverse(createdUniverse)
	if err != nil {
		if err == gorm.ErrDuplicatedKey {
			return "", err, http.StatusBadRequest
		} else {
			return "", err, http.StatusServiceUnavailable
		}
	}

	return publicID, nil, http.StatusCreated
}

func (us *UniverseService) SortUniverse(options model.SortOptions) (*[]model.Universe, error, int) {

	selectedUniverse, err := us.UniverseDB.SelectSortUniverse(options)
	if err == gorm.ErrInvalidData {
		return nil, err, http.StatusBadRequest
	} else if err == gorm.ErrRecordNotFound {
		return nil, err, http.StatusNotFound
	}

	return selectedUniverse, nil, http.StatusOK
}

func (us *UniverseService) GetUniverseByID(publicID string) (*model.Universe, error, int) {

	//ID, err := strconv.ParseUint(publicID, 10, 64)
	//if err != nil {
	//	fmt.Println("Error:", err)
	//	return nil, err, http.StatusBadRequest
	//}

	selectedUniverse, err := us.UniverseDB.GetUniverseByPublicID(publicID)
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return nil, err, http.StatusNotFound
		default:
			log.Error().Err(err).Msg("Failed to get tag")
			return nil, err, http.StatusServiceUnavailable
		}
	}

	return selectedUniverse, nil, http.StatusOK
}

func (us *UniverseService) EditUniverse(ID, newTitle string) (string, error, int) {

	newTag, err, httpStatus := us.GetUniverseByID(ID)
	if err != nil {
		return "", err, httpStatus
	}

	//imageID, err := strconv.ParseUint(ID, 10, 64)
	//if err != nil {
	//	fmt.Println("Error:", err)
	//	return "", err, http.StatusBadRequest
	//}

	newTag.Title = newTitle
	newTag.UpdatedAt = time.Now()

	//if err := is.GeneralService.InputValidation(*newImage); err != nil {
	//	return 0, err, http.StatusBadRequest
	//}

	err = us.UniverseDB.UpdateUniverse(*newTag)
	if err != nil {
		if gorm.ErrRecordNotFound == err {
			return "", err, http.StatusNotFound
		} else if err == gorm.ErrDuplicatedKey {
			return "", err, http.StatusBadRequest
		} else {
			log.Error().Err(err).Msg("Failed to update tag")
			return "", err, http.StatusServiceUnavailable
		}
	}

	return newTag.PublicID, nil, http.StatusOK
}

//func (is *ImageService) InputValidation(image model.Image) error {
//
//	if validationErr := is.validate.Struct(&image); validationErr != nil {
//		return validationErr
//	}
//	return nil
//}

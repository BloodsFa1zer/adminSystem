package tag

import (
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
	"gormTest/database/tag"
	"gormTest/model"
	"net/http"
	"time"
)

type TagServiceInterface interface {
	CreateTag(title, ID string) (string, error, int)
	SortTags(options model.SortOptions) (*[]model.Tag, error, int)
	GetTagByID(userID string) (*model.Tag, error, int)
	EditTag(ID, newTagName string) (string, error, int)
}

type TagService struct {
	//	GeneralServiceInterface ServiceInterface
	TagDB tag.TagDatabase
	//RedisClient redisDB.ClientRedisInterface
}

//func NewImageService(serviceInterface ServiceInterface, service Service) *ImageService {
//	return &ImageService{GeneralServiceInterface: serviceInterface, GeneralService: service}
//}

func NewTagService(tagDB *tag.TagDatabase) *TagService {
	return &TagService{TagDB: *tagDB}
}

func (ts *TagService) CreateTag(title, ID string) (string, error, int) {
	tag := model.Tag{
		Title:     title,
		PublicID:  ID,
		CreatedAt: time.Now(),
	}

	//if err := is.GeneralService.InputValidation(image); err != nil {
	//	return 0, err, http.StatusBadRequest
	//}

	publicID, err := ts.TagDB.InsertTag(tag)
	if err != nil {
		if err == gorm.ErrDuplicatedKey {
			return "", err, http.StatusBadRequest
		} else {
			return "", err, http.StatusServiceUnavailable
		}
	}

	return publicID, nil, http.StatusCreated
}

func (ts *TagService) SortTags(options model.SortOptions) (*[]model.Tag, error, int) {

	tags, err := ts.TagDB.SelectSortTags(options)
	if err == gorm.ErrInvalidData {
		return nil, err, http.StatusBadRequest
	} else if err == gorm.ErrRecordNotFound {
		return nil, err, http.StatusNotFound
	}

	return tags, nil, http.StatusOK
}

func (ts *TagService) GetTagByID(publicID string) (*model.Tag, error, int) {

	//ID, err := strconv.ParseUint(publicID, 10, 64)
	//if err != nil {
	//	fmt.Println("Error:", err)
	//	return nil, err, http.StatusBadRequest
	//}

	tag, err := ts.TagDB.GetTagByPublicID(publicID)
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return nil, err, http.StatusNotFound
		default:
			log.Error().Err(err).Msg("Failed to get tag")
			return nil, err, http.StatusServiceUnavailable
		}
	}

	return tag, nil, http.StatusOK
}

func (ts *TagService) EditTag(ID, newTitle string) (string, error, int) {

	newTag, err, httpStatus := ts.GetTagByID(ID)
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

	err = ts.TagDB.UpdateTag(*newTag)
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

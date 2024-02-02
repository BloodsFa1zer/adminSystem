package image

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
	"gormTest/config"
	"gormTest/database/image"
	"gormTest/model"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

type ImageServiceInterface interface {
	CreateImage(filename string, img *multipart.FileHeader) (string, error, int)
	SortImages(options model.SortOptions) (*[]model.Image, error, int)
	GetImageByID(userID string) (*model.Image, error, int)
	EditImage(ID, newFilename string) (string, error, int)
}

type ImageService struct {
	//	GeneralServiceInterface service.ServiceInterface
	imageDB image.ImageDbInterface
	//RedisClient redisDB.ClientRedisInterface
}

//func NewImageService(serviceInterface ServiceInterface, service Service) *ImageService {
//	return &ImageService{GeneralServiceInterface: serviceInterface, GeneralService: service}
//}

func NewImageService(imageDB image.ImageDbInterface) *ImageService {
	return &ImageService{imageDB: imageDB}
}

func (is *ImageService) CreateImage(filename string, img *multipart.FileHeader) (string, error, int) {

	//if err := is.GeneralService.InputValidation(image); err != nil {
	//	return 0, err, http.StatusBadRequest
	//}

	// Generate a unique path
	ID := config.GeneratePublicID()
	image := model.Image{
		Filename:  filename,
		Path:      "/server/static/" + ID + ".png",
		CreatedAt: time.Now(),
		PublicID:  ID,
	}

	imagePath := filepath.Join("/Users/mishashevnuk/GolandProjects/gormTest/server/static", ID) + ".png"
	fmt.Println(imagePath)
	if err := saveImage(img, imagePath); err != nil {
		return "", err, http.StatusServiceUnavailable
	}

	err := is.imageDB.InsertImage(image)
	if err != nil {
		if err == gorm.ErrDuplicatedKey {
			return "", err, http.StatusBadRequest
		} else {
			return "", err, http.StatusServiceUnavailable
		}
	}

	return ID, nil, http.StatusCreated
}

// saveImage saves the uploaded image to the specified path
func saveImage(file *multipart.FileHeader, destination string) error {
	src, err := file.Open()
	if err != nil {
		log.Info().Err(err)
		return err
	}

	defer src.Close()

	dst, err := os.Create(destination)
	if err != nil {
		log.Info().Err(err)
		return err
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	if err != nil {
		log.Info().Err(err)
		return err
	}

	return nil
}

func (is *ImageService) SortImages(options model.SortOptions) (*[]model.Image, error, int) {

	images, err := is.imageDB.SelectSortImages(options)
	if err == gorm.ErrInvalidData {
		return nil, err, http.StatusBadRequest
	} else if err == gorm.ErrRecordNotFound {
		return nil, err, http.StatusNotFound
	}

	return images, nil, http.StatusOK
}

func (is *ImageService) GetImageByID(publicID string) (*model.Image, error, int) {
	//fmt.Println(publicID)
	//ID, err := strconv.ParseUint(publicID, 10, 64)
	//if err != nil {
	//	fmt.Println("Error:", err)
	//	return nil, err, http.StatusBadRequest
	//}

	image, err := is.imageDB.GetImageByPublicID(publicID)
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return nil, err, http.StatusNotFound
		default:
			log.Error().Err(err).Msg("Failed to get image")
			return nil, err, http.StatusServiceUnavailable
		}
	}

	return image, nil, http.StatusOK
}

func (is *ImageService) EditImage(ID, newFilename string) (string, error, int) {

	newImage, err, httpStatus := is.GetImageByID(ID)
	if err != nil {
		return "", err, httpStatus
	}

	//imageID, err := strconv.ParseUint(ID, 10, 64)
	//if err != nil {
	//	fmt.Println("Error:", err)
	//	return 0, err, http.StatusBadRequest
	//}

	newImage.Filename = newFilename

	//if err := is.GeneralService.InputValidation(*newImage); err != nil {
	//	return 0, err, http.StatusBadRequest
	//}

	_, err = is.imageDB.UpdateImage(*newImage)
	if err != nil {
		if gorm.ErrRecordNotFound == err {
			return "", err, http.StatusNotFound
		} else if err == gorm.ErrDuplicatedKey {
			return "", err, http.StatusBadRequest
		} else {
			log.Error().Err(err).Msg("Failed to update image")
			return "", err, http.StatusServiceUnavailable
		}
	}

	return ID, nil, http.StatusOK
}

//func (is *ImageService) InputValidation(image model.Image) error {
//
//	if validationErr := is.validate.Struct(&image); validationErr != nil {
//		return validationErr
//	}
//	return nil
//}

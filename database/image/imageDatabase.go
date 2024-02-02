package image

import (
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
	"gormTest/database"
	"gormTest/model"
)

type ImageDatabase struct {
	ConnectionLayer *database.Database
}

func NewImageDatabase(connectionLayer *database.Database) *ImageDatabase {
	return &ImageDatabase{ConnectionLayer: connectionLayer}
}

func (ib *ImageDatabase) InsertImage(img model.Image) error {
	//err := db.Connection.AutoMigrate(&model.Image{})
	//if err != nil {
	//	log.Warn().Err(err).Msg("Failed to auto-migrate database")
	//	return 0, err
	//}

	result := ib.ConnectionLayer.PostgreConnection.Create(&img)
	if result.Error != nil {
		// Check if the error is a unique constraint violation
		pgError, _ := result.Error.(*pgconn.PgError)

		if pgError.Code == "23505" {
			log.Warn().Err(result.Error).Msg("image with that name already exists")
			return gorm.ErrDuplicatedKey
		} else {
			// Handle other errors
			log.Info().Err(result.Error)
			return result.Error
		}
	}

	return nil
}

func (ib *ImageDatabase) SelectSortImages(options model.SortOptions) (*[]model.Image, error) {
	//err := db.Connection.AutoMigrate(&model.Image{})
	//if err != nil {
	//	log.Warn().Err(err).Msg("Failed to auto-migrate database")
	//	return nil, err
	//}
	query := ib.ConnectionLayer.PostgreConnection.Model(&model.Image{})

	switch options.SortBy {
	case "new":
		query = query.Order("created_at DESC")
	case "old":
		query = query.Order("created_at ASC")
		// Handle other sorting options as needed
	case "", "None":
		// Don't apply any sorting by date
	default:
		// Handle invalid sortBy values
		return nil, gorm.ErrInvalidData
	}

	// Apply filtering based on the name criteria
	if options.Name != "None" && options.Name != "" {
		query = query.Where("filename LIKE ?", "%"+options.Name+"%")
	}

	// Apply offset and limit
	query = query.Offset(options.Offset).Limit(20)

	var images []model.Image
	if err := query.Select("filename", "path", "id").Find(&images).Error; err != nil {
		log.Warn().Err(err).Msg("can`t select images")
		return nil, err
	}

	return &images, nil
}

func (ib *ImageDatabase) GetImageByPublicID(publicID string) (*model.Image, error) {
	//err := db.Connection.AutoMigrate(&model.Image{})
	//if err != nil {
	//	log.Warn().Err(err).Msg("Failed to auto-migrate database")
	//	return nil, err
	//}
	var image model.Image
	if err := ib.ConnectionLayer.PostgreConnection.Select("filename", "path", "public_id").Where("public_id = ?", publicID).First(&image).Error; err != nil {
		return nil, err
	}

	return &image, nil
}

func (ib *ImageDatabase) UpdateImage(image model.Image) (uint, error) {

	if err := ib.ConnectionLayer.PostgreConnection.Model(&image).Where("public_id = ?", image.PublicID).Updates(&image).Error; err != nil {

		// Check if the error is a unique constraint violation
		pgError, _ := err.(*pgconn.PgError)

		if pgError.Code == "23505" {
			log.Warn().Err(err).Msg("image with that name already exists")
			return 0, gorm.ErrDuplicatedKey
		} else {
			// Handle other errors
			log.Info().Err(err)

			return 0, err
		}
	}

	return image.ID, nil
}

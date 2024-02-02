package tag

import (
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
	"gormTest/database"
	"gormTest/model"
)

type TagDatabase struct {
	ConnectionLayer *database.Database
}

func NewTagDatabase(connectionLayer *database.Database) *TagDatabase {
	return &TagDatabase{ConnectionLayer: connectionLayer}
}

func (tb *TagDatabase) InsertTag(tag model.Tag) (string, error) {
	//err := db.Connection.AutoMigrate(&model.Image{})
	//if err != nil {
	//	log.Warn().Err(err).Msg("Failed to auto-migrate database")
	//	return 0, err
	//}

	result := tb.ConnectionLayer.PostgreConnection.Create(&tag)
	if result.Error != nil {
		// Check if the error is a unique constraint violation
		pgError, _ := result.Error.(*pgconn.PgError)

		if pgError.Code == "23505" {
			log.Warn().Err(result.Error).Msg("image with that name already exists")
			return "", gorm.ErrDuplicatedKey
		} else {
			// Handle other errors
			log.Info().Err(result.Error)
			return "", result.Error
		}
	}

	return tag.PublicID, nil
}

func (tb *TagDatabase) SelectSortTags(options model.SortOptions) (*[]model.Tag, error) {
	//err := db.Connection.AutoMigrate(&model.Image{})
	//if err != nil {
	//	log.Warn().Err(err).Msg("Failed to auto-migrate database")
	//	return nil, err
	//}
	query := tb.ConnectionLayer.PostgreConnection.Model(&model.Tag{})

	switch options.SortBy {
	case "new":
		query = query.Order("created_at DESC")
	case "old":
		query = query.Order("created_at ASC")
		// Handle other sorting options as needed
	case "", "None":
		// Don't apply any sorting by date
		break
	default:
		// Handle invalid sortBy values
		return nil, gorm.ErrInvalidData
	}

	// Apply offset and limit
	query = query.Offset(options.Offset).Limit(20)

	var tags []model.Tag
	if err := query.Select("title", "public_id").Find(&tags).Error; err != nil {
		log.Warn().Err(err).Msg("can`t select images")
		return nil, err
	}

	return &tags, nil
}

func (tb *TagDatabase) GetTagByPublicID(publicID string) (*model.Tag, error) {
	//err := db.Connection.AutoMigrate(&model.Image{})
	//if err != nil {
	//	log.Warn().Err(err).Msg("Failed to auto-migrate database")
	//	return nil, err
	//}
	var tag model.Tag
	if err := tb.ConnectionLayer.PostgreConnection.Select("title", "public_id").Where("public_id = ?", publicID).First(&tag).Error; err != nil {
		return nil, err
	}

	return &tag, nil
}

func (tb *TagDatabase) UpdateTag(tag model.Tag) error {

	if err := tb.ConnectionLayer.PostgreConnection.Model(&tag).Where("public_id = ?", tag.PublicID).Updates(&tag).Error; err != nil {

		// Check if the error is a unique constraint violation
		pgError, _ := err.(*pgconn.PgError)

		if pgError.Code == "23505" {
			log.Warn().Err(err).Msg("tag with that name already exists")
			return gorm.ErrDuplicatedKey
		} else {
			// Handle other errors
			log.Info().Err(err)

			return err
		}

	}

	return nil
}

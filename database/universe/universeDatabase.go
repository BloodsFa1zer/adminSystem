package universe

import (
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
	"gormTest/database"
	"gormTest/model"
)

type UniverseDatabase struct {
	ConnectionLayer *database.Database
}

func NewUniverseDatabase(connectionLayer *database.Database) *UniverseDatabase {
	return &UniverseDatabase{ConnectionLayer: connectionLayer}
}

func (ub *UniverseDatabase) InsertUniverse(universe model.Universe) (string, error) {

	result := ub.ConnectionLayer.PostgreConnection.Create(&universe)
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

	return universe.PublicID, nil
}

func (ub *UniverseDatabase) SelectSortUniverse(options model.SortOptions) (*[]model.Universe, error) {
	//err := db.Connection.AutoMigrate(&model.Image{})
	//if err != nil {
	//	log.Warn().Err(err).Msg("Failed to auto-migrate database")
	//	return nil, err
	//}
	query := ub.ConnectionLayer.PostgreConnection.Model(&model.Universe{})

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

	var universe []model.Universe
	if err := query.Select("title", "public_id").Find(&universe).Error; err != nil {
		log.Warn().Err(err).Msg("can`t select images")
		return nil, err
	}

	return &universe, nil
}

func (ub *UniverseDatabase) GetUniverseByPublicID(publicID string) (*model.Universe, error) {

	var universe model.Universe
	if err := ub.ConnectionLayer.PostgreConnection.Select("title", "public_id").Where("public_id = ?", publicID).First(&universe).Error; err != nil {
		return nil, err
	}

	return &universe, nil
}

func (ub *UniverseDatabase) UpdateUniverse(universe model.Universe) error {

	if err := ub.ConnectionLayer.PostgreConnection.Model(&universe).Where("public_id = ?", universe.PublicID).Updates(&universe).Error; err != nil {

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

package repositories

import (
	"github.com/labstack/gommon/log"

	"github.com/fxivan/routeOptimizer/cockroach/entities"
	"github.com/fxivan/routeOptimizer/database"
)

type cockroachPostgresRepository struct {
	db database.Database
}

func NewCockroachPostgresRepository(db database.Database) CockroachRepository {
	return &cockroachPostgresRepository{db: db}
}

func (r *cockroachPostgresRepository) InsertCockroachData(in *entities.InsertCockroachDto) error {
	data := &entities.Cockroach{
		Amount: in.Amount,
	}

	result := r.db.GetDb().Create(data)

	if result.Error != nil {
		log.Errorf("InsertCockroachData: %v", result.Error)
		return result.Error
	}

	log.Debugf("InsertCockroachData: %v", result.RowsAffected)
	return nil
}

func (r *cockroachPostgresRepository) GetByIdCockroachData(in *entities.InsertCockroachDto) (entities.Cockroach, error) {
	var dataModel entities.Cockroach
	result := r.db.GetDb().First(&dataModel, in.Id)
	if result.Error != nil {
		log.Errorf("InsertCockroachData: %v", result.Error)
		return dataModel, result.Error
	}

	return dataModel, nil

}

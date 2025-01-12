package usecases

import (
	"github.com/fxivan/routeOptimizer/cockroach/entities"
	"github.com/fxivan/routeOptimizer/cockroach/models"
)

type CockroachUsecase interface {
	CockroachDataProcessing(in *models.AddCockroachData) error
	CockroachSearchTransactions(in *models.IdCockroachData) (entities.Cockroach, error)
}

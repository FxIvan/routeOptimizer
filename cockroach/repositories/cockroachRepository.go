package repositories

import "github.com/fxivan/routeOptimizer/cockroach/entities"

type CockroachRepository interface {
	InsertCockroachData(in *entities.InsertCockroachDto) error
	GetByIdCockroachData(in *entities.InsertCockroachDto) (entities.Cockroach, error)
}

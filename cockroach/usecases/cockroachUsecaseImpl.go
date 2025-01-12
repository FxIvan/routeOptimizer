package usecases

import (
	"fmt"
	"time"

	"github.com/fxivan/routeOptimizer/cockroach/entities"
	"github.com/fxivan/routeOptimizer/cockroach/models"
	"github.com/fxivan/routeOptimizer/cockroach/repositories"
)

type cockroachUsecaseImpl struct {
	cockroachRepository repositories.CockroachRepository
	cockroachMessaging  repositories.CockroachMessaging
}

func NewCockroachUsecaseImpl(
	cockroachRepository repositories.CockroachRepository,
	cockroachMessaging repositories.CockroachMessaging,
) CockroachUsecase {
	return &cockroachUsecaseImpl{
		cockroachRepository: cockroachRepository,
		cockroachMessaging:  cockroachMessaging,
	}
}

func (u *cockroachUsecaseImpl) CockroachDataProcessing(in *models.AddCockroachData) error {
	insertCockroachData := &entities.InsertCockroachDto{
		Amount: in.Amount,
	}

	if err := u.cockroachRepository.InsertCockroachData(insertCockroachData); err != nil {
		return err
	}

	pushCockroachData := &entities.CockroachPushNotificationDto{
		Title:        "Cockroach Detected 🪳 !!!",
		Amount:       in.Amount,
		ReportedTime: time.Now().Local().Format("2006-01-02 15:04:05"),
	}

	if err := u.cockroachMessaging.PushNotification(pushCockroachData); err != nil {
		return err
	}

	return nil
}

func (u *cockroachUsecaseImpl) CockroachSearchTransactions(in *models.IdCockroachData) (entities.Cockroach, error) {
	insertCockroachData := &entities.InsertCockroachDto{
		Id: in.Id,
	}
	fmt.Print("insertCockroachData --->", insertCockroachData)
	var result entities.Cockroach
	var error error
	result, error = u.cockroachRepository.GetByIdCockroachData(insertCockroachData)
	fmt.Print("---------------------------------------------")
	fmt.Print("CockroachSearchTransactions --->", result)
	fmt.Print("---------------------------------------------")
	if error != nil {
		return result, error
	}
	return result, nil
}

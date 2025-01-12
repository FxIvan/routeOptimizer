package repositories

import (
	"github.com/labstack/gommon/log"

	"github.com/fxivan/routeOptimizer/cockroach/entities"
)

type cockroachFCMMessaging struct{}

func NewCockroachFCMMessaging() CockroachMessaging {
	return &cockroachFCMMessaging{}
}

func (c *cockroachFCMMessaging) PushNotification(m *entities.CockroachPushNotificationDto) error {
	// ... handle logic to push FCM notification here ...
	log.Debugf("Pushed FCM notification with data: %v", m)

	return nil
}

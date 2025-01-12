package repositories

import "github.com/fxivan/routeOptimizer/cockroach/entities"

type CockroachMessaging interface {
	PushNotification(m *entities.CockroachPushNotificationDto) error
}

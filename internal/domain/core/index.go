package core

import (
	"github.com/rendau/sms/internal/interfaces"
)

type St struct {
	lg               interfaces.Logger
	cache            interfaces.Cache
	smscUsername     string
	smscPassword     string
	smscSender       string
	balanceNotifUrls map[float64]string
}

func New(lg interfaces.Logger, cache interfaces.Cache, smscUsername string, smscPassword string, smscSender string, array map[float64]string) *St {
	core := &St{
		lg:               lg,
		cache:            cache,
		smscUsername:     smscUsername,
		smscPassword:     smscPassword,
		smscSender:       smscSender,
		balanceNotifUrls: array,
	}

	return core
}

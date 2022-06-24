package core

import (
	"github.com/rendau/dop/adapters/cache"
	"github.com/rendau/dop/adapters/client/httpc"
	"github.com/rendau/dop/adapters/logger"
)

type St struct {
	lg    logger.Lite
	cache cache.Cache
	httpc httpc.HttpC
}

func New(lg logger.Lite, cache cache.Cache, httpc httpc.HttpC) *St {
	return &St{
		lg:    lg,
		cache: cache,
		httpc: httpc,
	}
}

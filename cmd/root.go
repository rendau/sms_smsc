package cmd

import (
	"crypto/tls"
	"net/http"
	"os"
	"time"

	dopCache "github.com/rendau/dop/adapters/cache"
	dopCacheMem "github.com/rendau/dop/adapters/cache/mem"
	dopCacheRedis "github.com/rendau/dop/adapters/cache/redis"
	"github.com/rendau/dop/adapters/client/httpc"
	"github.com/rendau/dop/adapters/client/httpc/httpclient"
	dopLoggerZap "github.com/rendau/dop/adapters/logger/zap"
	dopServerHttps "github.com/rendau/dop/adapters/server/https"
	"github.com/rendau/dop/dopTools"
	"github.com/rendau/sms_smsc/internal/adapters/server/rest"
	"github.com/rendau/sms_smsc/internal/domain/core"
)

func Execute() {
	app := struct {
		lg         *dopLoggerZap.St
		cache      dopCache.Cache
		core       *core.St
		restApiSrv *dopServerHttps.St
	}{}

	confLoad()

	app.lg = dopLoggerZap.New(conf.LogLevel, conf.Debug)

	if conf.RedisUrl == "" {
		app.cache = dopCacheMem.New()
	} else {
		app.cache = dopCacheRedis.New(
			app.lg,
			conf.RedisUrl,
			conf.RedisPsw,
			conf.RedisDb,
			conf.RedisKeyPrefix,
		)
	}

	app.core = core.New(
		app.lg,
		app.cache,
		httpclient.New(app.lg, httpc.OptionsSt{
			Client: &http.Client{
				Timeout: 15 * time.Second,
				Transport: &http.Transport{
					TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
				},
			},
			BaseUrl: "https://smsc.kz/sys",
			BaseParams: map[string][]string{
				"login": {conf.SmscUsername},
				"psw":   {conf.SmscPassword},
			},
		}),
	)

	// START

	app.lg.Infow("Starting")

	app.restApiSrv = dopServerHttps.Start(
		conf.HttpListen,
		rest.GetHandler(
			app.lg,
			app.core,
			conf.HttpCors,
		),
		app.lg,
	)

	var exitCode int

	select {
	case <-dopTools.StopSignal():
	case <-app.restApiSrv.Wait():
		exitCode = 1
	}

	// STOP

	app.lg.Infow("Shutting down...")

	if !app.restApiSrv.Shutdown(20 * time.Second) {
		exitCode = 1
	}

	app.lg.Infow("Exit")

	os.Exit(exitCode)
}

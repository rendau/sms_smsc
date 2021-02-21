package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"

	memCache "github.com/rendau/sms/internal/adapters/cache/mem"
	"github.com/rendau/sms/internal/adapters/cache/redis"
	"github.com/rendau/sms/internal/adapters/httpapi"
	"github.com/rendau/sms/internal/adapters/logger/zap"
	"github.com/rendau/sms/internal/domain/core"
	"github.com/rendau/sms/internal/interfaces"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use: "sms",
	Run: func(cmd *cobra.Command, args []string) {
		var err error

		loadConf()

		app := struct {
			log     *zap.St
			core    *core.St
			restApi *httpapi.St
			cache   interfaces.Cache
		}{}

		app.log, err = zap.New(viper.GetString("log_level"), viper.GetBool("debug"), false)
		if err != nil {
			log.Fatal(err)
		}

		if viper.GetString("redis.url") == "" {
			app.cache = memCache.New()
		} else {
			app.cache = redis.New(
				app.log,
				viper.GetString("redis.url"),
				viper.GetString("redis.psw"),
				viper.GetInt("redis.db"),
			)
		}

		balanceNotifyPars := viper.Get("balance_notify").(map[float64]string)

		app.core = core.New(
			app.log,
			app.cache,
			viper.GetString("smsc_username"),
			viper.GetString("smsc_password"),
			viper.GetString("smsc_sender"),
			balanceNotifyPars,
		)

		app.restApi = httpapi.New(app.log, viper.GetString("http_listen"), app.core)

		app.log.Infow(
			"Starting",
			"http_listen", viper.GetString("http_listen"),
		)

		for b, url := range balanceNotifyPars {
			app.log.Infow(
				"Balance alarm parameter",
				"balance", b,
				"url", url,
			)
		}

		app.restApi.Start()

		stop := make(chan os.Signal, 1)
		signal.Notify(stop, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

		var exitCode int

		select {
		case <-stop:
		case <-app.restApi.Wait():
			exitCode = 1
		}

		app.log.Infow("Shutting down...")

		ctx, ctxCancel := context.WithTimeout(context.Background(), 20*time.Second)
		defer ctxCancel()

		err = app.restApi.Shutdown(ctx)
		if err != nil {
			app.log.Errorw("Fail to shutdown http-api", err)
			exitCode = 1
		}

		os.Exit(exitCode)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func loadConf() {
	viper.SetDefault("debug", "false")
	viper.SetDefault("http_listen", ":80")
	viper.SetDefault("log_level", "debug")

	confFilePath := os.Getenv("CONF_PATH")
	if confFilePath == "" {
		confFilePath = "conf.yml"
	}
	viper.SetConfigFile(confFilePath)
	_ = viper.ReadInConfig()

	viper.AutomaticEnv()

	blnNotify := map[float64]string{}

	const bnEnvPrefix = "BALANCE_NOTIFY_"

	for _, element := range os.Environ() {
		pair := strings.SplitN(element, "=", 2)
		if len(pair) == 2 && strings.HasPrefix(pair[0], bnEnvPrefix) {
			key := strings.TrimPrefix(pair[0], bnEnvPrefix)
			if n, _ := strconv.ParseFloat(key, 64); n > 0 {
				blnNotify[n] = pair[1]
			}
		}
	}

	viper.Set("balance_notify", blnNotify)
}

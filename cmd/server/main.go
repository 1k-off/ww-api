package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"
	"ww-api/api"
	"ww-api/pkg/app"
	"ww-api/pkg/config"
)

func init() {
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)
}

func main() {
	cfg, err := config.Load("./config/")
	if err != nil {
		log.Fatal().Err(err).Msg("config load error")
	}
	appService, err := app.New(
		cfg.Database.ConnectionString,
		cfg.Server.AccessToken.PrivateKey,
		cfg.Server.AccessToken.PublicKey,
		cfg.Server.RefreshToken.PrivateKey,
		cfg.Server.RefreshToken.PublicKey,
		cfg.Server.AccessToken.ExpiresIn,
		cfg.Server.RefreshToken.ExpiresIn,
	)
	if err != nil {
		log.Fatal().Err(err).Msg("app init error")
	}

	manager, err := appService.NewManager(
		cfg.Queue.Memphis.ClientLogin,
		cfg.Queue.Memphis.ClientToken,
		cfg.Queue.Memphis.Url,
		cfg.Queue.Memphis.SslTargetsSN,
		cfg.Queue.Memphis.UptimeTargetsSN,
		cfg.Queue.Memphis.DomainExpirationTargetsSN,
		cfg.Queue.Memphis.SslMetricsSN,
		cfg.Queue.Memphis.UptimeMetricsSN,
		cfg.Queue.Memphis.DomainExpirationMetricsSN,
	)
	if err != nil {
		log.Fatal().Err(err).Msg("manager init error")
	}

	stopCh := make(chan os.Signal)
	signal.Notify(stopCh, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-stopCh
		err = manager.Stop()
		if err != nil {
			log.Error().Err(err).Msg("manager stop error")
			return
		}
		appService.Stop()
		log.Info().Msg("Application stopped")
		os.Exit(0)
	}()

	log.Info().Msg("Application job manager started")
	go manager.Run()

	go func() {
		if cfg.LogLevel == "debug" {
			select {
			case <-stopCh:
				return
			case <-time.After(60 * time.Second):
				log.Debug().Msgf("goroutines: %d", runtime.NumGoroutine())
			}
		}
	}()

	err = api.Start(appService, cfg.Server.Port)
	if err != nil {
		log.Fatal().Err(err).Msg("api start error")
	}
}

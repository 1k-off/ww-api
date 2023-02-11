package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"ww-api/api"
	"ww-api/pkg/app"
	"ww-api/pkg/config"
)

func main() {
	cfg, err := config.Load("./config/")
	if err != nil {
		log.Fatalln(err)
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
		panic(err)
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
		log.Fatalln(err)
	}

	stopCh := make(chan os.Signal)
	signal.Notify(stopCh, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-stopCh
		err = manager.Stop()
		if err != nil {
			log.Println(err)
			return
		}
		appService.Stop()
		os.Exit(255)
	}()

	go manager.Run()
	log.Fatalln(api.Start(appService, cfg.Server.Port))
}

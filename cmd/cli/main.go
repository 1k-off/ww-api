package main

import (
	"github.com/alecthomas/kong"
	"log"
	"ww-api/pkg/app"
	"ww-api/pkg/config"
)

type CLI struct {
	User UserCmd `cmd:"" help:"User commands"`
}

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

	cli := CLI{}
	cliCtx := kong.Parse(&cli,
		kong.Name("webwatch"),
		kong.Description("WebWatch CLI"),
		kong.UsageOnError(),
		kong.ConfigureHelp(kong.HelpOptions{Compact: false}),
	)

	err = cliCtx.Run(&app.Service{
		UserService: appService.UserService,
	})
	cliCtx.FatalIfErrorf(err)
}

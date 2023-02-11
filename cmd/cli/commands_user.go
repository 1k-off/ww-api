package main

import (
	"ww-api/pkg/app"
	"ww-api/pkg/entities"
)

type UserCmd struct {
	Create UserCreateCmd `cmd:"" help:"Create a new user"`
	Update UserUpdateCmd `cmd:"" help:"Update an existing user"`
	Delete UserDeleteCmd `cmd:"" help:"Delete an existing user"`
}

type UserCreateCmd struct {
	Login    string `arg required help:"Login name"`
	Password string `arg required help:"Password"`
}
type UserUpdateCmd struct {
	Login    string `arg required help:"Login name"`
	Password string `arg required help:"Password"`
}
type UserDeleteCmd struct {
	Login string `arg required help:"Login name"`
}

func (cmd *UserCmd) Run() error {
	return nil
}

func (cmd *UserCreateCmd) Run(app *app.Service) error {
	u := &entities.User{
		Login:    cmd.Login,
		Password: cmd.Password,
	}
	_, err := app.UserService.Create(u)
	if err != nil {
		return err
	}
	return nil
}

func (cmd *UserUpdateCmd) Run(app *app.Service) error {
	u := &entities.User{
		Login:    cmd.Login,
		Password: cmd.Password,
	}
	_, err := app.UserService.Update(u)
	if err != nil {
		return err
	}
	return nil
}

func (cmd *UserDeleteCmd) Run(app *app.Service) error {
	u, err := app.UserService.GetByLogin(cmd.Login)
	if err != nil {
		return err
	}
	err = app.UserService.Delete(u.ID.Hex())
	if err != nil {
		return err
	}
	return nil
}

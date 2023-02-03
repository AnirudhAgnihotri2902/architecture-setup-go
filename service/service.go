package service

import (
	"context"
	"github.com/AnirudhAgnihotri2902/architecture-setup-go/ent"
	"github.com/AnirudhAgnihotri2902/architecture-setup-go/repo"
)

type Getqueries interface {
	Getuserbyname(ctx context.Context, client *ent.Client, username string) ([]*ent.Registers, error)
	Createnewuser(ctx context.Context, client *ent.Client, username string, userpassword string) (*ent.Registers, error)
}

func Getuserbyname(ctx context.Context, client *ent.Client, username string) ([]*ent.Registers, error) {
	return repo.QueryUser(ctx, client, username)
}

func Createnewuser(ctx context.Context, client *ent.Client, username string, userpassword string) (*ent.Registers, error) {
	return repo.CreateUser(ctx, client, username, userpassword)
}

func Add(a int, b int) (res int) {
	c := a + b
	return c
}
func Subtract(a int, b int) (res int) {
	c := a - b
	return c
}
func Multiply(a int, b int) (res int) {
	return a * b
}
func Divide(a int, b int) (res int) {
	return a / b
}

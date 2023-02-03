package repo

import (
	"context"
	"fmt"
	"github.com/AnirudhAgnihotri2902/architecture-setup-go/ent"
	"github.com/AnirudhAgnihotri2902/architecture-setup-go/ent/registers"
	"log"
)

// creating user..
func CreateUser(ctx context.Context, client *ent.Client, username string, userpassword string) (*ent.Registers, error) {
	u, err := client.Registers.
		Create().
		SetUsername(username).
		SetPassword(userpassword).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}
	log.Println("user was created: ", u)
	return u, nil
}

// query checking..
func QueryUser(ctx context.Context, client *ent.Client, username string) ([]*ent.Registers, error) {
	u, err := client.Registers.
		Query().
		Where(registers.Username(username)).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}
	log.Println("user returned: ", u)
	return u, nil
}

//// delete user..
//func DelUser(ctx context.Context, client *ent.Client, username string) error {
//	_, err := client.Registers.
//		Delete().
//		Where(registers.Username(username)).
//		Exec(ctx)
//	if err != nil {
//		return fmt.Errorf("failed querying user: %w", err)
//	}
//	log.Println("user deleted ")
//	return nil
//}

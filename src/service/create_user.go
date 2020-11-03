package service

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/lgranade/minesweeperapi/dao"
	"github.com/lgranade/minesweeperapi/dao/minesweeper"
	"github.com/lgranade/minesweeperapi/model"
)

// HardcodedUserID is here to get first version without reading access token
// TODO: take this from access token
var HardcodedUserID uuid.UUID

func init() {
	HardcodedUserID, _ = uuid.Parse("e341410d-752a-404f-9acc-904764fd38f3")
}

// CreateUser creates a user
func CreateUser(ctx context.Context, name string, password string) (*model.User, error) {
	q, tx, err := dao.GetQuerier().WithTx()
	if err != nil {
		log.Println("An error occurred trying to establish db connection")
		return nil, ErrInternal
	}

	defer tx.Rollback()

	dbAccount, err := q.CreateAccount(ctx, minesweeper.CreateAccountParams{
		ID:            HardcodedUserID,
		LoginName:     name,
		LoginPassword: password,
	})
	if err != nil {
		// TODO: check for non unique constraint
		log.Println("Error occurred creating user in local db: ", err)
		return nil, ErrInternal
	}

	user := &model.User{}
	fillUserFromDB(user, &dbAccount)

	if err = tx.Commit(); err != nil {
		log.Println("Error occurred trying to commit tx: ", err)
		return nil, ErrInternal
	}
	return user, nil
}

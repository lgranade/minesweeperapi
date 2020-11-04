package service

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/lgranade/minesweeperapi/dao"
	"github.com/lgranade/minesweeperapi/dao/minesweeper"
	"github.com/lgranade/minesweeperapi/model"
)

// CreateUser creates a user
func CreateUser(ctx context.Context, name string, password string) (*model.User, error) {
	q, tx, err := dao.GetQuerier().WithTx()
	if err != nil {
		log.Println("An error occurred trying to establish db connection")
		return nil, ErrInternal
	}

	defer tx.Rollback()

	userID := uuid.New()

	dbAccount, err := q.CreateAccount(ctx, minesweeper.CreateAccountParams{
		ID:            userID,
		LoginName:     name,
		LoginPassword: password,
	})

	if dao.IsPQIntegrityViolationError(err) {
		log.Println("Can't create user in local db, violating constraint: ", err)
		return nil, ErrDuplicatedUser
	}
	if err != nil {
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

package service

import (
	"context"
	"database/sql"
	"log"

	"github.com/google/uuid"
	"github.com/lgranade/minesweeperapi/dao"
	"github.com/lgranade/minesweeperapi/dao/minesweeper"
	"github.com/lgranade/minesweeperapi/model"
)

// ReadUser reads a user from db
func ReadUser(ctx context.Context, userID uuid.UUID) (*model.User, error) {
	q := dao.GetQuerier().WithoutTx()

	dbAccount, err := q.GetAccountByID(ctx, userID)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Println("Error occurred querying game: ", err)
			return nil, ErrInternal
		}
		return nil, ErrNonexistentUser
	}

	user := &model.User{}
	fillUserFromDB(user, &dbAccount)

	return user, nil
}

func fillUserFromDB(mUser *model.User, dbAccount *minesweeper.Account) {
	mUser.ID = dbAccount.ID
	mUser.Name = dbAccount.LoginName
	mUser.CreatedAt = dbAccount.CreatedAt.Unix()
}

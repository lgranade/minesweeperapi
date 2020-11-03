package dao

import (
	"context"

	"github.com/google/uuid"
	"github.com/lgranade/minesweeperapi/dao/minesweeper"
)

// QuerierMockupBuilder is the transactioned querier mockup builder
type QuerierMockupBuilder struct {
	Querier *QuerierMockup
	Txs     []*TxMockup
}

// QuerierMockup is the querier mockup that allows to set the desired handlers to be used in each test
type QuerierMockup struct {
	CreateAccountHandler      func(ctx context.Context, arg minesweeper.CreateAccountParams) (minesweeper.Account, error)
	CreateGameHandler         func(ctx context.Context, arg minesweeper.CreateGameParams) (minesweeper.Game, error)
	GetGameByIDHandler        func(ctx context.Context, id uuid.UUID) (minesweeper.Game, error)
	GetAndLockGameByIDHandler func(ctx context.Context, id uuid.UUID) (minesweeper.Game, error)
	GetAccountByIDHandler     func(ctx context.Context, id uuid.UUID) (minesweeper.Account, error)
	UpdateGameHandler         func(ctx context.Context, arg minesweeper.UpdateGameParams) (minesweeper.Game, error)
}

// TxMockup is a transaction mockup
type TxMockup struct {
	Rolledback bool
	Committed  bool
}

// Commit simulates a commit in the transaction
func (t *TxMockup) Commit() error {
	if !t.Rolledback {
		t.Committed = true
	}
	return nil
}

// Rollback simulates a rollback in the transaction
func (t *TxMockup) Rollback() error {
	if !t.Committed {
		t.Rolledback = true
	}
	return nil
}

// NewQuerierMockupBuilder is the querier mockup builder
func NewQuerierMockupBuilder() *QuerierMockupBuilder {
	return &QuerierMockupBuilder{
		Querier: &QuerierMockup{},
		Txs:     []*TxMockup{},
	}
}

// WithTx is the transactioned querier mockup builder
func (b *QuerierMockupBuilder) WithTx() (minesweeper.Querier, Tx, error) {
	tx := &TxMockup{}
	b.Txs = append(b.Txs, tx)
	return b.Querier, tx, nil
}

// WithoutTx is the transactionless querier mockup builder
func (b *QuerierMockupBuilder) WithoutTx() minesweeper.Querier {
	return b.Querier
}

// CreateAccount simulates creation of an account
func (q *QuerierMockup) CreateAccount(ctx context.Context, arg minesweeper.CreateAccountParams) (minesweeper.Account, error) {
	if q.CreateAccountHandler != nil {
		return q.CreateAccountHandler(ctx, arg)
	}
	return minesweeper.Account{}, nil
}

// CreateGame simulates creation of a game
func (q *QuerierMockup) CreateGame(ctx context.Context, arg minesweeper.CreateGameParams) (minesweeper.Game, error) {
	if q.CreateGameHandler != nil {
		return q.CreateGameHandler(ctx, arg)
	}
	return minesweeper.Game{}, nil
}

// GetAndLockGameByID simulates reading and locking of a game
func (q *QuerierMockup) GetAndLockGameByID(ctx context.Context, id uuid.UUID) (minesweeper.Game, error) {
	if q.GetAndLockGameByIDHandler != nil {
		return q.GetAndLockGameByIDHandler(ctx, id)
	}
	return minesweeper.Game{}, nil
}

// GetGameByID simulates reading of a game
func (q *QuerierMockup) GetGameByID(ctx context.Context, id uuid.UUID) (minesweeper.Game, error) {
	if q.GetGameByIDHandler != nil {
		return q.GetGameByIDHandler(ctx, id)
	}
	return minesweeper.Game{}, nil
}

// GetAccountByID simulates reading of a user
func (q *QuerierMockup) GetAccountByID(ctx context.Context, id uuid.UUID) (minesweeper.Account, error) {
	if q.GetAccountByIDHandler != nil {
		return q.GetAccountByIDHandler(ctx, id)
	}
	return minesweeper.Account{}, nil
}

// UpdateGame simulates updating of a game
func (q *QuerierMockup) UpdateGame(ctx context.Context, arg minesweeper.UpdateGameParams) (minesweeper.Game, error) {
	if q.UpdateGameHandler != nil {
		return q.UpdateGameHandler(ctx, arg)
	}
	return minesweeper.Game{}, nil
}

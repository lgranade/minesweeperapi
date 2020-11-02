package dao

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/lgranade/minesweeperapi/dao/minesweeper"

	//Postgres Go driver
	_ "github.com/lib/pq"
)

// Tx is a transaction.
type Tx interface {
	Commit() error
	Rollback() error
}

// Querier is implemented in the object used by upper layer
type Querier interface {
	WithTx() (minesweeper.Querier, Tx, error)
	WithoutTx() minesweeper.Querier
}

// stores the db client object to be used. It could be a mock up
// set externally
var querier Querier

// SetQuerier allows an external agent to mock the client
func SetQuerier(q Querier) {
	querier = q
}

// GetQuerier returns the setted querier
func GetQuerier() Querier {
	return querier
}

// QuerierImpl is the internal implementation for db client
type QuerierImpl struct {
	db      *sql.DB
	queries *minesweeper.Queries
}

// WithTx creates a new Querier using a newly created transaction
func (q *QuerierImpl) WithTx() (minesweeper.Querier, Tx, error) {
	tx, err := q.db.Begin()
	if err != nil {
		return nil, nil, err
	}
	return q.queries.WithTx(tx), tx, nil
}

// WithoutTx returns a transactionless querier.
// It doesn't creates it, just returns the same one every time.
func (q *QuerierImpl) WithoutTx() minesweeper.Querier {
	return q.queries
}

// DBConfig represents the db configuration received inside config object
type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
	SSLMode  string
}

// Start starts the database client if it wasn't setted from the outside
// (if it wasn't for example mocked)
func Start(c *DBConfig) error {
	if querier == nil {
		psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
			"password=%s dbname=%s sslmode=%s",
			c.Host, c.Port, c.User, c.Password, c.Name, c.SSLMode)

		db, err := sql.Open("postgres", psqlInfo)
		if err != nil {
			return err
		}

		querier = &QuerierImpl{
			db:      db,
			queries: minesweeper.New(db),
		}

		if err = db.Ping(); err != nil {
			log.Println("Error doing ping: ", err)
		} else {
			log.Println("Ping DB successfully")
		}
	}

	return nil
}

// Stop stops the database client if it was instantiated internally
func Stop() {
	if qi, ok := querier.(*QuerierImpl); ok {
		err := qi.db.Close()
		if err != nil {
			log.Println("Error closing local database: ", err)
		}
	}
}

// ToNullString creates a valid NullString from a string
func ToNullString(s string) sql.NullString {
	return sql.NullString{
		String: s,
		Valid:  true,
	}
}

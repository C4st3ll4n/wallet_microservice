package database

import (
	"database/sql"
	"github.com/C4st3ll4n/wallet/internal/entity"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
	"testing"
)

type AccountDBTestSuite struct {
	suite.Suite
	db        *sql.DB
	accountDB *AccountDB
}

func (s *AccountDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Nil(err)

	s.db = db

	db.Exec("CREATE TABLE clients (id varchar(255), name varchar(255), email varchar(255), created_at date)")
	db.Exec("CREATE TABLE accounts (id varchar(255), client_id varchar(255), balance float, created_at date)")

	s.accountDB = NewAccountDB(db)
}

func (s *AccountDBTestSuite) TearDownSuite() {
	defer s.db.Close()
	s.db.Exec("DROP TABLE clients")
	s.db.Exec("DROP TABLE accounts")
}

func TestAccountDBTestSuite(t *testing.T) {
	suite.Run(t, new(AccountDBTestSuite))
}

func (s *AccountDBTestSuite) TestGetAccount() {

	client, _ := entity.NewClient("Fulano de Tal", "fulano@mail.com")
	account := entity.NewAccount(client)

	err := s.accountDB.Save(account)
	s.Nil(err)

	accountDB, err := s.accountDB.FindById(account.ID)
	s.Nil(err)
	s.Equal(account.Balance, accountDB.Balance)
	s.Equal(account.ID, accountDB.ID)
	s.Equal(account.Client.ID, accountDB.Client.ID)
}

func (s *AccountDBTestSuite) TestSaveAccount() {

	client, _ := entity.NewClient("Ciclano de Tal", "ciclano@mail.com")
	account := entity.NewAccount(client)

	err := s.accountDB.Save(account)
	s.Nil(err)
}

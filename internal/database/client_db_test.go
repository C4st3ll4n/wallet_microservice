package database

import (
	"database/sql"
	"github.com/C4st3ll4n/wallet/internal/entity"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ClientDBTestSuite struct {
	suite.Suite
	db       *sql.DB
	clientDB *ClientDB
}

func (s *ClientDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Nil(err)

	s.db = db

	db.Exec("CREATE TABLE clients (id varchar(255), name varchar(255), email varchar(255), created_at date)")

	s.clientDB = NewClientDB(db)
}

func (s *ClientDBTestSuite) TearDownSuite() {
	defer s.db.Close()
	s.db.Exec("DROP TABLE clients")
}

func TestClientDBTestSuite(t *testing.T) {
	suite.Run(t, new(ClientDBTestSuite))
}

func (s *ClientDBTestSuite) TestGetClient() {

	client, _ := entity.NewClient("Fulano de Tal", "fulano@mail.com")
	s.clientDB.Save(*client)

	clientDB, err := s.clientDB.Get(client.ID)
	s.Nil(err)
	s.Equal(client.Name, clientDB.Name)
	s.Equal(client.Email, clientDB.Email)
	s.Equal(client.ID, clientDB.ID)
}

func (s *ClientDBTestSuite) TestSaveClient() {

	client, _ := entity.NewClient("Fulano de Tal", "fulano@mail.com")
	err := s.clientDB.Save(*client)
	s.Nil(err)
}

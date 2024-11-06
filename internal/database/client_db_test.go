package database

import (
	"database/sql"
	"testing"

	"github.com.br/devfullcycle/fc-ms-wallet/internal/entity"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
)

type ClientDBTestSuite struct {
	suite.Suite
	db       *sql.DB
	ClientDB *ClientDB
}

func (s *ClientDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Nil(err)
	s.db = db
	db.Exec("CREATE TABLE clients (id VARCHAR(255), name VARCHAR(255), email VARCHAR(255), created_at DATE)")
	s.ClientDB = NewClientDB(db)
}

func (s *ClientDBTestSuite) TearDownSuite() {
	s.db.Exec("DROP TABLE clients")
	s.db.Close()
}

func TestClientDBTestSuite(t *testing.T) {
	suite.Run(t, new(ClientDBTestSuite))
}

func (s *ClientDBTestSuite) TestSave() {
	client := &entity.Client{
		ID:    "1",
		Name:  "Daniel",
		Email: "daniel@mail.com",
	}

	err := s.ClientDB.Save(client)
	s.Nil(err)
}

func (s *ClientDBTestSuite) TestGet() {
	// Criação de um novo cliente e salvamento
	client, _ := entity.NewClient("Daniel", "daniel@mail.com")

	// Verifica se o cliente é salvo corretamente
	err := s.ClientDB.Save(client)
	s.Nil(err, "Erro ao salvar o cliente")

	// Recuperação do cliente salvo
	ClientDB, err := s.ClientDB.Get(client.ID)
	s.Nil(err, "Erro ao buscar o cliente")

	// Verificações dos campos do cliente
	s.Equal(client.ID, ClientDB.ID)
	s.Equal(client.Name, ClientDB.Name)
	s.Equal(client.Email, ClientDB.Email)
}

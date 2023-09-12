package client

import (
	"github.com/C4st3ll4n/wallet/internal/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type ClientGatewayMock struct {
	mock.Mock
}

func (m *ClientGatewayMock) Get(id string) (*entity.Client, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Client), args.Error(1)
}

func (m *ClientGatewayMock) Save(client *entity.Client) error {
	args := m.Called(client)
	return args.Error(0)
}

func TestCreateClient(t *testing.T) {
	m := &ClientGatewayMock{}
	m.On("Save", mock.Anything).Return(nil)
	client, err := entity.NewClient("Any Email", "email@mail.com")

	assert.Nil(t, err)
	assert.NotNil(t, client)

	usecase := NewCreateClientUseCase(m)

	output, err := usecase.Execute(CreateClientInputDTO{
		Name:  "Any Name",
		Email: "Any Email",
	})

	assert.Nil(t, err)
	assert.NotNil(t, output)

	assert.Equal(t, output.Name, "Any Name")
	assert.Equal(t, output.Email, "Any Email")

	m.AssertExpectations(t)
	m.AssertNumberOfCalls(t, "Save", 1)
}

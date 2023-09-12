package account

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

type AccountGatewayMock struct {
	mock.Mock
}

func (m *AccountGatewayMock) Save(account entity.Account) error {
	args := m.Called(account)
	return args.Error(0)
}

func (m *AccountGatewayMock) FindById(id string) (*entity.Account, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Account), args.Error(1)
}

func TestNewCreateAccount(t *testing.T) {
	mockClient := ClientGatewayMock{}
	mockAccount := AccountGatewayMock{}

	client, err := entity.NewClient("Any Email", "email@mail.com")

	mockClient.On("Get", client.ID).Return(client, nil)
	mockAccount.On("Save", mock.Anything).Return(nil)

	assert.Nil(t, err)
	assert.NotNil(t, client)

	usecase := NewCreateAccountUsecase(&mockAccount, &mockClient)
	output, err := usecase.Execute(CreateAccountInputDTO{
		ClientID: client.ID,
	})

	assert.Nil(t, err)
	assert.NotNil(t, output)

	assert.NotNil(t, output.ClientID)

	mockAccount.AssertExpectations(t)
	mockAccount.AssertNumberOfCalls(t, "Save", 1)

	mockClient.AssertExpectations(t)
	mockClient.AssertNumberOfCalls(t, "Get", 1)

}
